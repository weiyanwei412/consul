package main

import (
	"database/sql"
//	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	//"reflect"
)

func slave() {
	db, err := sql.Open("mysql", "root:111111@tcp(localhost:3306)/")
	if err != nil {
		beego.Error("Connect to the mysql fails:", err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		beego.Error("连接mysql数据库失败", err)
		return
	}
	beego.Info("连接mysql数据库成功")
	_, err = db.Query("stop slave io_thread")
	if err != nil {
		beego.Error("关闭slave线程失败", err)
		return
	}
	beego.Info("关闭slave线程成功")
	row, err := db.Query("show slave status")
	if err != nil {
		beego.Error("查询slave状态失败",err)
		return
	}
	beego.Info("查询slave状态成功")
	cols, _ := row.Columns()
	buffer := make([]interface{}, len(cols))
	data := make([]interface{}, len(cols))
	for i, _ := range buffer {
		buffer[i] = &data[i]
	}
	for row.Next() {
		err = row.Scan(buffer...)
		if err != nil {
			beego.Error("scan err:", err)
			return
		}
	}
	mapField2Data := make(map[string]interface{}, len(cols))
	for k, col := range data {
		mapField2Data[cols[k]] = col
	}
	//	fmt.Printf("Master_Log_File=%s\n", mapField2Data["Master_Log_File"])
	//	fmt.Printf("Read_Master_Log_Pos=%s\n", mapField2Data["Read_Master_Log_Pos"])
	Master_Log_File := mapField2Data["Master_Log_File"]
	Read_Master_Log_Pos := mapField2Data["Read_Master_Log_Pos"]
	Slave_SQL_Running := mapField2Data["Slave_SQL_Running"]
//	fmt.Printf("Master_Log_File=%s\n", Master_Log_File)
//	fmt.Printf("Read_Master_Log_Pos=%s\n", Read_Master_Log_Pos)
//	fmt.Printf("Slave_SQL_Running=%s\n", Slave_SQL_Running)
	//fmt.Println(string(Slave_SQL_Running.([]uint8)))
	if string(Slave_SQL_Running.([]uint8)) != "Yes" {
		beego.Error("SQL的复制线程不正常!", err)
		return
	}
	beego.Info("SQL的复制线程正常")
	sqlstr := "select master_pos_wait(?,?)"
	rowss, err := db.Query(sqlstr, Master_Log_File, Read_Master_Log_Pos)
	if err != nil {
		beego.Error("执行master_pos_wait函数失败",err)
		return
	}
	beego.Info("执行master_pos_wait函数成功")
	var master_pos_wait string
	for rowss.Next() {
		err = rowss.Scan(&master_pos_wait)
		if err != nil {
			beego.Error("error", err)
			return
		}
//		fmt.Println("output", master_pos_wait)
		if master_pos_wait < "0" && master_pos_wait == "null" {
			beego.Error("切换数据库失败!", err)
			return
		}
		beego.Info("切换数据库成功")
		_, err := db.Query("set global read_only=0")
		if err != nil {
			beego.Error("设置可读可写失败!", err)
			//	return err
		}
		beego.Info("设置可读可写成功")
	}
	SetConn()
}
