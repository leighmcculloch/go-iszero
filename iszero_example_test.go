package iszero_test

import (
	"fmt"
	"reflect"

	"github.com/leighmcculloch/go-iszero"
)

func ExampleValue() {
	var s0 string
	s0Zero := iszero.Value(reflect.ValueOf(s0))

	var s1 string = "hello world"
	s1Zero := iszero.Value(reflect.ValueOf(s1))

	fmt.Println("s0", s0Zero)
	fmt.Println("s1", s1Zero)

	// Output:
	// s0 true
	// s1 false
}
