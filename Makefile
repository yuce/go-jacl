.PHONY: all cover parser test

all:
	go build ./cmd/jacl

cover:
	go test -coverprofile=cover.profile ./ && go tool cover -html=cover.profile

parser:
	antlr4 -Dlanguage=Go parser/Jacl.g4

test:
	go test ./