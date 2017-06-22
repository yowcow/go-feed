.PHONY: all test

all:

test:
	go test ./aggregator ./generator -v
