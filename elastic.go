package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/reconquest/hierr-go"
)

type ElasticClusterHealth struct {
	ClusterName                 string  `json:"cluster_name"`
	Status                      string  `json:"status"`
	TimedOut                    bool    `json:"timed_out"`
	NumderOfNodes               int64   `json:"number_of_nodes"`
	NumberOfDataNodes           int64   `json:"number_of_data_nodes"`
	ActivePrimaryShards         int64   `json:"active_primary_shards"`
	ActiveShards                int64   `json:"active_shards"`
	RelocatingShards            int64   `json:"relocating_shards"`
	InitializingShards          int64   `json:"initializing_shards"`
	UnassignedShards            int64   `json:"unassigned_shards"`
	DelayedUnassignedShards     int64   `json:"delayed_unassigned_shards"`
	NumberOfPendingTasks        int64   `json:"number_of_pending_tasks"`
	NumberOfInFlightFetch       int64   `json:"number_of_in_flight_fetch"`
	TaskMaxWaitingInQueueMillis int64   `json:"task_max_waiting_in_queue_millis"`
	ActiveShardsPercent         float64 `json:"active_shards_percent_as_number"`
}

type ElasticNodesStats struct {
	Nodes map[string]ElasticNodeStats `json:"nodes"`
}

type ElasticNodeStats struct {
	JVM         ElasticNodeStatsJVM       `json:"jvm"`
	ThreadPools map[string]NodeThreadPool `json:"thread_pool"`
	Indices     NodeIndices               `json:"indices"`
	Transport   ElasticNodeStatsTransport `json:"transport"`
	Http        ElasticNodeStatsHttp      `json:"http"`
}

type ElasticNodeStatsJVM struct {
	Timestamp      int64                                          `json:"timestamp"`
	UptimeInMillis int64                                          `json:"uptime_in_millis"`
	Mem            ElasticNodeStatsJVMMem                         `json:"mem"`
	Threads        ElasticNodeStatsJVMThreadsStats                `json:"threads"`
	GC             ElasticNodeStatsJVMGC                          `json:"gc"`
	BufferPools    map[string]ElasticNodeStatsJVMBufferPoolsStats `json:"buffer_pools"`
	Classes        ElasticNodeStatsJVMClassesStats                `json:"classes"`
}

type ElasticNodeStatsJVMMem struct {
	HeapUsedInBytes         int64                                       `json:"heap_used_in_bytes"`
	HeapUsedPercent         int64                                       `json:"heap_used_percent"`
	HeapCommittedInBytes    int64                                       `json:"heap_committed_in_bytes"`
	HeapMaxInBytes          int64                                       `json:"heap_max_in_bytes"`
	NonHeapUsedInBytes      int64                                       `json:"non_heap_used_in_bytes"`
	NonHeapCommittedInBytes int64                                       `json:"non_heap_committed_in_bytes"`
	Pools                   map[string]ElasticNodeStatsJVMMemPoolsStats `json:"pools"`
}

type ElasticNodeStatsJVMMemPoolsStats struct {
	UsedInBytes     int64 `json:"used_in_bytes"`
	MaxInBytes      int64 `json:"max_in_bytes"`
	PeakUsedInBytes int64 `json:"peak_used_in_bytes"`
	PeakMaxInBytes  int64 `json:"peak_max_in_bytes"`
}

type ElasticNodeStatsJVMThreadsStats struct {
	Count     int64 `json:"count"`
	PeakCount int64 `json:"peak_count"`
}

type ElasticNodeStatsJVMGC struct {
	Collectors map[string]ElasticNodeStatsJVMGCCollectorsStats `json:"collectors"`
}

type ElasticNodeStatsJVMGCCollectorsStats struct {
	CollectionCount        int64 `json:"collection_count"`
	CollectionTimeInMillis int64 `json:"collection_time_in_millis"`
}

type ElasticNodeStatsJVMBufferPoolsStats struct {
	Count                int64 `json:"count"`
	UsedInBytes          int64 `json:"used_in_bytes"`
	TotalCapacityInBytes int64 `json:"total_capacity_in_bytes"`
}

type ElasticNodeStatsJVMClassesStats struct {
	CurrentLoadedCount int64 `json:"current_loaded_count"`
	TotalLoadedCount   int64 `json:"total_loaded_count"`
	TotalUnloadedCount int64 `json:"total_unloaded_count"`
}

type ElasticNodeStatsTransport struct {
	ServerOpen    int64 `json:"server_open"`
	RxCount       int64 `json:"rx_count"`
	RxSizeInBytes int64 `json:"rx_size_in_bytes"`
	TxCount       int64 `json:"tx_count"`
	TxSizeInBytes int64 `json:"tx_size_in_bytes"`
}

type ElasticNodeStatsHttp struct {
	CurrentOpen int64 `json:"current_open"`
	TotalOpened int64 `json:"total_opened"`
}

type ElasticIndicesStats struct {
	Shards  ElasticIndicesStatsShards            `json:"_shards"`
	All     ElasticIndicesStatsAll               `json:"_all"`
	Indices map[string]ElasticIndicesStatsIndice `json:"indices"`
}

type ElasticIndicesStatsShards struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Failed     int64 `json:"failed"`
}

type ElasticIndicesStatsAll struct {
	Primaries ElasticIndicesStatsIndex `json:"primaries"`
	Total     ElasticIndicesStatsIndex `json:"total"`
}

type ElasticIndicesStatsIndice struct {
	Primaries ElasticIndicesStatsIndex `json:"primaries"`
	Total     ElasticIndicesStatsIndex `json:"total"`
}

type ElasticIndicesStatsIndex struct {
	Docs struct {
		Count   int64 `json:"count"`
		Deleted int64 `json:"deleted"`
	} `json:"docs"`
	Store struct {
		SizeInBytes          int64 `json:"size_in_bytes"`
		ThrottleTimeInMillis int64 `json:"throttle_time_in_millis"`
	} `json:"store"`
}

func getClusterHealth(
	elasticDSN string,
	elasticsearchAuthToken string,
) (*ElasticClusterHealth, error) {

	client := &http.Client{}

	var elasticClusterHealth ElasticClusterHealth

	clutserHealthURL := fmt.Sprintf("http://%s/_cluster/health", elasticDSN)
	request, err := http.NewRequest("GET", clutserHealthURL, nil)
	if err != nil {
		return nil, hierr.Errorf(
			err,
			"can`t create new HTTP request to %s",
			elasticDSN,
		)
	}

	if elasticsearchAuthToken != noneValue {
		request.Header.Add("Authorization", "Basic "+elasticsearchAuthToken)
	}

	clusterHealthResponse, err := client.Do(request)
	if err != nil {
		return nil, hierr.Errorf(
			err.Error(),
			"can`t get cluster health from Elasticsearch %s",
			elasticDSN,
		)
	}

	defer clusterHealthResponse.Body.Close()

	if clusterHealthResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"can`t get cluster health, Elasticsearch cluster returned %d HTTP code, expected %d HTTP code",
			clusterHealthResponse.StatusCode,
			http.StatusOK,
		)
	}

	err = json.NewDecoder(clusterHealthResponse.Body).Decode(&elasticClusterHealth)
	if err != nil {
		return nil, hierr.Errorf(
			err.Error(),
			"can`t decode cluster health response from Elasticsearch %s",
			elasticDSN,
		)
	}

	return &elasticClusterHealth, nil
}

func getNodeStats(
	elasticDSN string,
	elasticsearchAuthToken string,
) (*ElasticNodesStats, error) {

	client := &http.Client{}

	var elasticNodesStats ElasticNodesStats

	nodeStatsURL := fmt.Sprintf("http://%s/_nodes/_local/stats", elasticDSN)
	request, err := http.NewRequest("GET", nodeStatsURL, nil)
	if err != nil {
		return nil, hierr.Errorf(
			err,
			"can`t create new HTTP request to %s",
			elasticDSN,
		)
	}

	if elasticsearchAuthToken != noneValue {
		request.Header.Add("Authorization", "Basic "+elasticsearchAuthToken)
	}

	nodeStatsResponse, err := client.Do(request)
	if err != nil {
		return nil, hierr.Errorf(
			err.Error(),
			"can`t get node stats from Elasticsearch %s",
			elasticDSN,
		)
	}

	defer nodeStatsResponse.Body.Close()

	if nodeStatsResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"can`t get node stats, Elasticsearch node returned %d HTTP code",
			nodeStatsResponse.StatusCode,
			http.StatusOK,
		)
	}

	err = json.NewDecoder(nodeStatsResponse.Body).Decode(&elasticNodesStats)
	if err != nil {
		return nil, hierr.Errorf(
			err.Error(),
			"can`t decode node stats response from Elasticsearch %s",
			elasticDSN,
		)
	}

	return &elasticNodesStats, nil
}

func getIndicesStats(
	elasticDSN string,
	elasticsearchAuthToken string,
) (*ElasticIndicesStats, error) {

	client := &http.Client{}

	var elasticIndicesStats ElasticIndicesStats

	indicesStatsURL := fmt.Sprintf("http://%s/_stats", elasticDSN)
	request, err := http.NewRequest("GET", indicesStatsURL, nil)
	if err != nil {
		return nil, hierr.Errorf(
			err,
			"can`t create new HTTP request to %s",
			elasticDSN,
		)
	}

	if elasticsearchAuthToken != noneValue {
		request.Header.Add("Authorization", "Basic "+elasticsearchAuthToken)
	}

	indicesStatsResponse, err := client.Do(request)
	if err != nil {
		return nil, hierr.Errorf(
			err.Error(),
			"can`t get indices stats from Elasticsearch %s",
			elasticDSN,
		)
	}

	defer indicesStatsResponse.Body.Close()

	if indicesStatsResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"can`t get indices stats, Elasticsearch node returned %d HTTP code, expected %d HTTP code",
			indicesStatsResponse.StatusCode,
			http.StatusOK,
		)
	}

	err = json.NewDecoder(indicesStatsResponse.Body).Decode(&elasticIndicesStats)
	if err != nil {
		return nil, hierr.Errorf(
			err.Error(),
			"can`t decode indices stats response from Elasticsearch %s",
			elasticDSN,
		)
	}

	return &elasticIndicesStats, nil
}
