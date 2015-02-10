package main

import (
	consulapi "github.com/armon/consul-api"
	"github.com/astaxie/beego"
	"time"
)

func SetConn() {
	//Config is used to configure the creation of a client
	config := &consulapi.Config{
		Datacenter: beego.AppConfig.String("datacenter"),
		Token:      beego.AppConfig.String("token"),
		Address:    beego.AppConfig.String("service_ip") + ":" + beego.AppConfig.String("service_port"),
	}
	hostname := beego.AppConfig.String("hostname")
	ip := beego.AppConfig.String("ip")
	port := beego.AppConfig.String("port")
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	servicename := beego.AppConfig.String("servicename")
	//NewClient returns a new client
	client, err := consulapi.NewClient(config)
	if err != nil {
		beego.Error("Create a consul-api client fails", err)
		return
	}
	//KV is used to return a handle to the K/V apis
	kv := client.KV()
	//Get is used to lookup a single key
	qm, _, err := kv.Get("service/"+servicename+"/leader", nil)
	if err != nil {
		beego.Error("Get a key failure", err)
		return
	}
	//Are there external connection string provided
	if qm.Session != "" {
		beego.Info("There are external connection string provided")
		time.Sleep(1000)
		return
	}
	//Health returns a handle to the health endpoints
	health := client.Health()
	//Checks is used to return the checks associated with a service
	healthvalue, _, err := health.Checks(servicename, nil)
	if err != nil {
		beego.Error("Return to service-related checks fail", err)
		return
	}
	if len(healthvalue) <= 0 {
		beego.Info("Without this service, or service is not a healthy state")
		return
	}
	var islocal bool
	for index := range healthvalue {
		if healthvalue[index].Node == hostname {
			islocal = true
			beego.Info("Native " + servicename + " service is healthy")
			break
		}

	}
	if !islocal {
		beego.Info("Native " + servicename + " service unhealthy or the service does not exist")
		return
	} else {
		//Session returns a handle to the session endpoints
		session := client.Session()
		sessionEntry := consulapi.SessionEntry{
			LockDelay: 15 * time.Second,
			Name:      servicename,
			Node:      hostname,
			Checks:    []string{"serfHealth", "service:" + servicename},
		}
		//Create makes a new session. Providing a session entry can customize the session. It can also be nil to use defaults.
		sessionvalue, _, err := session.Create(&sessionEntry, nil)
		if err != nil {
			beego.Error("Session creation failed", err)
			return
		}
		acquirejson := `{"Node":"` + hostname + `","IP":"` + ip + `","Port":` + port + `,"username":"` + username + `","password":" ` + password + `"}`
		value := []byte(acquirejson)
		kvpair := consulapi.KVPair{
			Key:     "service/" + servicename + "/leader",
			Value:   value,
			Session: sessionvalue,
		}
		//Acquire is used for a lock acquisiiton operation. The Key, Flags, Value and Session are respected. Returns true on success or false on failures.
		_, _, err = kv.Acquire(&kvpair, nil)
		if err != nil {
			beego.Error("Set the connection string master fails", err)
			return
		}
	}
}