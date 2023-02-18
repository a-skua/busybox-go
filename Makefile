fmt:
	cd option; go fmt ./...
	cd log; go fmt ./...

test:
	cd option; go test ./... -cover
	cd log; go test ./... -cover

bench:
	cd log; go test ./... -bench Message
