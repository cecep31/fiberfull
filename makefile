bin:
	@mkdir bin
run:
	@go build -o bin
	@./bin/fiberfull
build:
	@go build -o bin