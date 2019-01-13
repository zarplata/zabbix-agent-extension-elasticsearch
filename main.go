package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"

	zsend "github.com/blacked/go-zabbix"
	docopt "github.com/docopt/docopt-go"
)

const (
	noneValue = "None"
)

var (
	version = "[manual build]"
	err     error
)

func main() {
	usage := `zabbix-agent-extension-elasticsearch

Usage:
  zabbix-agent-extension-elasticsearch [options]

Options:
  --type <type>                 Type of statistics: global (cluster and nodes)
                                  or indices [default: global].
  -e --elasticsearch <dsn>      DSN of Elasticsearch server
                                  [default: 127.0.0.1:9200].
  --agg-group <group>           Group name which will be use for aggregate
                                  item values [default: None].
  -u --user <name>              User for authenticate through 
                                  Elasticsearch API [default: None].
  -x --password <string>        Password for user [default: None].

Stats options:
  -z --zabbix <zabbix>          Hostname or IP address of zabbix server
                                  [default: 127.0.0.1].
  -p --port <port>              Port of zabbix server [default: 10051]
  --prefix <prefix>             Add part of your prefix for key
                                  [default: None_pfx].
  --hostname <hostname>         Override hostname used to identify in zabbix server
                                  [default: None].

Discovery options:
  --discovery                   Run low-level discovery for determine
                                  gc collectors, mem pools, boofer pools, etc.

Misc options:
  --version                     Show version.
  -h --help                     Show this screen.
`

	args, _ := docopt.Parse(usage, nil, true, version, false)

	elasticDSN := args["--elasticsearch"].(string)

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

	elasticsearchAuthToken := noneValue

	elasticsearchUser := args["--user"].(string)
	elasticsearchPassword := args["--password"].(string)

	if elasticsearchUser != noneValue && elasticsearchPassword != noneValue {
		elasticsearchAuthToken = base64.StdEncoding.EncodeToString(
			[]byte(
				fmt.Sprintf(
					"%s:%s",
					elasticsearchUser,
					elasticsearchPassword,
				),
			),
		)
	}

	hostname := args["--hostname"].(string)
	if hostname == "None" {
		hostname, err = os.Hostname()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	var metrics []*zsend.Metric

	statsType := args["--type"].(string)

	switch statsType {
	case "indices":
		if aggGroup == "None" {
			fmt.Println("indices work only master node with --agg-group set")
			os.Exit(0)
		}

		indicesStats, err := getIndicesStats(
			elasticDSN,
			elasticsearchAuthToken,
		)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if args["--discovery"].(bool) {
			err = discoveryIndices(indicesStats)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			os.Exit(0)
		}

		metrics = createIndicesStats(
			hostname,
			indicesStats,
			metrics,
			prefix,
		)

	case "global":
		clusterHealth, err := getClusterHealth(
			elasticDSN,
			elasticsearchAuthToken,
		)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		nodesStats, err := getNodeStats(
			elasticDSN,
			elasticsearchAuthToken,
		)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

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
			prefix,
		)
		metrics = createNodeStatsJVMMetrics(
			hostname,
			nodesStats,
			metrics,
			prefix,
		)

		metrics = createNodeStatsThreadPool(
			hostname,
			nodesStats,
			metrics,
			prefix,
		)

		metrics = createNodeStatsIndices(
			hostname,
			nodesStats,
			metrics,
			prefix,
		)

		metrics = createNodeStatsTransport(
			hostname,
			nodesStats,
			metrics,
			prefix,
		)

		metrics = createNodeStatsHttp(
			hostname,
			nodesStats,
			metrics,
			prefix,
		)

	default:
		fmt.Println("Unsupported type of stats.")
		os.Exit(0)
	}

	packet := zsend.NewPacket(metrics)
	sender := zsend.NewSender(
		zabbix,
		port,
	)
	sender.Send(packet)

	fmt.Println("OK")
}
