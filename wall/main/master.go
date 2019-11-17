package main

import (
	"aaa/wall"
	"aaa/wall/configs"
	"aaa/wall/dao/DB"
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
	flag.StringVar(&confFile, "config", "E:/GOWorks/src/aaa/wall/main/master.json", "指定master.json")
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

	// 服务注册
	if err = wall.InitRegister(); err != nil {
		goto ERR
	}

	//任务管理器
	if err = wall.InitJobMgr(); err != nil {
		goto ERR
	}

	//启动API HTTP服务
	if err = wall.InitApiServer(); err != nil {
		goto ERR
	}

	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}

	return

ERR:
	fmt.Println(err)
}
