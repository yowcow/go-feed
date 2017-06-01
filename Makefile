.PHONY: all test

all:
	make -C ./aggregator

test:
	make test -C ./aggregator
