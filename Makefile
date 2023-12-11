.PHONY: test
test:
	go test ./...

.PHONY: build
build: bin
	go build -o bin/gogetter .

.PHONY: install
install:
	go install .

generate:
	go generate ./...

bin:
	mkdir bin
