.PHONY: all
all: escape-analysis benchmark

.PHONY: escape-analysis
escape-analysis:
	go build -gcflags="-m" .

.PHONY: benchmark
benchmark:
	go test -bench=. -benchmem