GOPATH := $(shell pwd)

default: build

build:
	cd src/lifx/ && GOPATH=${GOPATH} go install

config:
	GOPATH=${GOPATH} go get ./...

clean:
	rm -rf pkg/
	rm -rf bin/

install: clean build
	sudo cp ./bin/lifx /usr/local/bin

uninstall:
	sudo rm /usr/local/bin/lifx
