package main

import (
	"strings"
	"testing"
)

// Global variable to prevent compiler optimizations
var stackResult interface{}

// Benchmarks for stack optimization techniques

func BenchmarkStackFriendlyComputation(b *testing.B) {
	var r int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = stackFriendlyComputation()
	}
	stackResult = r
}

func BenchmarkProcessValueSemantics(b *testing.B) {
	data := [100]int{}
	for i := range data {
		data[i] = i
	}
	var r [100]int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = processValueSemantics(data)
	}
	stackResult = r
}

func BenchmarkSpecificTypeProcessing(b *testing.B) {
	var r int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = specificTypeProcessing(42)
	}
	stackResult = r
}

func BenchmarkFixedArrayProcessing(b *testing.B) {
	var r [10]int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = fixedArrayProcessing()
	}
	stackResult = r
}

func BenchmarkNoClosureCapture(b *testing.B) {
	var r func() int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = noClosureCapture()
	}
	stackResult = r
}

func BenchmarkLocalVariableProcessing(b *testing.B) {
	var a, x, c int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a, x, c = localVariableProcessing()
	}
	stackResult = [3]int{a, x, c}
}

func BenchmarkDirectValueAccess(b *testing.B) {
	data := [5]int{1, 2, 3, 4, 5}
	var r int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = directValueAccess(data)
	}
	stackResult = r
}

func BenchmarkUseSyncPool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		obj := useSyncPool()
		returnToPool(obj)
	}
}

func BenchmarkInlineableFunction(b *testing.B) {
	var r int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = inlineableFunction(42)
	}
	stackResult = r
}

func BenchmarkBufferProcessor(b *testing.B) {
	bp := NewBufferProcessor()
	data := []byte("hello world")
	var r []byte
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = bp.ProcessData(data)
	}
	stackResult = r
}

func BenchmarkStackFriendlyStringBuild(b *testing.B) {
	var r string
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = stackFriendlyStringBuild()
	}
	stackResult = r
}

func BenchmarkPreallocatedSlice(b *testing.B) {
	var r []int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = preallocatedSlice()
	}
	stackResult = r
}

func BenchmarkCopyInsteadOfNew(b *testing.B) {
	src := make([]int, 100)
	for i := range src {
		src[i] = i
	}
	var r []int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = copyInsteadOfNew(src)
	}
	stackResult = r
}

func BenchmarkCreateStackFriendlyStruct(b *testing.B) {
	var r StackFriendlyStruct
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = createStackFriendlyStruct()
	}
	stackResult = r
}

func BenchmarkUseArrayInsteadOfMap(b *testing.B) {
	var r string
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = useArrayInsteadOfMap(2)
	}
	stackResult = r
}

// Comparison benchmarks showing heap allocation alternatives

func BenchmarkInterfaceBoxing(b *testing.B) {
	var r interface{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = interfaceBoxing(42)
	}
	stackResult = r
}

func BenchmarkClosureCapture(b *testing.B) {
	var r func() int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = closureCapture()
	}
	stackResult = r
}

func BenchmarkSliceGrowth(b *testing.B) {
	var r []int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = sliceGrowth()
	}
	stackResult = r
}

func BenchmarkMapLookup(b *testing.B) {
	var r string
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = mapLookup(2)
	}
	stackResult = r
}

func BenchmarkReturnLocalPointer(b *testing.B) {
	var r *int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = returnLocalPointer()
	}
	stackResult = r
}

func BenchmarkHeapStringBuilder(b *testing.B) {
	var r string
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.WriteString("Hello")
		builder.WriteString(" ")
		builder.WriteString("World")
		r = builder.String()
	}
	stackResult = r
}

// Combined comparison benchmarks to highlight differences

func BenchmarkComparison_StackVsHeapString(b *testing.B) {
	b.Run("Stack-Friendly", func(b *testing.B) {
		var r string
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r = stackFriendlyStringBuild()
		}
		stackResult = r
	})

	b.Run("Heap-StringBuilder", func(b *testing.B) {
		var r string
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var builder strings.Builder
			builder.WriteString("Hello")
			builder.WriteString(" ")
			builder.WriteString("World")
			r = builder.String()
		}
		stackResult = r
	})
}

func BenchmarkComparison_ArrayVsSlice(b *testing.B) {
	b.Run("Fixed-Array", func(b *testing.B) {
		var r [10]int
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r = fixedArrayProcessing()
		}
		stackResult = r
	})

	b.Run("Preallocated-Slice", func(b *testing.B) {
		var r []int
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// Simulate same size as array
			result := make([]int, 10)
			for j := range result {
				result[j] = j * j
			}
			r = result
		}
		stackResult = r
	})
}

func BenchmarkComparison_DirectVsInterface(b *testing.B) {
	b.Run("Direct-Type", func(b *testing.B) {
		var r int
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r = specificTypeProcessing(42)
		}
		stackResult = r
	})

	b.Run("Interface-Boxing", func(b *testing.B) {
		var r interface{}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r = interfaceBoxing(42)
		}
		stackResult = r
	})
}

func BenchmarkComparison_ArrayVsMap(b *testing.B) {
	b.Run("Array-Lookup", func(b *testing.B) {
		var r string
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r = useArrayInsteadOfMap(2)
		}
		stackResult = r
	})

	b.Run("Map-Lookup", func(b *testing.B) {
		var r string
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r = mapLookup(2)
		}
		stackResult = r
	})
}

func BenchmarkComparison_SliceAllocation(b *testing.B) {
	b.Run("Preallocated", func(b *testing.B) {
		var r []int
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r = preallocatedSlice()
		}
		stackResult = r
	})

	b.Run("Growing", func(b *testing.B) {
		var r []int
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r = sliceGrowth()
		}
		stackResult = r
	})
}
