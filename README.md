# zabbix-agent-extension-elasticsearch

zabbix-agent-extension-elasticsearch - this extension for monitoring Elasticsearch cluster and node health/status.

### Supported features

This extension obtains stats of two types:

#### Node stat
https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-stats.html

- [ ] roles
- [ ] attributes
- [x] indices (partly)
- [ ] os
- [ ] processes
- [x] jvm
- [x] thread_pool
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
- [x] total indices docs count
- [x] total indices deleted docs count
- [x] primary indices docs count
- [x] primary indices deleted docs count
- [x] total indices store size
- [x] primary indices store size
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

./build.sh

#Installing
pacman -U *.tar.xz
```

### Dependencies

zabbix-agent-extension-elasticsearch requires [zabbix-agent](http://www.zabbix.com/download) v2.4+ to run.

### Zabbix configuration
In order to start getting metrics, it is enough to import template and attach it to monitored node.

`WARNING:` You must define macro with name - `{$ZABBIX_SERVER_IP}` in global or local (template) scope with IP address of  zabbix server.

On one node of cluster set MACRO `{$GROUPNAME}` = `REAL_ZABBIX_GROUP`. This group must include all nodes of the cluster.
Only this one node will be triggered cluster status (low level discovery added aggregate checks of cluster health).

### Customize key prefix
It may you need if key in template already used.

If you need change key `elasticsearch.*` -> `YOUR_PREFIX_PART.elasticsearch.*`, run script `custom_key_template.sh` whit `YOUR_PREFIX_PART` and import updated zabbix template `template_elasticsearch_service.xml`.

```sh
./custom_key_template.sh YOUR_PREFIX_PART
```

### Elasticsearch API authentication (X-Pack security)

This extension support basic authentication which provided by X-Pack. For authentication in Elasticsearch you must set valid values in template macros - `${ES_USER}` and `${ES_PASSWORD}`

### Customize Elasticsearch ip/port.

Just change `{$ES_ADDRESS}` macro in template.
