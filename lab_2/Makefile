default: out/example

clean:
	rm -rf out

test: 
	go test -v

out/example: implementation.go cmd/example/main.go
	mkdir -p out
	go build -o out/example ./cmd/example
