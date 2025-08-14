package heapescapeanalysis

// Case 1: Returning pointer to local variable
//
//go:noinline
func returnPointer() *int {
	x := 42 // This will escape to heap
	return &x
}

// Case 2: Returning value (stays on stack)
//
//go:noinline
func returnValue() int {
	x := 42 // This stays on stack
	return x
}

// Case 3: Large struct that escapes due to size
type LargeStruct struct {
	data [3000]int // Large array
}

//go:noinline
func returnLargePointer() *LargeStruct {
	s := LargeStruct{} // Will escape to heap
	return &s
}

//go:noinline
func returnLargeValue() LargeStruct {
	s := LargeStruct{} // May escape due to size
	return s
}

// Case 4: Interface assignment causes escape
//
//go:noinline
func assignToInterface() interface{} {
	x := 42 // Will escape to heap when assigned to interface{}
	return x
}

// Case 5: Slice with dynamic size
//
//go:noinline
func createSlice(size int) []int {
	s := make([]int, size) // May escape to heap depending on size
	return s
}

// Case 6: Closure capturing variables
//
//go:noinline
func createClosure() func() int {
	x := 42 // Will escape to heap because closure captures it
	return func() int {
		return x
	}
}

// Case 7: Channel operations
//
//go:noinline
func sendToChannel(ch chan int) {
	x := 42 // May escape depending on channel usage
	ch <- x
}

// Case 8: Map operations
//
//go:noinline
func assignToMap(m map[string]*int) {
	x := 42 // Will escape to heap when stored in map
	m["key"] = &x
}

// Optimized versions using pointer parameters to avoid allocations

// Case 9: Using pointer parameter instead of returning value
//
//go:noinline
func setValueViaPointer(result *int) {
	*result = 42 // No allocation - modifies existing memory
}

// Case 10: Using pointer parameter for large struct
//
//go:noinline
func setLargeStructViaPointer(result *LargeStruct) {
	// Initialize the struct in-place, no allocation
	for i := range result.data {
		result.data[i] = i
	}
}

// Case 11: Slice initialization via pointer parameter
//
//go:noinline
func initSliceViaPointer(result *[]int, size int) {
	// Reuse existing slice capacity if possible
	if cap(*result) >= size {
		*result = (*result)[:size]
	} else {
		*result = make([]int, size) // Only allocate if needed
	}
	for i := range *result {
		(*result)[i] = i
	}
}
