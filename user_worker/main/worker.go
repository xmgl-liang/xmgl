package main

import (
	"aaa/user_worker"
	"aaa/user_worker/configs"
	"aaa/user_worker/tasks/DB"
	"flag"
	"fmt"
	"runtime"
	"time"
)

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	// worker -config ./worker.json
	// worker -h
	flag.StringVar(&confFile, "config", "E:/GOWorks/src/aaa/user_worker/main/worker.json", "worker.json")
	flag.Parse()
}

// 初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	// 初始化命令行参数
	initArgs()

	// 初始化线程
	initEnv()

	// 加载配置
	if err = configs.InitConfig(confFile); err != nil {
		goto ERR
	}

	//连接数据库
	if err = DB.InitDB(); err != nil {
		goto ERR
	}

	// 服务注册
	if err = user_worker.InitRegister(); err != nil {
		goto ERR
	}

	// 启动日志协程
	if err = user_worker.InitLogSink(); err != nil {
		goto ERR
	}

	// 启动执行器
	if err = user_worker.InitExecutor(); err != nil {
		goto ERR
	}

	// 启动调度器
	if err = user_worker.InitScheduler1(); err != nil {
		goto ERR
	}

	if err = user_worker.InitScheduler2(); err != nil {
		goto ERR
	}

	// 初始化任务管理器
	if err = user_worker.InitJobMgr(); err != nil {
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
