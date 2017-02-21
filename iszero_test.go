package iszero_test

import (
	"reflect"
	"testing"

	"github.com/leighmcculloch/go-iszero"
)

func TestValue(t *testing.T) {
	tests := []struct {
		Value        reflect.Value
		ExpectedZero bool
	}{
		// Array
		{reflect.ValueOf([0]int{}), true},
		{reflect.ValueOf([1]int{0}), false},
		// Map
		{reflect.ValueOf(map[int]int(nil)), true},
		{reflect.ValueOf(map[int]int{}), true},
		{reflect.ValueOf(map[int]int{1: 2}), false},
		// Slice
		{reflect.ValueOf([]int(nil)), true},
		{reflect.ValueOf([]int{}), true},
		{reflect.ValueOf([]int{1}), false},
		// String
		{reflect.ValueOf(""), true},
		{reflect.ValueOf("1"), false},
		// Bool
		{reflect.ValueOf(false), true},
		{reflect.ValueOf(true), false},
		// Int
		{reflect.ValueOf(int(0)), true},
		{reflect.ValueOf(int(1)), false},
		{reflect.ValueOf(int8(0)), true},
		{reflect.ValueOf(int8(1)), false},
		{reflect.ValueOf(int16(0)), true},
		{reflect.ValueOf(int16(1)), false},
		{reflect.ValueOf(int32(0)), true},
		{reflect.ValueOf(int32(1)), false},
		{reflect.ValueOf(int64(0)), true},
		{reflect.ValueOf(int64(1)), false},
		// Uint
		{reflect.ValueOf(uint(0)), true},
		{reflect.ValueOf(uint(1)), false},
		{reflect.ValueOf(uint8(0)), true},
		{reflect.ValueOf(uint8(1)), false},
		{reflect.ValueOf(uint16(0)), true},
		{reflect.ValueOf(uint16(1)), false},
		{reflect.ValueOf(uint32(0)), true},
		{reflect.ValueOf(uint32(1)), false},
		{reflect.ValueOf(uint64(0)), true},
		{reflect.ValueOf(uint64(1)), false},
		{reflect.ValueOf(uintptr(0)), true},
		{reflect.ValueOf(uintptr(1)), false},
		// Float
		{reflect.ValueOf(float32(0)), true},
		{reflect.ValueOf(float32(1)), false},
		{reflect.ValueOf(float64(0)), true},
		{reflect.ValueOf(float64(1)), false},
		// Interface
		{reflect.ValueOf(testInterface((*testStruct)(nil))), true},
		{reflect.ValueOf(testInterface(&testStruct{})), false},
		{reflect.ValueOf(testInterface(testStruct{})), false},
		// Ptr
		{reflect.ValueOf((*testStruct)(nil)), true},
		{reflect.ValueOf(&testStruct{}), false},
	}

	for _, test := range tests {
		zero := iszero.Value(test.Value)
		if zero != test.ExpectedZero {
			t.Errorf("Value %#v iszero defined as %v, want %v", test.Value, zero, test.ExpectedZero)
		}
	}
}

type testInterface interface {
	IsTestInterface() bool
}

type testStruct struct{}

func (tis testStruct) IsTestInterface() bool {
	return true
}
