package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"strconv" 

	"database/sql"
        _"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

type Kv struct {
	Node     string
	Ip       string
	Port     int
	Username string
	Password string
	Holdon   string
}



func GetLeader() error {
	client := new(http.Client)
	url := &url.URL{
		Scheme:   "http",
		Host:     "192.168.2.71:8500",
		Path:     "/v1/kv/service/mysql-1/leader",
		RawQuery: "raw",
	}
	reg, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		beego.Error("Create a get request fails:", err)
		return err
	}
	reg.Header.Set("Content-Type", "application/json")
	//执行get请求
	resp, err := client.Do(reg)
	if err != nil {
		beego.Error("Send http request fails:", err)
		return err
	}
	//分隔符
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error("Failed to read data:", err)
		return err
	}
	//      fmt.Println(string(body))
	//      var bodybyte = []byte(body)
	//      var kv = make([]Kv,0)
	kv := Kv{}
	//解析body到kv
	err = json.Unmarshal([]byte(body), &kv)
	if err != nil {
		beego.Error("Failed parsing json:", err)
		return err
	}
	//      fmt.Println(string(kv[0].value))
	//连接mysql数据库
 	db, err := sql.Open("mysql", kv.Username + ":" + kv.Password + "@tcp(" +kv.Ip + ":" + strconv.Itoa(kv.Port) + ")/upm_testdb?charset=utf8")
    	if err != nil {
		beego.Error("Connect to the database fails:", err)
        	panic(err.Error())
   	}
	err = db.Ping()
	if err != nil {
		beego.Error("db.ping()",err)
		return err
	}
    	defer db.Close()	
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	raw, err := db.Query("select * from test_tbl"); 
	if err != nil {
		beego.Error("select data failed:", err)
		return err
	}
        defer raw.Close()
        for raw.Next() {
		var id int
		var names string
		var created_at string
                err = raw.Scan(&id,&names,&created_at)
                if err != nil{
                        beego.Error("error", err)
                }
                fmt.Println("out",id,names,created_at)
        }
	//当前时间
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	sqlstr := "insert into test_tbl(`name`,`created_at`) values(?,?)"
	if _, err := tx.Exec(sqlstr,"lindan",nowtime); err != nil 	{
		beego.Error("Insert data failed:",err)
		tx.Rollback()
		return err
	}
	row, err := db.Query("select name from test_tbl order by id desc limit 1")
	if err != nil {
                beego.Error("select data failed:", err)
                tx.Rollback()
                return err
        }
	defer row.Close()
	var name string
	for row.Next() {
    		err = row.Scan(&name)
		if err != nil{
			beego.Error("error", err)
		}
		fmt.Println("output",name)
	}
	sqlstr = "update test_tbl set name='lindan1' where name=? order by id desc limit 1"
        if _, err := tx.Exec(sqlstr,name); err != nil         {
                beego.Error("Insert data failed:",err)
                tx.Rollback()
                return err
        }

	if err := tx.Commit(); err != nil {
		beego.Error("commit failed:",err)
		tx.Rollback()
		return err
	}
	return nil
	
}
