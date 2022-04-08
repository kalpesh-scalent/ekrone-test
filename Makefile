run-server:
	go run cmd/main.go

run-tests:
	go clean -testcache && go test ./...