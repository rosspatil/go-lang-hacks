package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data := heapOptimised()
	fmt.Println(data)
	// normal()

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
