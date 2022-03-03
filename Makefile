fmt:
	gofmt -s -l -w $$(go list -f {{.Dir}} ./... | grep -v /vendor/)

_go_test:
	./cobrademo commaTest --slice 'foo=bar,baz' --array 'foo=bar,baz' --slice 'a=b,c' --array 'a=b,c'

_shipad_dry:
	go build .

test: build _go_test

build:
	go build .