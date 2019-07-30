.PHONY: build

bin_name ?= slit-$$(uname -o)
bin_path ?= build/$(bin_name)

build:
	ENABLE_CGO=0 go build -tags netgo -o $(bin_path) main.go

install: build
	sudo cp $(bin_path) /usr/local/bin/slit
