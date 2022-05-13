MAIN_PKG:=wechatDemo
MAIN_PREFIX=$(dir $(MAIN_PKG))
MAIN=$(subst $(MAIN_PREFIX), , $(MAIN_PKG))
BIN=$(strip $(MAIN))

BASEDIR=$(shell pwd)
export GOPATH=$(shell pwd)/../../../../../
export AIBEE_KUBERNETES_IDC=suzhou

build:
	go build -tags=jsoniter -x -o run/$(BIN) /root/go/src/$(MAIN_PKG)

dev:
	glr -main /root/go/src/$(MAIN_PKG) -wd run -delay 2000 -args $(ARG)

run: build
	cd run && ./$(BIN) $(ARG)

.PHONY: build
