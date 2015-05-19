package main

import (
	"fmt"
	"strings"
	//	"fmt"
	consulapi "github.com/armon/consul-api"
	"github.com/astaxie/beego"
)

func ClusterMembers() error {
	config := &consulapi.Config{
		Datacenter: beego.AppConfig.String("cmha-datacenter"),
		Token:      beego.AppConfig.String("cmha-token"),
		Address:    beego.AppConfig.String("cmha-server-ip") + ":8500",
	}

	fmt.Println(beego.AppConfig.String("cmha-server-ip"))
	client, err := consulapi.NewClient(config)
	if err != nil {
		beego.Error("Create consul-api client failure!", err)
		return err
	}
	beego.Info(" Create consul-api client success!")
	status := client.Status()
	peers, err := status.Peers()
	if err != nil {
		beego.Error("Query consul service faiiure!", err)
		return err
	}
	agent := client.Agent()
	members, err := agent.Members(false)
	if err != nil {
		beego.Error("Query cluster members failure", err)
		return err
	}
	//	peerssize := len(peers)
	for _, peersvalue := range peers {
		fmt.Println(strings.Split(peersvalue, ":"))
		for _, membersvalue := range members {
			fmt.Println(membersvalue.Addr)
		}
	}

	return nil
}
