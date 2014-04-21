libraries = exec var_dump

.PHONY: all/% format/% test/%

all: $(foreach library,$(libraries),all/$(library))

all/%: % go_get
	cd "$*"; go build

test: $(foreach library,$(libraries),test/$(library))

test/%: % go_get
	cd "$*"; go test

format: $(foreach library,$(libraries),format/$(library))

format/%: %
	cd "$*"; gofmt -d=true -tabs=false -tabwidth=2 -w=true .

go_get:
	go get github.com/imos/go/var_dump
