package heapescapeanalysis

import "sync"

// Stack-friendly techniques to avoid heap allocation

// 1. Keep variables local and don't return pointers
//
//go:noinline
func stackFriendlyComputation() int {
	x := 42
	y := 24
	return x + y // Values, not pointers
}

// 2. Use value semantics instead of pointer semantics
//
//go:noinline
func processValueSemantics(data [100]int) [100]int {
	for i := range data {
		data[i] *= 2
	}
	return data // Copy returned, but may be optimized
}

// 3. Avoid interface{} when possible
//
//go:noinline
func specificTypeProcessing(x int) int {
	return x * 2 // No boxing to interface{}
}

// 4. Use fixed-size arrays instead of slices when size is known
//
//go:noinline
func fixedArrayProcessing() [10]int {
	var arr [10]int
	for i := range arr {
		arr[i] = i * i
	}
	return arr // Fixed size, likely stack allocated
}

// 5. Avoid capturing variables in closures
//
//go:noinline
func noClosureCapture() func() int {
	// Don't capture local variables
	return func() int {
		return 42 // No variable capture
	}
}

// 6. Use local variables instead of heap-allocated structures
//
//go:noinline
func localVariableProcessing() (int, int, int) {
	a, b, c := 1, 2, 3
	a *= 2
	b *= 3
	c *= 4
	return a, b, c // Multiple return values, no struct needed
}

// 7. Minimize pointer indirection
//
//go:noinline
func directValueAccess(data [5]int) int {
	sum := 0
	for _, v := range data {
		sum += v // Direct value access, no pointers
	}
	return sum
}

// 8. Use sync.Pool for frequently allocated objects to reduce heap pressure
var largeStructPool = sync.Pool{
	New: func() interface{} {
		return &LargeStruct{}
	},
}

//go:noinline
func useSyncPool() *LargeStruct {
	obj := largeStructPool.Get().(*LargeStruct)
	// Reset the object
	*obj = LargeStruct{}
	return obj
}

//go:noinline
func returnToPool(obj *LargeStruct) {
	largeStructPool.Put(obj)
}

// 9. Inline small functions to avoid call overhead
//
//go:noinline
func inlineableFunction(x int) int {
	return x * 2
}

// 10. Use buffer reuse patterns
type BufferProcessor struct {
	buffer []byte
}

func NewBufferProcessor() *BufferProcessor {
	return &BufferProcessor{
		buffer: make([]byte, 0, 1024), // Pre-allocate capacity
	}
}

//go:noinline
func (bp *BufferProcessor) ProcessData(data []byte) []byte {
	// Reuse existing buffer capacity
	bp.buffer = bp.buffer[:0] // Reset length, keep capacity
	bp.buffer = append(bp.buffer, data...)

	// Process in-place
	for i := range bp.buffer {
		bp.buffer[i] = bp.buffer[i] ^ 0xFF
	}

	return bp.buffer
}

// 11. Stack-friendly string building (for small strings)
//
//go:noinline
func stackFriendlyStringBuild() string {
	// For small, known-size strings, direct concatenation can be optimized
	prefix := "Hello"
	suffix := "World"
	return prefix + " " + suffix // May be optimized to stack allocation
}

// 12. Avoid unnecessary slice growth
//
//go:noinline
func preallocatedSlice() []int {
	// Pre-allocate with known capacity
	result := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		result = append(result, i*i)
	}
	return result
}

// 13. Use copy instead of creating new slices
//
//go:noinline
func copyInsteadOfNew(src []int) []int {
	// Create appropriately sized slice
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

// 14. Embedded structs instead of pointer fields
type StackFriendlyStruct struct {
	// Embed instead of pointer
	Data [10]int
	Name [32]byte // Fixed-size instead of string
}

//go:noinline
func createStackFriendlyStruct() StackFriendlyStruct {
	s := StackFriendlyStruct{}
	for i := range s.Data {
		s.Data[i] = i
	}
	copy(s.Name[:], "example")
	return s
}

// 15. Avoid maps for small, known sets
//
//go:noinline
func useArrayInsteadOfMap(key int) string {
	// For small, known sets, arrays can be more stack-friendly
	values := [4]string{"zero", "one", "two", "three"}
	if key >= 0 && key < len(values) {
		return values[key]
	}
	return "unknown"
}

// Comparison functions that cause heap allocations

// Heap allocation - boxing to interface{}
//
//go:noinline
func interfaceBoxing(x int) interface{} {
	return x // int boxed to interface{}
}

// Heap allocation - closure captures variable
//
//go:noinline
func closureCapture() func() int {
	x := 42 // This will escape to heap
	return func() int {
		return x // Captures x
	}
}

// Heap allocation - slice growth without pre-allocation
//
//go:noinline
func sliceGrowth() []int {
	slice := []int{} // Start with zero capacity
	for i := 0; i < 100; i++ {
		slice = append(slice, i*i) // Will grow multiple times
	}
	return slice
}

// Heap allocation - map for small lookup
//
//go:noinline
func mapLookup(key int) string {
	m := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
	}
	if val, ok := m[key]; ok {
		return val
	}
	return "unknown"
}

// Heap allocation - returning pointer to local variable
//
//go:noinline
func returnLocalPointer() *int {
	x := 42 // Will escape to heap
	return &x
}
