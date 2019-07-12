package main

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

// ReferenceCountable following reference countable interface
// We have provided inbuilt embeddable implementation of the reference countable pool
// This interface just provides the extensibility for the implementation
type ReferenceCountable interface {
	// Method to set the current instance
	SetInstance(i interface{})
	// Method to increment the reference count
	IncrementReferenceCount()
	// Method to decrement reference count
	DecrementReferenceCount()
}

// ReferenceCounter representing reference
// This struct is supposed to be embedded inside the object to be pooled
// Along with that incrementing and decrementing the references is highly important specifically around routines
type ReferenceCounter struct {
	count       *uint32                 `sql:"-" json:"-" yaml:"-"`
	destination *sync.Pool              `sql:"-" json:"-" yaml:"-"`
	released    *uint32                 `sql:"-" json:"-" yaml:"-"`
	Instance    interface{}             `sql:"-" json:"-" yaml:"-"`
	reset       func(interface{}) error `sql:"-" json:"-" yaml:"-"`
	id          uint32                  `sql:"-" json:"-" yaml:"-"`
}

// IncrementReferenceCount to increment a reference
func (r ReferenceCounter) IncrementReferenceCount() {
	atomic.AddUint32(r.count, 1)
}

// DecrementReferenceCount to decrement a reference
// If the reference count goes to zero, the object is put back inside the pool
func (r ReferenceCounter) DecrementReferenceCount() {
	if atomic.LoadUint32(r.count) == 0 {
		panic("this should not happen =>" + reflect.TypeOf(r.Instance).String())
	}
	if atomic.AddUint32(r.count, ^uint32(0)) == 0 {
		atomic.AddUint32(r.released, 1)
		if err := r.reset(r.Instance); err != nil {
			panic("error while resetting an instance => " + err.Error())
		}
		r.destination.Put(r.Instance)
		r.Instance = nil
	}
}

// SetInstance to set the current instance
func (r *ReferenceCounter) SetInstance(i interface{}) {
	r.Instance = i
}

type referenceCountedPool struct {
	pool       *sync.Pool
	factory    func() ReferenceCountable
	returned   uint32
	allocated  uint32
	referenced uint32
}

// Method to create a new pool
func NewReferenceCountedPool(factory func(referenceCounter ReferenceCounter) ReferenceCountable, reset func(interface{}) error) *referenceCountedPool {
	p := new(referenceCountedPool)
	p.pool = new(sync.Pool)
	p.pool.New = func() interface{} {
		// fmt.Println("New Object requested "
		// Incrementing allocated count
		atomic.AddUint32(&p.allocated, 1)
		fmt.Println(atomic.LoadUint32(&p.allocated))
		c := factory(ReferenceCounter{
			count:       new(uint32),
			destination: p.pool,
			released:    &p.returned,
			reset:       reset,
			id:          p.allocated,
		})
		return c
	}
	return p
}

// Method to get new object
func (p *referenceCountedPool) Get() ReferenceCountable {
	c := p.pool.Get().(ReferenceCountable)
	c.SetInstance(c)
	atomic.AddUint32(&p.referenced, 1)
	c.IncrementReferenceCount()
	return c
}

// Method to return reference counted pool stats
func (p *referenceCountedPool) Stats() map[string]interface{} {
	return map[string]interface{}{"allocated": atomic.LoadUint32(&p.allocated), "referenced": p.referenced, "returned": p.returned}
}

type Event struct {
	ReferenceCounter
	Name      string    `json:"name"`
	Log       string    `json:"log"`
	Timestamp time.Time `json:"timestamp"`
}

// Method to reset event
func (e *Event) Reset() {
	e.Name = ""
	e.Log = ""
	e.Timestamp = time.Time{}
}

// Method to reset Event
// Used by reference countable pool
func ResetEvent(i interface{}) error {
	obj, ok := i.(*Event)
	if !ok {
		errors.New("illegal object sent to ResetEvent")
	}
	obj.Reset()
	return nil
}

// NewEvent to create new event
func NewEvent(name, log string, timestamp time.Time) *Event {
	e := AcquireEvent()
	e.Name = name
	e.Log = log
	e.Timestamp = timestamp
	return e
}

var eventPool = NewReferenceCountedPool(
	func(counter ReferenceCounter) ReferenceCountable {

		br := new(Event)
		br.ReferenceCounter = counter
		return br
	}, ResetEvent)

// Method to get new Event
func AcquireEvent() *Event {
	return eventPool.Get().(*Event)
}

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			e := NewEvent(string(i), "this is a test log", time.Now())
			e.DecrementReferenceCount()
		}(i)
		// <-time.NewTicker(time.Second * 1).C

	}
	fmt.Println("finished  ", eventPool.Stats())
	<-time.NewTicker(time.Second * time.Duration(300)).C
}
