.PHONY: all clean-all build cleand-deps deps ver

DATE := $(shell git log -1 --format="%cd" --date=short | sed s/-//g)
COUNT := $(shell git rev-list --count HEAD)
COMMIT := $(shell git rev-parse --short HEAD)

BINARYNAME := zabbix-agent-extension-elasticsearch
CONFIG := zabbix-agent-extension-elasticsearch.conf
VERSION := "${DATE}.${COUNT}_${COMMIT}"

LDFLAGS := "-X main.version=${VERSION}"


default: all 

all: clean-all deps build

ver:
	@echo ${VERSION}

clean-all: clean-deps
	@echo Clean builded binaries
	rm -rf .out/
	@echo Done

build:
	@echo Build
	ln -s ${PWD}/vendor/ ${PWD}/vendor/src
	GOPATH="${PWD}/vendor" go build -v -o .out/${BINARYNAME} -ldflags ${LDFLAGS} *.go
	@echo Done

clean-deps:
	@echo Clean dependencies
	rm -rf vendor/*

deps:
	@echo Fetch dependencies
	git submodule update --init

install:
	@echo Install
	cp .out/${BINARYNAME} /usr/bin/${BINARYNAME}
	cp zabbix-agent-extension-elasticsearch.conf /etc/zabbix/zabbix_agentd.d/zabbix-agent-extension-elasticsearch.conf
	@echo Done

remove:
	@echo Remove
	rm /usr/bin/${BINARYNAME}
	rm /etc/zabbix/zabbix_agentd.conf.d/zabbix-agent-extension-elasticsearch.conf
	@echo Done

