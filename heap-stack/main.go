package main

func main() {
	a := 10
	stack(a)
	heap(&a)
	b := make([]byte, 10)
	readByteStack(b)
}

func stack(a int) int {
	return a * a
}

func heap(a *int) {
	*a++
}

func readByteHeap() []byte {
	return []byte{'a'}
}

func readByteStack(b []byte) {
	b[0] = 'a'
}

/*
# heap-stack
./main.go:17:11: heap a does not escape
./main.go:25:20: readByteStack b does not escape
./main.go:7:7: main &a does not escape
./main.go:9:11: main make([]byte, 10) does not escape
./main.go:22:15: []byte literal escapes to heap
*/
