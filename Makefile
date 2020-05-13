.PHONY: all clean-all build cleand-deps deps ver make-gopath

DATE := $(shell git log -1 --format="%cd" --date=short | sed s/-//g)
COUNT := $(shell git rev-list --count HEAD)
COMMIT := $(shell git rev-parse --short HEAD)
CWD := $(shell pwd)

BINARYNAME := zabbix-agent-extension-elasticsearch
CONFIG := configs/zabbix_agentd.d/zabbix-agent-extension-elasticsearch.conf
VERSION := "${DATE}.${COUNT}_${COMMIT}"

LDFLAGS := "-X main.version=${VERSION}"


default: all

all: clean-all make-gopath deps build

ver:
	@echo ${VERSION}

clean-all: clean-deps
	@echo Clean builded binaries
	rm -rf .out/
	rm -rf .gopath/
	@echo Done

build:
	@echo Build
	cd ${CWD}/.gopath/src/${BINARYNAME}; \
		GOPATH=${CWD}/.gopath \
		go build  -v -o ${CWD}/.out/${BINARYNAME} -ldflags ${LDFLAGS} *.go
	@echo Done

clean-deps:
	@echo Clean dependencies
	rm -rf ${CWD}/.gopath/src/${BINARYNAME}/vendor/

deps:
	@echo Fetch dependencies
	cd ${CWD}/.gopath/src/${BINARYNAME}; \
		GOPATH=${CWD}/.gopath \
		dep ensure -v

make-gopath:
	@echo Creating GOPATH
	mkdir -p .gopath/src
	ln -s ${CWD}/cmd/${BINARYNAME} ${CWD}/.gopath/src/${BINARYNAME}

install:
	@echo Install
	cp .out/${BINARYNAME} /usr/bin/${BINARYNAME}
	cp configs/zabbix_agentd.d/zabbix-agent-extension-elasticsearch.conf \
		/etc/zabbix/zabbix_agentd.conf.d/zabbix-agent-extension-elasticsearch.conf
	@echo Done

remove:
	@echo Remove
	rm -f /usr/bin/${BINARYNAME}
	rm -f /etc/zabbix/zabbix_agentd.conf.d/zabbix-agent-extension-elasticsearch.conf
	@echo Done
