.PHONY: all gom build test clean

BINARY := generator

all: gom build

gom: Gomfile
	gom -test install

build:
	cd main && gom build -o ../$(BINARY)

clean:
	rm -rf $(BINARY) vendor

test:
	gom test -v
