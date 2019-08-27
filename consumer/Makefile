default: build

build:
	go build -mod=vendor -o ./bin/consumer main.go

install:
	@(go install github.com/kaplanmaxe/maxwell-kafka-go/consumer)

.PHONY: build install