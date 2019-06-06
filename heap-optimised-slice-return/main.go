package main

import (
	"unsafe"
)

func main() {
	heapOptimised()
	normal()

}

func heapOptimised() string {
	ba := returnBytes()
	ptr := unsafe.Pointer(&ba)
	return *(*string)(ptr)
}

func normal() string {
	ba := returnBytes()
	return string(ba)
}

func returnBytes() []byte {
	str := "roshan"
	ptr := unsafe.Pointer(&str)
	return *(*[]byte)(ptr)
}
