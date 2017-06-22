.PHONY: build gom clean test

BINARY := aggregator

all: urls.json build

urls.json:
	echo "[\n]" > $@

build:
	cd main && go build -o ../$(BINARY)

test:
	go test -v

clean:
	rm -rf $(BINARY)
