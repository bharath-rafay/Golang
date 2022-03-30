BIN=Project

all:build

init:clean
	go mod init

tidy:init
	go mod tidy

build:tidy
	go build -o ${BIN} main.go

run:build
	./${BIN}

clean:
	rm -rf go.* ${BIN}