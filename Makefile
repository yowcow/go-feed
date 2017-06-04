.PHONY: all test

all:
	make -C ./aggregator
	make -C ./generator

test:
	make test -C ./aggregator
	make test -C ./generator
