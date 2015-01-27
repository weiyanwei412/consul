package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	//"reflect"
)

func slave() {
	db, err := sql.Open("mysql", "root:111111@tcp(192.168.2.62:3306)/")
	if err != nil {
		beego.Error("Connect to the mysql fails:", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		beego.Error("db.ping():", err)
	}
	_, err = db.Query("stop slave io_thread")
	if err != nil {
		beego.Error("stop slave io_thread faild:", err)
	}
	row, err := db.Query("show slave status")
	cols, _ := row.Columns()
	fmt.Println()
	buffer := make([]interface{}, len(cols))
	data := make([]interface{}, len(cols))
	for i, _ := range buffer {
		buffer[i] = &data[i]
	}
	for row.Next() {
		err = row.Scan(buffer...)
		if err != nil {
			beego.Error("scan err:", err)
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
	//	fmt.Printf("Master_Log_File=%s\n", Master_Log_File)
	//	fmt.Printf("Read_Master_Log_Pos=%s\n", Read_Master_Log_Pos)
	sqlstr := "select master_pos_wait(?,?)"
	rowss, err := db.Query(sqlstr, Master_Log_File, Read_Master_Log_Pos)
	var master_pos_wait string
	for rowss.Next() {
		err = rowss.Scan(&master_pos_wait)
		if err != nil {
			beego.Error("error", err)
		}
		fmt.Println("output", master_pos_wait)
		if master_pos_wait < "0" && master_pos_wait == "null" {
			beego.Error("切换数据库失败!", err)
			//return err
		}
		_, err := db.Query("set global read_only=0")
		if err != nil {
			beego.Error("设置可读可写失败!", err)
			//	return err
		}
	}

}
