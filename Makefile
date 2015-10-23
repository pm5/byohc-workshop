
BIN=json2lambda

.PHONY: all test tags

all: $(BIN)

test:
	go test ./ulambda

json2lambda: cmd/json2lambda/main.go
	go build -o bin/json2lambda $<

tags:
	gotags **/*.go > tags
