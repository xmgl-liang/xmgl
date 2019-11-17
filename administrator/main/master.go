package main

import (
	"aaa/administrator"
	"aaa/administrator/configs"
	"aaa/administrator/dao/DB"
	"flag"
	"fmt"
	"runtime"
	"time"
)

var (
	confFile string //配置文件路径
)

//解析命令行参数
func initArgs() {
	// master -config ./master.json -xxx 123 -yyy ddd
	// master -h
	flag.StringVar(&confFile, "config", "E:/GOWorks/src/aaa/administrator/main/master.json", "指定master.json")
	flag.Parse()
}

//初始化线程数量
func initEnv() {
	//线程数量和CPU核数相同
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	//初始化命令行参数
	initArgs()

	//初始化线程
	initEnv()

	//加载配置
	if err = configs.InitConfig(confFile); err != nil {
		goto ERR
	}

	//连接数据库
	if err = DB.InitDB(); err != nil {
		goto ERR
	}

	//初始化服务发现模块
	if err = administrator.InitWorkerMgr(); err != nil {
		goto ERR
	}

	//日志管理器
	if err = administrator.InitLogMgr(); err != nil {
		goto ERR
	}

	//任务管理器
	if err = administrator.InitJobMgr(); err != nil {
		goto ERR
	}

	//启动API HTTP服务
	if err = administrator.InitApiServer(); err != nil {
		goto ERR
	}

	//启动12点自动更新用户的发表和接受次数
	//dao.UpdateUserMsg()

	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}

	return

ERR:
	fmt.Println(err)
}
