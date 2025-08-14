package heapescapeanalysis

import (
	"testing"
)

// Global variable to prevent compiler optimizations
var result interface{}

// Benchmarks comparing stack vs heap allocations
func BenchmarkReturnPointer(b *testing.B) {
	var r *int

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = returnPointer()
	}
	result = r
}

func BenchmarkReturnValue(b *testing.B) {
	var r int

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = returnValue()
	}
	result = r
}

func BenchmarkReturnLargePointer(b *testing.B) {
	var r *LargeStruct

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = returnLargePointer()
	}
	result = r
}

func BenchmarkReturnLargeValue(b *testing.B) {
	var r LargeStruct

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = returnLargeValue()
	}
	result = r
}

func BenchmarkAssignToInterface(b *testing.B) {
	var r interface{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = assignToInterface()
	}
	result = r
}

func BenchmarkCreateSliceSmall(b *testing.B) {
	var r []int

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = createSlice(10) // Small slice
	}
	result = r
}

func BenchmarkCreateSliceLarge(b *testing.B) {
	var r []int

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = createSlice(10000) // Large slice
	}
	result = r
}

func BenchmarkCreateClosure(b *testing.B) {
	var r func() int

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = createClosure()
	}
	result = r
}

func BenchmarkSendToChannel(b *testing.B) {
	ch := make(chan int, 1)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sendToChannel(ch)
		<-ch // Consume to prevent blocking
	}
}

func BenchmarkAssignToMap(b *testing.B) {
	m := make(map[string]*int)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assignToMap(m)
	}
}

// Benchmarks for optimized pointer parameter versions
func BenchmarkSetValueViaPointer(b *testing.B) {
	var r int

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setValueViaPointer(&r)
	}
	result = r
}

func BenchmarkSetLargeStructViaPointer(b *testing.B) {
	var r LargeStruct

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setLargeStructViaPointer(&r)
	}
	result = r
}

func BenchmarkInitSliceViaPointer(b *testing.B) {
	var r []int

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		initSliceViaPointer(&r, 10)
	}
	result = r
}

func BenchmarkInitSliceViaPointerLarge(b *testing.B) {
	var r []int

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		initSliceViaPointer(&r, 10000)
	}
	result = r
}
