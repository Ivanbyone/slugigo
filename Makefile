fmt:
	go fmt ./...

# Start CI code statis analysis
lint:
	go vet ./...

# Start CI unit tests 
unit:
	go test ./...

# Test with verbose log and coverage info
tcov:
	go test -v -cover
