package main

import (
	"encoding/json"
	"fmt"
)

func discovery(
	nodesStats *ElasticNodesStats,
	aggGroup string,
) error {
	discoveryData := make(map[string][]map[string]string)

	var discoveredItems []map[string]string

	if aggGroup != "None" {
		aggregateItem := make(map[string]string)
		aggregateItem["{#GROUPNAME}"] = aggGroup
		discoveredItems = append(discoveredItems, aggregateItem)
	}

	for _, nodeStats := range nodesStats.Nodes {

		for collectorsName := range nodeStats.JVM.GC.Collectors {
			discoveredItem := make(map[string]string)
			discoveredItem["{#JVMGCCOLLECTORS}"] = collectorsName
			discoveredItems = append(discoveredItems, discoveredItem)
		}

		for bufferPoolsName := range nodeStats.JVM.BufferPools {
			discoveredItem := make(map[string]string)
			discoveredItem["{#JVMBUFFERSPOOLS}"] = bufferPoolsName
			discoveredItems = append(discoveredItems, discoveredItem)
		}

		for poolsName := range nodeStats.JVM.Mem.Pools {
			discoveredItem := make(map[string]string)
			discoveredItem["{#JVMMEMPOOLS}"] = poolsName
			discoveredItems = append(discoveredItems, discoveredItem)
		}

		for threadPoolName := range nodeStats.ThreadPools {
			discoveredItem := make(map[string]string)
			discoveredItem["{#THREADPOOLNAME}"] = threadPoolName
			discoveredItems = append(discoveredItems, discoveredItem)
		}

	}

	discoveryData["data"] = discoveredItems

	out, err := json.Marshal(discoveryData)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", out)
	return nil
}
func discoveryIndices(
	indicesStats *ElasticIndicesStats,
) error {
	discoveryData := make(map[string][]map[string]string)

	var discoveredItems []map[string]string

	for name, _ := range indicesStats.Indices {
		discoveredItem := make(map[string]string)
		discoveredItem["{#INDEX}"] = name
		discoveredItems = append(discoveredItems, discoveredItem)
	}

	discoveryData["data"] = discoveredItems

	out, err := json.Marshal(discoveryData)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", out)
	return nil
}
