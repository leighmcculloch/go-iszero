package iszero_test

import (
	"reflect"
	"testing"

	"github.com/leighmcculloch/go-iszero"
)

func BenchmarkValueArray(b *testing.B) {
	v := reflect.ValueOf([1]int{0})
	for i := 0; i < b.N; i++ {
		iszero.Value(v)
	}
}

func BenchmarkValueMap(b *testing.B) {
	v := reflect.ValueOf(map[int]int{1: 2})
	for i := 0; i < b.N; i++ {
		iszero.Value(v)
	}
}

func BenchmarkValueSlice(b *testing.B) {
	v := reflect.ValueOf([]int{1})
	for i := 0; i < b.N; i++ {
		iszero.Value(v)
	}
}

func BenchmarkValueString(b *testing.B) {
	v := reflect.ValueOf("1")
	for i := 0; i < b.N; i++ {
		iszero.Value(v)
	}
}

func BenchmarkValueBool(b *testing.B) {
	v := reflect.ValueOf(true)
	for i := 0; i < b.N; i++ {
		iszero.Value(v)
	}
}

func BenchmarkValueInt(b *testing.B) {
	v := reflect.ValueOf(int(1))
	for i := 0; i < b.N; i++ {
		iszero.Value(v)
	}
}

func BenchmarkValueUint(b *testing.B) {
	v := reflect.ValueOf(uint(1))
	for i := 0; i < b.N; i++ {
		iszero.Value(v)
	}
}

func BenchmarkValueFloat32(b *testing.B) {
	v := reflect.ValueOf(float32(1))
	for i := 0; i < b.N; i++ {
		iszero.Value(v)
	}
}

func BenchmarkValueFloat64(b *testing.B) {
	v := reflect.ValueOf(float64(1))
	for i := 0; i < b.N; i++ {
		iszero.Value(v)
	}
}

func BenchmarkValueInterface(b *testing.B) {
	v := reflect.ValueOf(testInterface(testStruct{}))
	for i := 0; i < b.N; i++ {
		iszero.Value(v)
	}
}

func BenchmarkValuePtr(b *testing.B) {
	v := reflect.ValueOf(&testStruct{})
	for i := 0; i < b.N; i++ {
		iszero.Value(v)
	}
}
