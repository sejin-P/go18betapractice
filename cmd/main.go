package main

import (
	"constraints"
	"fmt"
)

func min[T constraints.Ordered](x, y T) T {
	if x < y { // in goland, the compile error occurred, mentioned with 'invalid operation: the operator < is not defined on T
		return x
	}
	return y
}

type Tree[T interface{}] struct {
	left, right *Tree[T]
	data        T
}

func main() {
	fmt.Println(min(1, 2))
}
