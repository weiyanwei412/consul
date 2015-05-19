package main

import (
	//	"fmt"
	//	"fmt"
	"log"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/mitchellh/cli"
)

func init() {
	beego.SetLogger("file", `{"filename":"logs/cmha-check.log"}`)
	beego.SetLogFuncCall(true)
}

func main() {
	defer beego.BeeLogger.Close()
	defer time.Sleep(100 * time.Millisecond)
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-v" || arg == "--version" {
			beego.Info("version 0.1.0")
			return
		}
	}
	c := &cli.CLI{
		Args: args,
		Commands: map[string]cli.CommandFactory{
			"check": checkCommandFactory,
		},
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)

}

func checkCommandFactory() (cli.Command, error) {
	//fmt.Println("Check")
	return &CheckCommandFactory{}, nil
}
