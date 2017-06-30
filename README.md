# zabbix-agent-extension-elasticsearch

zabbix-agent-extension-elasticsearch - this extension for monitoring Elasticsearch cluster and node health/status.

### Tech

Binary `/usr/bin/zabbix-agent-extension-elasticsearch`

Zabbix-agent config `/etc/zabbix/zabbix_agentd.conf.d/zabbix-agent-extension-elasticsearch.conf`

Zabbix template `template_elasticsearch_service.xml`

Add global or local (template) MARCO {$ZABBIX_SERVER_IP} with your zabbix-server IP.

Restart zabbix-agent after install extension.

### Installation

```sh
git clone ssh://git@git.rn/devops/zabbix-agent-extension-elasticsearch.git
cd zabbix-agent-extension-elasticsearch
make
make install
```

### Deletion

```
make remove
```

zabbix-agent-extension-elasticsearch requires [zabbix-agent](http://www.zabbix.com/download) v2.4+ to run.

### Customize key prefix

Change key `elasticsearch.*` -> `service.elasticsearch.*`:

Replace `template_elasticsearch_service.xml` whit your prefix:

```sh
sed 's/elasticsearch./service.elasticsearch./g' -i template_elasticsearch_service.xml
sed 's/None_pfx/service/g' -i template_elasticsearch_service.xml
```

### Customize Elasticsearch ip/port, zabbix-server port

Default 127.0.0.1:9200

Add new param -e $3 (--elasticsearch $3) in `/etc/zabbix/zabbix_agentd.conf.d/zabbix-agent-extension-elasticsearch.conf` and add new param in key Item `Elasticsearch get stats`:

```sh
elasticsearch.stats[{$ZABBIX_SERVER_IP},{$ES_ZBX_PREFIX},ESIP:ESPORT]
```
