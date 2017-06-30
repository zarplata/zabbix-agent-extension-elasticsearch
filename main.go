package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	zsend "github.com/blacked/go-zabbix"
	docopt "github.com/docopt/docopt-go"
)

var version = "[manual build]"

func main() {
	usage := `zabbix-agent-extension-elastic

Usage:
  zabbix-agent-extension-elastic [-e --elasticsearch <dsn>] [-z --zabbix-host <zhost>] [-p --zabbix-port <zport>] [--zabbix-prefix <prefix>]
  zabbix-agent-extension-elastic [-e --elasticsearch <dsn>] [--discovery] [--agg-group <group>]
  zabbix-agent-extension-elastic [-h | --help]

Options:
	-e --elasticsearch <dsn>      DSN of Elasticsearch server [default: 127.0.0.1:9200]
	-z --zabbix-host <zhost>      Hostname or IP address of zabbix server [default: 127.0.0.1]
	-p --zabbix-port <zport>      Port of zabbix server [default: 10051]
	--zabbix-prefix <prefix>      Add part of your prefix for key [default: None_pfx]
	--discovery                   Run low-level discovery for determine gc collectors, mem pools, boofer pools, etc.
	--agg-group <group>           Group name which will be use for aggregate item values.[default: None]
	-h --help                     Show this screen.
`
	args, _ := docopt.Parse(usage, nil, true, version, false)
	elasticDSN := args["--elasticsearch"].(string)

	aggGroup := args["--agg-group"].(string)

	zabbixHost := args["--zabbix-host"].(string)
	zabbixPort, err := strconv.Atoi(args["--zabbix-port"].(string))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	zabbixPrefix := args["--zabbix-prefix"].(string)
	if zabbixPrefix != "None_pfx" {
		zabbixPrefix = strings.Join([]string{zabbixPrefix, "elasticsearch"}, ".")
	} else {
		zabbixPrefix = "elasticsearch"
	}

	clusterHealth, err := getClusterHealth(elasticDSN)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	nodesStats, err := getNodeStats(elasticDSN)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var metrics []*zsend.Metric

	if args["--discovery"].(bool) {
		err = discovery(nodesStats, aggGroup)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	metrics = createClusterHealthMetrics(
		hostname,
		clusterHealth,
		metrics,
		zabbixPrefix,
	)
	metrics = createNodeStatsJVMMetrics(
		hostname,
		nodesStats,
		metrics,
		zabbixPrefix,
	)

	packet := zsend.NewPacket(metrics)
	sender := zsend.NewSender(
		zabbixHost,
		zabbixPort,
	)
	sender.Send(packet)
	fmt.Println("OK")
}
