default: out/example

clean:
	rm -rf out

test: *.go
	go vet ./...
	go test ./...

out/example: implementation.go cmd/example/main.go
	mkdir -p out
	go build -o out/example ./cmd/example
