.PHONY: mod/download mod/tidy

mod/download:
	go mod download

mod/tidy:
	go mod tidy

.PHONY: test run

test:
	go test  -v -p 1 ./internal/...

run:
	go run ./main.go
