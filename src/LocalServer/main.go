package main

import (
	"os"
)

import (
	"github.com/funny/link"
	"global"
	"module"
	"proxys/redisProxy"
	"proxys/dbProxy"
	. "tools"
	"tools/cfg"
	"tools/db"
)

//各个模块
import (
	_ "module/cache"
	_ "module/config"
	_ "module/user"
	"proxys/gameProxy"
)

var (
	local_ip   string
	local_port string
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			ERR("caught panic in main()", x)
		}
	}()

	// 获取端口号
	getPort()

	//启动
	global.Startup(global.ServerName, "local_log", stopLocalServer)

	//开启LocalServer监听
	startLocalServer()

	//保持进程
	global.Run()
}

func getPort() {
	//端口号
	local_ip = cfg.GetValue("local_ip")
	local_port = cfg.GetValue("local_port")
	global.ServerName = "LocalServer[" + local_port + "]"
}

func startLocalServer() {
	//连接Redis
	redisProxyErr := redisProxy.InitClient(cfg.GetValue("redis_ip"), cfg.GetValue("redis_port"), cfg.GetValue("redis_pwd"))
	checkError(redisProxyErr)

	//开启DB
	db.Init()

	//开启同步DB数据到数据库
	dbProxy.StartSysDB()

	//开启客户端监听
	addr := "0.0.0.0:" + local_port
	err := global.Listener("tcp", addr, global.PackCodecType_UnSafe,
		func(session *link.Session) {
			global.AddSession(session)
		},
		gameProxy.MsgDispatch,
	)
	checkError(err)
}

func stopLocalServer() {
	INFO("Waiting SyncDB...")
	dbProxy.SyncDB()
	INFO("SyncDB Success")
}

func checkError(err error) {
	if err != nil {
		ERR("Fatal error: %v", err)
		os.Exit(-1)
	}
}
