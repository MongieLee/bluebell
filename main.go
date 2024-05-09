package main

import (
	"bluebell/controller"
	"bluebell/dao/mdb"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/setting"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("need config file.eg: bluebell config.yaml")
		return
	}
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config file failed, err:%v\n", err)
		return
	}
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineId); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	if err := mdb.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mdb failed, err:%v\n", err)
		return
	}
	defer mdb.Close()
	controller.InitTrans("zh")
	r := router.SetRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
