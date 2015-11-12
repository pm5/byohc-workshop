
SRC=./ulambda
.PHONY: all build test install tags

all: test install

build:
	go build $(SRC)

test:
	go test $(SRC)

install:
	go install $(SRC)

tags:
	gotags **/*.go > tags
