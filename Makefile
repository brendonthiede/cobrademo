fmt:
	gofmt -s -l -w $$(go list -f {{.Dir}} ./... | grep -v /vendor/)

build:
	go build .

_go_comma_test:
	./cobrademo commaTest --slice 'foo=bar,baz' --array 'foo=bar,baz' --slice 'a=b,c' --array 'a=b,c'

comma_test: build _go_comma_test

_go_map_test:
	./cobrademo mapTest

map_test: build _go_map_test
