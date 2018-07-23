package main

import (
	"fmt"
	"strconv"

	zsend "github.com/blacked/go-zabbix"
)

type NodeThreadPool struct {
	Threads   int64 `json:"threads"`
	Queue     int64 `json:"queue"`
	Active    int64 `json:"active"`
	Rejected  int64 `json:"rejected"`
	Largest   int64 `json:"largest"`
	Completed int64 `json:"completed"`
}

func createNodeStatsThreadPool(
	hostname string,
	nodesStats *ElasticNodesStats,
	metrics []*zsend.Metric,
	prefix string,
) []*zsend.Metric {

	var nodeStats ElasticNodeStats

	for _, nodeStat := range nodesStats.Nodes {
		nodeStats = nodeStat
		break
	}

	for threadPoolName, threadPoolMetric := range nodeStats.ThreadPools {
		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf(
						"node_stats.thread_pool.threads[%s]",
						threadPoolName,
					),
				),
				strconv.Itoa(int(threadPoolMetric.Threads)),
			),
		)

		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf(
						"node_stats.thread_pool.queue[%s]",
						threadPoolName,
					),
				),
				strconv.Itoa(int(threadPoolMetric.Queue)),
			),
		)

		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf(
						"node_stats.thread_pool.active[%s]",
						threadPoolName,
					),
				),
				strconv.Itoa(int(threadPoolMetric.Active)),
			),
		)

		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf(
						"node_stats.thread_pool.rejected[%s]",
						threadPoolName,
					),
				),
				strconv.Itoa(int(threadPoolMetric.Rejected)),
			),
		)

		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf(
						"node_stats.thread_pool.largest[%s]",
						threadPoolName,
					),
				),
				strconv.Itoa(int(threadPoolMetric.Largest)),
			),
		)

		metrics = append(
			metrics,
			zsend.NewMetric(
				hostname,
				makePrefix(
					prefix,
					fmt.Sprintf(
						"node_stats.thread_pool.completed[%s]",
						threadPoolName,
					),
				),
				strconv.Itoa(int(threadPoolMetric.Completed)),
			),
		)

	}

	return metrics
}
