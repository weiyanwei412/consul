package main

import (
	consulapi "github.com/armon/consul-api"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func SessionAndChecks() {
	config := &consulapi.Config{
		Datacenter: beego.AppConfig.String("datacenter"),
		Token:      beego.AppConfig.String("token"),
		Address:    beego.AppConfig.String("service_ip") + ":" + beego.AppConfig.String("service_port"),
	}
	hostname := beego.AppConfig.String("hostname")
	servicename := beego.AppConfig.String("servicename")
	//NewClient returns a new client
	client, err := consulapi.NewClient(config)
	if err != nil {
		beego.Error("Create consul-api client failure!", err)
		return
	}
	beego.Info(" Create consul-api client success!")
	//KV is used to return a handle to the K/V apis
	kv := client.KV()
	//Get is used to lookup a single key
	kvPair, _, err := kv.Get("service/"+servicename+"/leader", nil)
	if err != nil {
		beego.Error("Get service/"+servicename+"/leader key failure!", err)
		return
	}
	beego.Info("Get service/" + servicename + "/leader key success!")
	if kvPair == nil {
		beego.Error("service/" + servicename + "/leader not found, Please create!")
		return
	}

	//Are there external connection string provided
	if kvPair.Session != "" {
		beego.Info("service/" + servicename + "/leader Connection string  set up,Leader presence!")
		time.Sleep(1 * time.Second)
		return
	}
	beego.Info("service/" + servicename + "/leader Connection string not set,Leader does not exist!")
	//Health returns a handle to the health endpoints
	health := client.Health()
	//Checks is used to return the checks associated with a service
	healthvalue, _, err := health.Checks(servicename, nil)
	if err != nil {
		beego.Error("Carried out health check /v1/health/checks/"+servicename+" failure!", err)
		return
	}
	beego.Error("Carried out health check /v1/health/checks/" + servicename + "success!")
	if len(healthvalue) <= 0 {
		beego.Info("No " + servicename + "service!")
		return
	}
	var islocal bool
	for index := range healthvalue {
		if healthvalue[index].Node == hostname {
			islocal = true
			beego.Info("Native " + servicename + " service is healthy!")
			break
		}

	}
	if !islocal {
		beego.Info("Native " + servicename + " service unhealthy or the service does not exist!")
		return
	} else {
		//slave()
		updatevalue := consulapi.KVPair{
			Key:   "service/" + servicename + "/leader",
			Value: []byte(""),
		}
		_, err = kv.Put(&updatevalue, nil)
		if err != nil {
			beego.Error("Update service/"+servicename+"/leader value failure!", err)
			return
		}
		beego.Info("Upata service/" + servicename + "/leader value success!")
		ip := beego.AppConfig.String("ip")
		port := beego.AppConfig.String("port")
		username := beego.AppConfig.String("username")
		password := beego.AppConfig.String("password")
		slave(ip, port, username, password)
	}
}

func SetConn(ip, port, username, password string) {
	config := &consulapi.Config{
		Datacenter: beego.AppConfig.String("datacenter"),
		Token:      beego.AppConfig.String("token"),
		Address:    beego.AppConfig.String("service_ip") + ":" + beego.AppConfig.String("service_port"),
	}
	hostname := beego.AppConfig.String("hostname")

	servicename := beego.AppConfig.String("servicename")
	//NewClient returns a new client
	client, err := consulapi.NewClient(config)
	if err != nil {
		beego.Error("Create  consul-api client failure! ", err)
		return
	}
	beego.Info("Create  consul-api client success!")
	kv := client.KV()
	session := client.Session()
	//	check := "serfHealth" + "," + "service:" + servicename +"""
	sessionEntry := consulapi.SessionEntry{
		LockDelay: 15 * time.Second,
		Name:      servicename,
		Node:      hostname,
		Checks:    []string{"serfHealth", "service:" + servicename},
	}
	//Create makes a new session. Providing a session entry can customize the session. It can also be nil to use defaults.
	sessionvalue, _, err := session.Create(&sessionEntry, nil)
	if err != nil {
		beego.Error("Session create failure!", err)
		return
	}
	beego.Info("Session create success!")
	acquirejson := `{"Node":"` + hostname + `","Ip":"` + ip + `","Port":` + port + `,"Username":"` + username + `","Password":"` + password + `"}`

	value := []byte(acquirejson)
	kvpair := consulapi.KVPair{
		Key:     "service/" + servicename + "/leader",
		Value:   value,
		Session: sessionvalue,
	}
	//Acquire is used for a lock acquisiiton operation. The Key, Flags, Value and Session are respected. Returns true on success or false on failures.
	time.Sleep(15 * time.Second)
	ok, _, err := kv.Acquire(&kvpair, nil)
	if err != nil {
		beego.Error("Set the connection string master failure! ", err)
		return
	}
	beego.Info("Set the connection string master success!")
	if !ok {
		time.Sleep(5 * time.Second)
		beego.Warn("kv acquire failure!")
		return
	} else {
		beego.Info("kv acquire success!")
	}
}
