.PHONY: all verion test clean install

VERSION := $(shell git log -1 --format=%cd.$(shell git rev-list --count HEAD).%h --date=format:%y%m%d)
LDFLAGS := "-X main.version=$(VERSION)"
BINARY  := zabbix-agent-extension-elasticsearch

all: $(BINARY)

$(BINARY): *.go
	go build -v -o $@ -ldflags $(LDFLAGS)

version:
	@echo $(VERSION)

install: $(BINARY) $(BINARY).conf
	mkdir -p $(PREFIX){etc/zabbix/zabbix_agentd.conf.d,usr/bin}
	cp $(BINARY)      $(PREFIX)usr/bin/
	cp $(BINARY).conf $(PREFIX)etc/zabbix/zabbix_agentd.conf.d/

clean:
	rm -f $(BINARY)
