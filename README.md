# zabbix-agent-extension-elasticsearch

zabbix-agent-extension-elasticsearch - this extension for monitoring Elasticsearch cluster and node health/status.

### Description

This extension monitored:
[x] Get stats [trigger]
[x] process is NOT running [trigger]

#### Elasticsearch cluster:
[x] health
[x] health (integer -> aggregate health for all cluster nodes)
[x] name
[x] number of nodes
[x] number of data nodes
[x] number of in flight fetch
[x] number of pending tasks
[x] active primary shards
[x] active shards
[x] active shards percent
[x] initializing shards
[x] relocating shards
[x] unassigned shards
[x] delayed unassigned shards
[x] task max waiting in queue
[x] timeout

#### Elasticsearch node
[x] jvm classes current loaded count
[x] jvm classes total loaded count
[x] jvm classes total unloaded count
[x] jvm mem heap committed in bytes
[x] jvm mem heap max in bytes
[x] jvm mem heap used in bytes
[x] jvm mem heap used percent [trigger]
[x] jvm mem non heap committed in bytes
[x] jvm mem non heap used in bytes
[x] jvm threads count
[x] jvm threads peak count
[x] jvm timestamp
[x] jvm uptime

### Discovery
[x] JVM buffer pools {#JVMBUFFERSPOOLS} count
[x] JVM buffer pools {#JVMBUFFERSPOOLS} total_capacity_in_bytes
[x] JVM buffer pools {#JVMBUFFERSPOOLS} used_in_bytes
[x] JVM gc collectors {#JVMGCCOLLECTORS} collection count
[x] JVM gc collectors {#JVMGCCOLLECTORS} collection time
[x] JVM mem pools {#JVMMEMPOOLS} max in bytes
[x] JVM mem pools {#JVMMEMPOOLS} peak max in bytes
[x] JVM mem pools {#JVMMEMPOOLS} peak used in bytes
[x] JVM mem pools {#JVMMEMPOOLS} used in bytes [trigger]

#### Discovery (aggregate main node)
[x] Aggregate cluster active shards percent [trigger]
[x] Aggregate cluster delayed unassigned shards [trigger]
[x] Aggregate cluster health [trigger]
[x] Aggregate cluster unassigned shards [trigger]

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
