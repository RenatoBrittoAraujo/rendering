run:
	go build cmd/main.go; \
	./main

update:
	go mod tidy