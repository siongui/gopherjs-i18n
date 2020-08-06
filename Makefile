# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
ifndef TRAVIS
	export GOROOT=$(realpath ../paligo/go)
	export GOPATH=$(realpath ../paligo)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

PKG="github.com/siongui/gopherjs-i18n"
PKGDIR=${GOPATH}/src/${PKG}/

poToJson: fmt
	@echo "\033[92mTest converting PO file to JSON file ...\033[0m"
	cd po2json; go test -v

test:
	@echo "\033[92mTest this package ...\033[0m"
	go test -v locale.go locale_test.go

run: js
	go run example/server.go

js:
	cd example; gopherjs build demo.go -o demo.js

local:
	@echo "\033[92mInstall this package locally ...\033[0m"
	@[ -d $(PKGDIR) ] || mkdir -p $(PKGDIR)
	@cp *.go $(PKGDIR)

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
	@go fmt example/*.go
	@go fmt po2json/*.go

install:
	go get -u github.com/gopherjs/gopherjs
	go get -u github.com/siongui/gopherjs-i18n

clean:
	rm example/demo.js
	rm example/demo.js.map
	rm -rf ${GOPATH}/src/${PKG}/
