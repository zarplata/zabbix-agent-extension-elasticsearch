package main

import (
	"fmt"
	"strconv"

	zsend "github.com/blacked/go-zabbix"
)

func makePrefix(prefix, key string) string {
	return fmt.Sprintf(
		"%s.%s", prefix, key,
	)

}

func createClusterHealthMetrics(
	hostname string,
	clusterHealth *ElasticClusterHealth,
	metrics []*zsend.Metric,
	prefix string,
) []*zsend.Metric {

	healthToInt := make(map[string]int)
	healthToInt["green"] = 0
	healthToInt["yellow"] = 1
	healthToInt["red"] = 2

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.cluster_name",
			),
			clusterHealth.ClusterName,
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.status_int",
			),
			strconv.Itoa(int(healthToInt[clusterHealth.Status])),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.status",
			),
			clusterHealth.Status,
		),
	)

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.timed_out",
			),
			strconv.FormatBool(clusterHealth.TimedOut),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.number_of_nodes",
			),
			strconv.Itoa(int(clusterHealth.NumderOfNodes)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.number_of_data_nodes",
			),
			strconv.Itoa(int(clusterHealth.NumberOfDataNodes)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.active_primary_shards",
			),
			strconv.Itoa(int(clusterHealth.ActivePrimaryShards)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.active_shards",
			),
			strconv.Itoa(int(clusterHealth.ActiveShards)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.relocating_shards",
			),
			strconv.Itoa(int(clusterHealth.RelocatingShards)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.initializing_shards",
			),
			strconv.Itoa(int(clusterHealth.InitializingShards)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.unassigned_shards",
			),
			strconv.Itoa(int(clusterHealth.UnassignedShards)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.delayed_unassigned_shards",
			),
			strconv.Itoa(int(clusterHealth.DelayedUnassignedShards)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.number_of_pending_tasks",
			),
			strconv.Itoa(int(clusterHealth.NumberOfPendingTasks)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.number_of_in_flight_fetch",
			),
			strconv.Itoa(int(clusterHealth.NumberOfInFlightFetch)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.task_max_waiting_in_queue_millis",
			),
			strconv.Itoa(int(clusterHealth.TaskMaxWaitingInQueueMillis)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"cluster_health.active_shards_percent",
			),
			strconv.Itoa(int(clusterHealth.ActiveShardsPercent)),
		),
	)

	return metrics
}

func createNodeStatsTransport(
	hostname string,
	nodesStats *ElasticNodesStats,
	metrics []*zsend.Metric,
	prefix string,
) []*zsend.Metric {

	for _, nodeStats := range nodesStats.Nodes {
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					"node_stats.transport.server_open",
				),
				strconv.Itoa(int(nodeStats.Transport.ServerOpen)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					"node_stats.transport.rx_count",
				),
				strconv.Itoa(int(nodeStats.Transport.RxCount)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					"node_stats.transport.rx_size_in_bytes",
				),
				strconv.Itoa(int(nodeStats.Transport.RxSizeInBytes)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					"node_stats.transport.tx_count",
				),
				strconv.Itoa(int(nodeStats.Transport.TxCount)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					"node_stats.transport.tx_size_in_bytes",
				),
				strconv.Itoa(int(nodeStats.Transport.TxSizeInBytes)),
			),
		)
	}

	return metrics
}

func createNodeStatsHttp(
	hostname string,
	nodesStats *ElasticNodesStats,
	metrics []*zsend.Metric,
	prefix string,
) []*zsend.Metric {

	for _, nodeStats := range nodesStats.Nodes {
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					"node_stats.http.current_open",
				),
				strconv.Itoa(int(nodeStats.Http.CurrentOpen)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					"node_stats.http.total_opened",
				),
				strconv.Itoa(int(nodeStats.Http.TotalOpened)),
			),
		)
	}

	return metrics
}

func createNodeStatsJVMMetrics(
	hostname string,
	nodesStats *ElasticNodesStats,
	metrics []*zsend.Metric,
	prefix string,
) []*zsend.Metric {

	for _, nodeStats := range nodesStats.Nodes {
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					"node_stats.jvm.timestamp",
				),
				strconv.Itoa(int(nodeStats.JVM.Timestamp)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					"node_stats.jvm.uptime_in_millis",
				),
				strconv.Itoa(int(nodeStats.JVM.UptimeInMillis)),
			),
		)

		metrics = createNodeStatsJVMMemMetrics(hostname, metrics, &nodeStats, prefix)
		metrics = createNodeStatsJVMThreadsMetrics(hostname, metrics, &nodeStats, prefix)

		for collectorsName, nodeStatsJVMGCColletorsStats := range nodeStats.JVM.GC.Collectors {
			metrics = createNodeStatsJVMGCCollectorsMetrics(hostname, metrics, &nodeStatsJVMGCColletorsStats, collectorsName, prefix)
		}

		for bufferPoolsName, nodeStatsJVMBufferPoolsStats := range nodeStats.JVM.BufferPools {
			metrics = createNodeStatsJVMBufferPoolsMetrics(hostname, metrics, &nodeStatsJVMBufferPoolsStats, bufferPoolsName, prefix)
		}
		metrics = createNodeStatsJVMClassesMetrics(hostname, metrics, &nodeStats, prefix)
	}

	return metrics
}

func createNodeStatsJVMMemMetrics(
	hostname string,
	metrics []*zsend.Metric,
	nodeStats *ElasticNodeStats,
	prefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.jvm.mem.heap_used_in_bytes",
			),
			strconv.Itoa(int(nodeStats.JVM.Mem.HeapUsedInBytes)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.jvm.mem.heap_used_percent",
			),
			strconv.Itoa(int(nodeStats.JVM.Mem.HeapUsedPercent)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.jvm.mem.heap_committed_in_bytes",
			),
			strconv.Itoa(int(nodeStats.JVM.Mem.NonHeapCommittedInBytes)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.jvm.mem.heap_max_in_bytes",
			),
			strconv.Itoa(int(nodeStats.JVM.Mem.HeapMaxInBytes)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.jvm.mem.non_heap_used_in_bytes",
			),
			strconv.Itoa(int(nodeStats.JVM.Mem.NonHeapUsedInBytes)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.jvm.mem.non_heap_committed_in_bytes",
			),
			strconv.Itoa(int(nodeStats.JVM.Mem.NonHeapCommittedInBytes)),
		),
	)

	for poolsName, nodeStatsJVMMemPoolsStats := range nodeStats.JVM.Mem.Pools {
		metrics = createNodeStatsJVMMemPoolsMetrics(hostname, metrics, &nodeStatsJVMMemPoolsStats, poolsName, prefix)
	}

	return metrics
}

func createNodeStatsJVMMemPoolsMetrics(
	hostname string,
	metrics []*zsend.Metric,
	nodeStatsJVMMemPoolsStats *ElasticNodeStatsJVMMemPoolsStats,
	poolsName string,
	prefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				fmt.Sprintf("node_stats.jvm.mem.pools.used_in_bytes.[%s]", poolsName),
			),
			strconv.Itoa(int(nodeStatsJVMMemPoolsStats.UsedInBytes)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				fmt.Sprintf("node_stats.jvm.mem.pools.max_in_bytes.[%s]", poolsName),
			),
			strconv.Itoa(int(nodeStatsJVMMemPoolsStats.MaxInBytes)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				fmt.Sprintf("node_stats.jvm.mem.pools.peak_used_in_bytes.[%s]", poolsName),
			),
			strconv.Itoa(int(nodeStatsJVMMemPoolsStats.PeakUsedInBytes)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				fmt.Sprintf("node_stats.jvm.mem.pools.peak_max_in_bytes.[%s]", poolsName),
			),
			strconv.Itoa(int(nodeStatsJVMMemPoolsStats.PeakMaxInBytes)),
		),
	)

	return metrics
}

func createNodeStatsJVMThreadsMetrics(
	hostname string,
	metrics []*zsend.Metric,
	nodeStats *ElasticNodeStats,
	prefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.jvm.threads.count",
			),
			strconv.Itoa(int(nodeStats.JVM.Threads.Count)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.jvm.threads.peak_count",
			),
			strconv.Itoa(int(nodeStats.JVM.Threads.PeakCount)),
		),
	)

	return metrics
}

func createNodeStatsJVMGCCollectorsMetrics(
	hostname string,
	metrics []*zsend.Metric,
	nodeStatsJVMGCColletorsStats *ElasticNodeStatsJVMGCCollectorsStats,
	collectorsName string,
	prefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				fmt.Sprintf("node_stats.jvm.gc.collectors.collection_cout.[%s]", collectorsName),
			),
			strconv.Itoa(int(nodeStatsJVMGCColletorsStats.CollectionCount)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				fmt.Sprintf("node_stats.jvm.gc.collectors.collection_time_in_millis.[%s]", collectorsName),
			),
			strconv.Itoa(int(nodeStatsJVMGCColletorsStats.CollectionTimeInMillis)),
		),
	)

	return metrics
}

func createNodeStatsJVMBufferPoolsMetrics(
	hostname string,
	metrics []*zsend.Metric,
	nodeStatsJVMBufferPoolsStats *ElasticNodeStatsJVMBufferPoolsStats,
	bufferPoolsName string,
	prefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				fmt.Sprintf("node_stats.jvm.buffer_polls.count.[%s]", bufferPoolsName),
			),
			strconv.Itoa(int(nodeStatsJVMBufferPoolsStats.Count)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				fmt.Sprintf("node_stats.jvm.buffer_polls.used_in_bytes.[%s]", bufferPoolsName),
			),
			strconv.Itoa(int(nodeStatsJVMBufferPoolsStats.UsedInBytes)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				fmt.Sprintf("node_stats.jvm.buffer_polls.total_capacity_in_bytes.[%s]", bufferPoolsName),
			),
			strconv.Itoa(int(nodeStatsJVMBufferPoolsStats.TotalCapacityInBytes)),
		),
	)

	return metrics
}

func createNodeStatsJVMClassesMetrics(
	hostname string,
	metrics []*zsend.Metric,
	nodeStats *ElasticNodeStats,
	prefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.jvm.classes.current_loaded_count",
			),
			strconv.Itoa(int(nodeStats.JVM.Classes.CurrentLoadedCount)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.jvm.classes.total_loaded_count",
			),
			strconv.Itoa(int(nodeStats.JVM.Classes.TotalLoadedCount)),
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				prefix,
				"node_stats.jvm.classes.total_unloaded_count",
			),
			strconv.Itoa(int(nodeStats.JVM.Classes.TotalUnloadedCount)),
		),
	)

	return metrics
}

func createIndicesStats(
	hostname string,
	indicesStats *ElasticIndicesStats,
	metrics []*zsend.Metric,
	prefix string,
) []*zsend.Metric {

	for indexName, indexStats := range indicesStats.Indices {

		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf(
						"indices_stats.total.docs.count.[%s]",
						indexName,
					),
				),
				strconv.Itoa(int(indexStats.Total.Docs.Count)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf(
						"indices_stats.total.docs.deleted.[%s]",
						indexName,
					),
				),
				strconv.Itoa(int(indexStats.Total.Docs.Deleted)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf(
						"indices_stats.primaries.docs.count.[%s]",
						indexName,
					),
				),
				strconv.Itoa(int(indexStats.Primaries.Docs.Count)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf(
						"indices_stats.primaries.docs.deleted.[%s]",
						indexName,
					),
				),
				strconv.Itoa(int(indexStats.Primaries.Docs.Deleted)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf(
						"indices_stats.total.store.size.[%s]",
						indexName,
					),
				),
				strconv.Itoa(int(indexStats.Total.Store.SizeInBytes)),
			),
		)
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf(
						"indices_stats.primaries.store.size.[%s]",
						indexName,
					),
				),
				strconv.Itoa(int(indexStats.Primaries.Store.SizeInBytes)),
			),
		)

	}

	return metrics
}
