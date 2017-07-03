# zabbix-agent-extension-elasticsearch

zabbix-agent-extension-elasticsearch - this extension for monitoring Elasticsearch cluster and node health/status.

### Supported features

This extension obtains stats of two types:

#### Node stat
https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-stats.html

_This version supports only JVM stats, because now we are only interested in this metric._

- [ ] roles
- [ ] attributes
- [ ] indices
- [ ] os
- [ ] processes
- [x] jvm
- [ ] thread_pool
- [ ] fs
- [ ] transport
- [ ] http
- [ ] breakers
- [ ] script
- [ ] discovery
- [ ] ingest

#### Cluster health
https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html
- [x] cluster_name
- [x] status
- [x] timed_out
- [x] number_of_nodes
- [x] number_of_data_nodes
- [x] active_primary_shards
- [x] active_shards
- [x] relocating_shards
- [x] initializing_shards
- [x] unassigned_shards
- [x] delayed_unassigned_shards
- [x] number_of_pending_tasks
- [x] number_of_in_flight_fetch
- [x] task_max_waiting_in_queue_millis
- [x] active_shards_percent_as_number

### Installation

#### Manual build

```sh
# Building
git clone https://github.com/zarplata/zabbix-agent-extension-elasticsearch.git
cd zabbix-agent-extension-elasticsearch
make

#Installing
make install

# By default, binary installs into /usr/bin/ and zabbix config in /etc/zabbix/zabbix_agentd.conf.d/ but,
# you may manually copy binary to your executable path and zabbix config to specific include directory
```

#### Arch Linux package
```sh
# Building
git clone https://github.com/zarplata/zabbix-agent-extension-elasticsearch.git
git checkout pkgbuild

makepkg

#Installing
pacman -U *.tar.xz
```

### Dependencies

zabbix-agent-extension-elasticsearch requires [zabbix-agent](http://www.zabbix.com/download) v2.4+ to run.

### Zabbix configuration
In order to start getting metrics, it is enough to import template and attach it to monitored node.

`WARNING:` You must define macro with name - `{$ZABBIX_SERVER_IP}` in global or local (template) scope with IP address of  zabbix server.

### Customize key prefix
It may you need if key in template already used.

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
