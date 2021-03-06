package main

import (
	"os"
)

import (
	"global"
	"proxys/dbProxy"
	"proxys/transferProxy"
	. "tools"
	"tools/cfg"
)

//各个模块
import (
	_ "module/cache"
	_ "module/config"
	_ "module/user"
	"proxys/logProxy"
)

var (
	login_ip   string
	login_port string
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			ERR("caught panic in main()", x)
		}
	}()

	// 获取监听端口
	getPort()

	//启动
	global.Startup(global.ServerName, "login_log", nil)

	//连接TransferServer
	err := transferProxy.InitClient(cfg.GetValue("transfer_ip"), cfg.GetValue("transfer_port"))
	checkError(err)

	//连接DB
	dbProxyErr := dbProxy.InitClient(cfg.GetValue("db_ip"), cfg.GetValue("db_port"))
	checkError(dbProxyErr)

	//连接LogServer
	logProxyErr := logProxy.InitClient(cfg.GetValue("log_ip"), cfg.GetValue("log_port"))
	checkError(logProxyErr)

	//保持进程
	global.Run()
}

func getPort() {
	login_ip = cfg.GetValue("login_ip")
	login_port = cfg.GetValue("login_port")
	global.ServerName = "LoginServer[" + login_port + "]"
}

func checkError(err error) {
	if err != nil {
		ERR("Fatal error: %v", err)
		os.Exit(-1)
	}
}
