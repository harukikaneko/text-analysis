run:
	TEXT_ANALYSIS_ENDPOINT=http://localhost:3008 go run main.go

install:
	go mod download

build:
	go build -o target/bff

clean:
	rm -rf ./target && go clean && go mod tidy && rm -f di/wire_gen.go

gen:
	rm -f di/wire_gen.go && /Users/kanekoharuaki/go/bin/wire di/wire.go

fmt:
	go fmt ./...

test:
	go test ./...