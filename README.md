# Go Heap Escape Analysis Examples

This repository demonstrates various cases in Go where local variables "escape to the heap" instead of being allocated on the stack, and more importantly, **how to write Go code that avoids heap allocations** for better performance.

## What is Heap Escape?

When a function creates a local variable, the Go compiler decides whether to allocate it on the stack or heap based on escape analysis. Variables "escape to the heap" when:

1. They need to outlive the function that created them
2. They are too large for the stack
3. Their size is not known at compile time
4. They are assigned to interfaces
5. They are captured by closures

## Repository Structure

This repository contains several files demonstrating different aspects of Go memory allocation:

### **Core Files**
- **`functions.go`** - Functions that cause heap escapes (what NOT to do)
- **`keep_on_stack.go`** - Stack-optimized functions (what TO do)

### **Benchmark Files**
- **`benchmark_test.go`** - Benchmarks heap-escaping functions to show allocation costs
- **`stack_benchmark_test.go`** - Benchmarks stack-optimized functions and comparisons

### **How to run**

To run the benchmarks, use the following command:

```bash
make
```

### **Goals of Each Benchmark**

| Benchmark Category | Purpose | What It Demonstrates |
|-------------------|---------|---------------------|
| `BenchmarkReturnPointer` vs `BenchmarkReturnValue` | Show pointer return vs value return cost | Returning pointers forces heap allocation |
| `BenchmarkReturnLarge*` | Compare large struct allocation strategies | Size matters for escape analysis |
| `BenchmarkAssignToInterface` | Show interface boxing overhead | interface{} causes escapes |
| `BenchmarkCreateSlice*` | Compare slice allocation patterns | Dynamic allocation vs pre-allocation |
| `BenchmarkComparison_*` | Side-by-side performance comparisons | Direct measurement of optimization impact |
| Stack optimization benchmarks | Validate stack-friendly patterns | Prove techniques actually work |

## How to Write Go Code That Avoids Heap Allocations

### 🎯 **Core Principles**

1. **Keep data local and scoped**
2. **Use value semantics over pointer semantics**
3. **Pre-allocate when size is known**
4. **Avoid interface{} when possible**
5. **Minimize pointer indirection**

### 📋 **Detailed Techniques**

#### **1. Function Design Patterns**

```go
// ❌ BAD: Returns pointer (forces heap allocation)
func createData() *Data {
    d := Data{value: 42}
    return &d  // d escapes to heap
}

// ✅ GOOD: Return by value (stack allocation)
func createData() Data {
    return Data{value: 42}  // Stays on stack
}

// ✅ BETTER: Modify via pointer parameter (reuse memory)
func initData(d *Data) {
    d.value = 42  // No allocation, modifies existing memory
}
```

#### **2. Data Structure Choices**

```go
// ❌ BAD: Dynamic slice without pre-allocation
func processItems() []int {
    result := []int{}  // Will grow multiple times
    for i := 0; i < 100; i++ {
        result = append(result, i)
    }
    return result
}

// ✅ GOOD: Pre-allocated slice
func processItems() []int {
    result := make([]int, 0, 100)  // Pre-allocate capacity
    for i := 0; i < 100; i++ {
        result = append(result, i)
    }
    return result
}

// ✅ BETTER: Fixed-size array when size is known
func processItems() [100]int {
    var result [100]int  // Stack allocated
    for i := range result {
        result[i] = i
    }
    return result
}
```

#### **3. Lookup Optimizations**

```go
// ❌ BAD: Map for small, static datasets
func getStatusName(code int) string {
    statusMap := map[int]string{
        200: "OK",
        404: "Not Found",
        500: "Internal Error",
    }
    return statusMap[code]
}

// ✅ GOOD: Array lookup for small, known sets
func getStatusName(code int) string {
    statuses := map[int]string{
        200: "OK",
        404: "Not Found", 
        500: "Internal Error",
    }
    if code >= 0 && code < len(statuses) {
        return statuses[code]
    }
    return "Unknown"
}
```

#### **4. Interface Usage**

```go
// ❌ BAD: Unnecessary interface{} boxing
func process(data interface{}) {
    // Forces heap allocation for boxing
}

// ✅ GOOD: Specific types
func processInt(data int) {
    // No boxing, stays on stack
}

// ✅ GOOD: Type-specific functions
func processString(data string) {
    // Type-safe and efficient
}
```

#### **5. String Building**

```go
// ❌ BAD: Multiple allocations for complex strings
func buildMessage(name, action string) string {
    var builder strings.Builder
    builder.WriteString("User ")
    builder.WriteString(name)
    builder.WriteString(" performed ")
    builder.WriteString(action)
    return builder.String()
}

// ✅ GOOD: Simple concatenation for small strings
func buildMessage(name, action string) string {
    return "User " + name + " performed " + action
}
```

#### **6. Buffer Reuse Patterns**

```go
// ✅ EXCELLENT: Reusable buffer pattern
type Processor struct {
    buffer []byte
}

func NewProcessor() *Processor {
    return &Processor{
        buffer: make([]byte, 0, 4096),  // Pre-allocate
    }
}

func (p *Processor) Process(data []byte) []byte {
    p.buffer = p.buffer[:0]  // Reset length, keep capacity
    p.buffer = append(p.buffer, data...)
    // Process in place...
    return p.buffer
}
```

#### **7. Closure Optimization**

```go
// ❌ BAD: Closure captures large variables
func createHandler(largeData []byte) func() {
    return func() {
        // largeData escapes to heap
        process(largeData)
    }
}

// ✅ GOOD: Pass what you need
func createHandler() func([]byte) {
    return func(data []byte) {
        // No capture, data passed as parameter
        process(data)
    }
}
```

### 🛠 **Tools and Verification**

#### **Escape Analysis**
```bash
# Basic escape analysis
go build -gcflags="-m" .

# Detailed escape analysis  
go build -gcflags="-m -m" .

# Look for:
# - "moved to heap: variable" (bad)
# - "does not escape" (good)
# - "escapes to heap" (allocation needed)
```

#### **Benchmarking**
```bash
# Run all benchmarks with memory stats
go test -bench=. -benchmem

# Focus on specific patterns
go test -bench=BenchmarkComparison -benchmem

# Look for:
# - 0 B/op, 0 allocs/op (perfect - stack allocation)
# - Low B/op, 1 allocs/op (acceptable - single allocation)
# - High allocs/op (bad - multiple allocations)
```

#### **Profiling**
```bash
# Memory profiling
go test -bench=. -memprofile=mem.prof
go tool pprof -http=:8080 mem.prof

# CPU profiling  
go test -bench=. -cpuprofile=cpu.prof
go tool pprof -http=:8080 cpu.prof
```

### 📊 **Quick Reference: Stack vs Heap Indicators**

| Pattern | Stack Allocation | Heap Allocation |
|---------|------------------|-----------------|
| **Function returns** | `return value` | `return &localVar` |
| **Data structures** | `[N]Type` (arrays) | `[]Type` (slices), `map[K]V` |
| **Interfaces** | Specific types | `interface{}` boxing |
| **Closures** | No captures | Captures variables |
| **Size** | Small, fixed | Large, dynamic |
| **Lifetime** | Function scope | Beyond function |

### ⚡ **Performance Impact**

- **Stack allocation**: ~1-10 ns/op, 0 B/op, 0 allocs/op
- **Single heap allocation**: ~10-50 ns/op, size B/op, 1 allocs/op  
- **Multiple allocations**: ~100+ ns/op, large B/op, many allocs/op

### 🎯 **Best Practices Summary**

#### **DO:**
- ✅ Return values instead of pointers
- ✅ Use fixed arrays for known sizes
- ✅ Pre-allocate slice capacity
- ✅ Use specific types over interface{}
- ✅ Keep variable scope minimal
- ✅ Reuse buffers with sync.Pool
- ✅ Profile before optimizing

#### **DON'T:**
- ❌ Return pointers to local variables
- ❌ Use interface{} unnecessarily
- ❌ Create closures that capture large data
- ❌ Grow slices without pre-allocation
- ❌ Use maps for small, static data
- ❌ Optimize without measuring

**Remember**: Premature optimization is the root of all evil. Always profile first, then optimize the actual bottlenecks with these techniques!
