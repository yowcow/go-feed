.PHONY: build gom clean test

BINARY  := aggregator
SUCCESS := \033[1;32m
FAILURE := \033[1;31m
RESET   := \033[m

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
	gom test -v \
		| sed ''/PASS/s//$$(printf "$(SUCCESS)PASS$(RESET)")/'' \
		| sed ''/FAIL/s//$$(printf "$(FAILURE)FAIL$(RESET)")/''
