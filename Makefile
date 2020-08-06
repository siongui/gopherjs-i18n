# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
ifndef TRAVIS
	export GOROOT=$(realpath ../paligo/go)
	export GOPATH=$(realpath ../paligo)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

PKG="github.com/siongui/gopherjs-i18n"

test:
	go test -v locale.go locale_test.go

run: js
	go run example/server.go

js:
	cd example; gopherjs build demo.go -o demo.js

local:
	@[ -d ${GOPATH}/src/${PKG}/ ] || mkdir -p ${GOPATH}/src/${PKG}/
	@cp *.go ${GOPATH}/src/${PKG}/

fmt:
	@go fmt *.go
	@go fmt example/*.go
	@go fmt tool/*.go

install:
	go get -u github.com/gopherjs/gopherjs
	go get -u github.com/siongui/gopherjs-i18n

clean:
	rm example/demo.js
	rm example/demo.js.map
	rm -rf ${GOPATH}/src/${PKG}/
