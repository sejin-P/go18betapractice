package main

import (
	"constraints"
	"fmt"
)

type Point []int32

func min[T constraints.Ordered](x, y T) T {
	if x < y { // in goland, the compile error occurred, mentioned with 'invalid operation: the operator < is not defined on T
		return x
	}
	return y
}

type Tree[T any] struct {
	cmp  func(T, T) int
	root *node[T]
}

type node[T any] struct {
	left, right *node[T]
	data        T
}

func (bt *Tree[T]) find(val T) **node[T] {
	pl := &bt.root
	for *pl != nil {
		switch cmp := bt.cmp(val, (*pl).data); {
		case cmp < 0:
			pl = &(*pl).left
		case cmp > 0:
			pl = &(*pl).right
		default:
			return pl
		}
	}

	return pl
}

// tilde(~) syntax -> A type approximation such as ~int matches not only int itself, as you'd expect, but also any type derived from int (like MyInt).
func Scale[S ~[]E, E constraints.Integer](s S, c E) S { // no type hinting for constraints package in goland latest... it's GoLand 2021.3.2
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}

	return r
}

func ScaleAndPrint(p Point) {
	r := Scale(p, 2) // calls Scale[Point, int32]
	fmt.Println(r)
}

func main() {
	fmt.Println(min[int](1, 2))
	ScaleAndPrint([]int32{1})

}
