bin_name := slit-$(shell uname -s | tr '[:upper:]' '[:lower:]')-$(shell uname -m)
bin_path := build/$(bin_name)

.PHONY: test
test:
	go vet
	go test

.PHONY: build
build: test
	ENABLE_CGO=0 go build -tags netgo -o $(bin_path) main.go

.PHONY: install
install: build
	sudo cp $(bin_path) /usr/local/bin/slit
