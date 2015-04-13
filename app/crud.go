package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"database/sql"
	consulapi "github.com/armon/consul-api"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type Kv struct {
	Node     string
	Ip       string
	Port     int
	Username string
	Password string
}

func GetLeader() error {
	config := &consulapi.Config{
		Datacenter: beego.AppConfig.String("datacenter"),
		Token:      beego.AppConfig.String("token"),
		Address:    beego.AppConfig.String("server_ip") + ":" + beego.AppConfig.String("server_port"),
	}
	servicename := beego.AppConfig.String("servicename")
	client, err := consulapi.NewClient(config)
	if err != nil {
		beego.Error("Create consul-api client failure!", err)
		return err
	}
	beego.Info(" Create consul-api client success!")
	kv := client.KV()
	//Get is used to lookup a single key
	for {
		kvPair, _, err := kv.Get("service/"+servicename+"/leader", nil)
		if err != nil {
			beego.Error("Get service/"+servicename+"/leader key failure!", err)
			return err
		}
		if kvPair == nil {
			beego.Error("service/" + servicename + "/leader not found, please create!")
			return err
		}
		beego.Info("Get service/" + servicename + "/leader key success!")
		//Are there external connection string provided
		if kvPair.Session != "" {
			beego.Info("service/" + servicename + "/leader Connection string  set up,Leader presence!")
			if kvPair.Value != nil {
				beego.Info("service/" + servicename + "/leader Connection string is not empty,Leader presence!")
				break
			} else {
				beego.Info("service/" + servicename + "/leader Connection string null,leader presence!")
				return nil
			}
		} else {
			beego.Info("Leader does not exist,Wait for the election leader!")
			continue
		}
	}
	err = Conn()
	if err != nil {
		beego.Error("Connect to the database failure!", err)
		return nil
	}
	beego.Info("Connect to the database success!")
	return nil

}

func Conn() error {
	server_ip := beego.AppConfig.String("server_ip")
	server_port := beego.AppConfig.String("server_port")
	table_name := beego.AppConfig.String("table_name")
	servicename := beego.AppConfig.String("servicename")
	database := beego.AppConfig.String("database")
	cli := new(http.Client)
	url := &url.URL{
		Scheme:   "http",
		Host:     server_ip + ":" + server_port,
		Path:     "/v1/kv/service/" + servicename + "/leader",
		RawQuery: "raw",
	}
	reg, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		beego.Error("Create Get request failure!:", err)
		return err
	}
	beego.Info("Create Get request success!")
	reg.Header.Set("Content-Type", "application/json")
	//执行get请求
	resp, err := cli.Do(reg)
	if err != nil {
		beego.Error("Send Get request failure!", err)
		return err
	}
	beego.Info("Send Get request success!")
	//分隔符
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error("Read leader data failure!", err)
		return err
	}
	beego.Info("Read leader data success!")
	kv := Kv{}
	//解析body到kv
	err = json.Unmarshal([]byte(body), &kv)
	if err != nil {
		beego.Error("Resolve leader data failure!", err)
		return err
	}
	beego.Info("Resolve leader data success!")
	//连接mysql数据库
	db, err := sql.Open("mysql", kv.Username+":"+kv.Password+"@tcp("+kv.Ip+":"+strconv.Itoa(kv.Port)+")/"+database+"?charset=utf8")
	if err != nil {
		beego.Error("open database connection string failure!", err)
		panic(err.Error())
	}
	beego.Info("open database connection string success!")
	err = db.Ping()
	if err != nil {
		beego.Error("ping() database failure!", err)
		return err
	}
	beego.Info("ping() database success!")
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	//sql := "select * from" + table_name
	raw, err := db.Query("select * from " + table_name)
	if err != nil {
		beego.Error("Inquiry data failure!", err)
		return err
	}
	beego.Info("Inquiry data success!")
	defer raw.Close()
	for raw.Next() {
		var id int
		var names string
		var created_at string
		err = raw.Scan(&id, &names, &created_at)
		if err != nil {
			beego.Error("scan() traversal data failure!", err)
		}
		//	beego.Info("scan() traversal data success!")
		fmt.Println("out", id, names, created_at)
	}
	//当前时间
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	sqlstr := "insert into " + table_name + "(`name`,`created_at`) values(?,?)"
	if _, err := tx.Exec(sqlstr, "test", nowtime); err != nil {
		beego.Error("Insert test value failure!", err)
		tx.Rollback()
		return err
	}
	beego.Info("Insert test value success!")
	row, err := db.Query("select name from " + table_name + " order by id desc limit 1")
	if err != nil {
		beego.Error("Inquiry the last line name failure!:", err)
		tx.Rollback()
		return err
	}
	beego.Info("Inquiry the last line name success!")
	defer row.Close()
	var name string
	for row.Next() {
		err = row.Scan(&name)
		if err != nil {
			beego.Error("scan() traversal name data failure!", err)
		}
		//beego.Info("scan() traversal name data success!")
		fmt.Println("output", name)
	}
	sqlstr = "update " + table_name + " set name='test1' where name=? order by id desc limit 1"
	if _, err := tx.Exec(sqlstr, name); err != nil {
		beego.Error("Update the last line data failure!", err)
		tx.Rollback()
		return err
	}
	beego.Info("Update the last line data success!")
	if err := tx.Commit(); err != nil {
		beego.Error("Submit data failure,End of things!", err)
		tx.Rollback()
		return err
	}
	beego.Info("Submit data success.End of things!")
	return nil
}
