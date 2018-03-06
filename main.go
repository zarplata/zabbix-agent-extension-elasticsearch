package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	zsend "github.com/blacked/go-zabbix"
	docopt "github.com/docopt/docopt-go"
	lorg "github.com/kovetskiy/lorg"
)

var (
	version = "[manual build]"
	logger  *lorg.Log
	err     error
)

func main() {
	usage := `zabbix-agent-extension-elasticsearch

Usage:
  zabbix-agent-extension-elastic [options]

Options:
  --type <type>                 Type of statistics: global (cluster and nodes)
                                  or indices [default: global].
  -e --elasticsearch <dsn>      DSN of Elasticsearch server
                                  [default: 127.0.0.1:9200].
  --agg-group <group>           Group name which will be use for aggregate
                                  item values [default: None].

Stats options:
  -z --zabbix <zabbix>          Hostname or IP address of zabbix server
                                  [default: 127.0.0.1].
  -p --port <port>              Port of zabbix server [default: 10051]
  --prefix <prefix>             Add part of your prefix for key
                                  [default: None_pfx].

Discovery options:
  --discovery                   Run low-level discovery for determine
                                  gc collectors, mem pools, boofer pools, etc.

Misc options:
  -v --verbose                  Print verbose messages.
  --version                     Show version.
  -h --help                     Show this screen.
`

	start := time.Now()

	args, _ := docopt.Parse(usage, nil, true, version, false)

	mustSetupLogger(args["--verbose"].(bool))

	elasticDSN := args["--elasticsearch"].(string)
	logger.Debugf("set elasticsearch dsn %s", elasticDSN)

	aggGroup := args["--agg-group"].(string)

	zabbix := args["--zabbix"].(string)
	port, err := strconv.Atoi(args["--port"].(string))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	prefix := args["--prefix"].(string)
	if prefix != "None_pfx" {
		prefix = strings.Join([]string{prefix, "elasticsearch"}, ".")
	} else {
		prefix = "elasticsearch"
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var metrics []*zsend.Metric

	statsType := args["--type"].(string)

	switch statsType {
	case "indices":
		if aggGroup == "None" {
			fmt.Println("indices work only master node with --agg-group set")
			os.Exit(0)
		}

		logger.Debug("try get indices stats")
		indicesStats, err := getIndicesStats(elasticDSN)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if args["--discovery"].(bool) {
			logger.Debug("discovery indices")
			err = discoveryIndices(indicesStats)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			os.Exit(0)
		}

		logger.Debug("create metrics from indices stats")
		metrics = createIndicesStats(
			hostname,
			indicesStats,
			metrics,
			prefix,
		)
	default:
		logger.Debug("try get cluster health")
		clusterHealth, err := getClusterHealth(elasticDSN)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		logger.Debug("try get node stats")
		nodesStats, err := getNodeStats(elasticDSN)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if args["--discovery"].(bool) {
			logger.Debug("discovery node stats")
			err = discovery(nodesStats, aggGroup)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			os.Exit(0)
		}

		logger.Debug("create metrics from cluster health stats")
		metrics = createClusterHealthMetrics(
			hostname,
			clusterHealth,
			metrics,
			prefix,
		)
		logger.Debug("create metrics from node stats")
		metrics = createNodeStatsJVMMetrics(
			hostname,
			nodesStats,
			metrics,
			prefix,
		)
	}
	packet := zsend.NewPacket(metrics)
	sender := zsend.NewSender(
		zabbix,
		port,
	)
	logger.Debugf("send metrics to zabbix %s:%d", zabbix, port)
	sender.Send(packet)

	elapsed := time.Since(start)
	logger.Debugf("execution time %s", elapsed)

	fmt.Println("OK")
}
