
SRC=./ulambda
BIN=./cmd/ulc
.PHONY: all build test install tags

all: test install

bin:
	go build $(BIN)

build:
	go build $(SRC)

test:
	go test $(SRC)

install:
	go install $(SRC)

tags:
	gotags **/*.go > tags
