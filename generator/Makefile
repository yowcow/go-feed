.PHONY: all build test clean

BINARY := generator

all: build

build:
	cd main && go build -o ../$(BINARY)

test:
	go test -v

clean:
	rm -rf $(BINARY)
