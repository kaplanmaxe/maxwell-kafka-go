default: build

build:
	go build -mod=vendor -o ./bin/producer main.go

install:
	@(go install github.com/kaplanmaxe/maxwell-kafka-go/producer)

.PHONY: build install