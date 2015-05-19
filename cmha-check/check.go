package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CheckCommandFactory struct {
}

func (c *CheckCommandFactory) Help() string {
	helpText := `
Usage: cmha check [options]

Options:

  -all		Query all
  -cluster		Query cluster 
  -chap		Query chap 
  -db		Query db
  -config		Query config file 
  -leader		Query leader
  -service	Query service
 `
	return strings.TrimSpace(helpText)
}

func (c *CheckCommandFactory) Run(args []string) int {
	var All bool
	var Cluster bool
	var Chap bool
	var Db bool
	var Config bool
	var Leader bool
	var Service bool
	cmdFlags := flag.NewFlagSet("check", flag.ContinueOnError)
	cmdFlags.BoolVar(&All, "all", false, "query all")
	cmdFlags.BoolVar(&Cluster, "cluster", false, "query cluster")
	cmdFlags.BoolVar(&Chap, "chap", false, "query chap")
	cmdFlags.BoolVar(&Db, "db", false, "query db")
	cmdFlags.BoolVar(&Config, "config", false, "query config")
	cmdFlags.BoolVar(&Leader, "leader", false, "query leader")
	cmdFlags.BoolVar(&Service, "service", false, "query service")

	err := cmdFlags.Parse(args)
	if err != nil {
		fmt.Println("Parse failure!")
	}
	if All {
		fmt.Println("all:", All)
	}
	if Cluster {
		fmt.Println("cluster:", Cluster)
		err := ClusterMembers()
		if err != nil {
			beego.Error("Query cluster mebers failure!", err)
		}
	}
	if Chap {
		fmt.Println("chap:", Chap)
	}
	if Db {
		fmt.Println("db:", Db)
	}
	if Config {
		fmt.Println("config:", Config)
	}
	if Leader {
		fmt.Println("leader:", Leader)
	}
	if Service {
		fmt.Println("service:", Service)
	}
	return 0
}

func (c *CheckCommandFactory) Synopsis() string {
	return "Runs a chma check"
}
