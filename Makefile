run:
	@go build -o bin/rendering cmd/main.go 
	@./bin/rendering 2>&1 | zap-pretty

update:
	go mod tidy