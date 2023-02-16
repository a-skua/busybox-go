fmt:
	go fmt ./...

test:
	go test ./... -cover

bench:
	go test ./log -bench Message
