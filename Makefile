.PHONY: all verion test clean install

VERSION := $(shell git log -1 --format=%cd.$(shell git rev-list --count HEAD).%h --date=format:%Y%m%d)
BINARY  := zabbix-agent-extension-elasticsearch
PREFIX  := .

LDFLAGS := "-X main.version=$(VERSION)"
GOFLAGS := -buildmode=pie -trimpath -mod=readonly -modcacherw -ldflags $(LDFLAGS)

all: $(BINARY)

$(BINARY): cmd/*.go
	go build -v -o $@ $(GOFLAGS) ./cmd/...

version:
	@echo $(VERSION)

install: $(BINARY) $(BINARY).conf
	mkdir -p $(PREFIX)/{etc/zabbix/zabbix_agentd.conf.d,usr/bin}
	cp $(BINARY)      $(PREFIX)/usr/bin/
	cp $(BINARY).conf $(PREFIX)/etc/zabbix/zabbix_agentd.conf.d/

clean:
	rm -f $(BINARY)
