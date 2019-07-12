package main

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

func main() {
	aObj := a{}
	InitMe(aObj)
	for i := 0; ; i++ {
		go GetInstance().ImplementMe(strconv.Itoa(i))
		if i == 1000 {
			<-time.NewTicker(time.Duration(time.Second * 20)).C
			break
		}
	}
}

type a struct{}

func (obj a) ImplementMe(name string) error {

	fmt.Println(unsafe.Pointer(&obj))
	return nil
}

// Interface - need to be implemented by vendor
type Interface interface {
	ImplementMe(str string) error
}

var intObj Interface

// InitMe - InitMe
func InitMe(obj Interface) {
	intObj = obj
}

// GetInstance - GetInstance
func GetInstance() Interface {
	tmp := intObj
	fmt.Println(unsafe.Pointer(&tmp), unsafe.Pointer(&intObj))
	return tmp
}
