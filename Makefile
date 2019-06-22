.PHONY: all parser test

all:
	go build ./cmd/jacl

parser:
	antlr4 -Dlanguage=Go parser/Jacl.g4

test:
	go test ./