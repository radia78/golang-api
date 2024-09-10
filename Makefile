build:
	@go build -o bin/pf_tracker cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/pf_tracker