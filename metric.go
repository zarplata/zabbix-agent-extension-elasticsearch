package main

import (
	"fmt"
	"strconv"

	zsend "github.com/blacked/go-zabbix"
)

func makePrefix(zabbixPrefix, key string) string {
	return fmt.Sprintf(
		"%s.%s", zabbixPrefix, key,
	)

}

func createClusterHealthMetrics(
	hostname string,
	clusterHealth *ElasticClusterHealth,
	metrics []*zsend.Metric,
	zabbixPrefix string,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
				"cluster_health.active_shards_percent",
			),
			strconv.Itoa(int(clusterHealth.ActiveShardsPercent)),
		),
	)

	return metrics
}

func createNodeStatsJVMMetrics(
	hostname string,
	nodesStats *ElasticNodesStats,
	metrics []*zsend.Metric,
	zabbixPrefix string,
) []*zsend.Metric {

	for _, nodeStats := range nodesStats.Nodes {
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					zabbixPrefix,
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
					zabbixPrefix,
					"node_stats.jvm.uptime_in_millis",
				),
				strconv.Itoa(int(nodeStats.JVM.UptimeInMillis)),
			),
		)

		metrics = createNodeStatsJVMMemMetrics(hostname, metrics, &nodeStats, zabbixPrefix)
		metrics = createNodeStatsJVMThreadsMetrics(hostname, metrics, &nodeStats, zabbixPrefix)

		for collectorsName, nodeStatsJVMGCColletorsStats := range nodeStats.JVM.GC.Collectors {
			metrics = createNodeStatsJVMGCCollectorsMetrics(hostname, metrics, &nodeStatsJVMGCColletorsStats, collectorsName, zabbixPrefix)
		}

		for bufferPoolsName, nodeStatsJVMBufferPoolsStats := range nodeStats.JVM.BufferPools {
			metrics = createNodeStatsJVMBufferPoolsMetrics(hostname, metrics, &nodeStatsJVMBufferPoolsStats, bufferPoolsName, zabbixPrefix)
		}
		metrics = createNodeStatsJVMClassesMetrics(hostname, metrics, &nodeStats, zabbixPrefix)
	}

	return metrics
}

func createNodeStatsJVMMemMetrics(
	hostname string,
	metrics []*zsend.Metric,
	nodeStats *ElasticNodeStats,
	zabbixPrefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
				"node_stats.jvm.mem.non_heap_committed_in_bytes",
			),
			strconv.Itoa(int(nodeStats.JVM.Mem.NonHeapCommittedInBytes)),
		),
	)

	for poolsName, nodeStatsJVMMemPoolsStats := range nodeStats.JVM.Mem.Pools {
		metrics = createNodeStatsJVMMemPoolsMetrics(hostname, metrics, &nodeStatsJVMMemPoolsStats, poolsName, zabbixPrefix)
	}

	return metrics
}

func createNodeStatsJVMMemPoolsMetrics(
	hostname string,
	metrics []*zsend.Metric,
	nodeStatsJVMMemPoolsStats *ElasticNodeStatsJVMMemPoolsStats,
	poolsName string,
	zabbixPrefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
	zabbixPrefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				zabbixPrefix,
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
				zabbixPrefix,
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
	zabbixPrefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				zabbixPrefix,
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
				zabbixPrefix,
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
	zabbixPrefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
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
	zabbixPrefix string,
) []*zsend.Metric {

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				zabbixPrefix,
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
				zabbixPrefix,
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
				zabbixPrefix,
				"node_stats.jvm.classes.total_unloaded_count",
			),
			strconv.Itoa(int(nodeStats.JVM.Classes.TotalUnloadedCount)),
		),
	)

	return metrics
}
