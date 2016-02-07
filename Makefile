# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


run: js
	go run example/server.go

js:
	cd example; gopherjs build demo.go -o demo.js

fmt:
	@go fmt *.go
	@go fmt example/*.go

install:
	go get -u github.com/gopherjs/gopherjs
	go get -u github.com/siongui/gopherjs-i18n

clean:
	rm example/demo.js
	rm example/demo.js.map
