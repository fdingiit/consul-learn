package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	conf := consulapi.DefaultConfig()
	cli, err := consulapi.NewClient(conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		hc, q, err := cli.Health().Checks("dummy-service", &consulapi.QueryOptions{
			UseCache:     true,
			MaxAge:       time.Hour,
			StaleIfError: time.Hour,
		})
		fmt.Printf("%+v, %+v, %+v\n", prettyPrint(hc), q, err)

		time.Sleep(time.Second)
	}
}

func prettyPrint(hcs consulapi.HealthChecks) string {
	var hcsStrs []string
	for _, hc := range hcs {
		hcsStrs = append(hcsStrs, fmt.Sprintf("%+v", hc))
	}
	return strings.Join(hcsStrs, ",")
}
