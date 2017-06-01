.PHONY: build gom clean test

BINARY  := aggregator

all: gom urls.json build

gom: Gomfile
	gom -test install

urls.json:
	echo "[\n]" > $@

build:
	cd main && gom build -o ../$(BINARY)

clean:
	rm -rf aggregator vendor

test:
	gom test -v
