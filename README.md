# Go Dependency Injection Example

This is example project to demonstrate DI implementation in Go using dig from Uber.

## Setup

You need to create database by running `./init_db.sh` script.

## Usage

1. `go run cmd/server/main.go` - run server
2. `go run cmd/create_message_command/main.go "Hi, %s."` - run command that creates new message in database
3. `go test -v ./...` - run tests

## API endpoints

`curl 127.0.0.1:8080/hello` - greeting message for unknown user

`curl 127.0.0.1:8080/hello?userId=1` - greeting message for concrete user

`curl -X POST --data "test" 127.0.0.1:8080/echo` - sends back in response request body
