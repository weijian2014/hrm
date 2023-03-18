package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"hrm/common"
	"hrm/db"
	"hrm/handler"
	"hrm/log"

	"github.com/julienschmidt/httprouter"
)

func init() {
	flag.BoolVar(&common.FlagInfos.IsHelp, "h", false, "Show help and exit")
	flag.StringVar(&common.FlagInfos.ConfigFileFullPath, "f", common.CurrDir+"/config.json", "The path of config.json file, support for absolute and relative paths")
	flag.BoolVar(&common.FlagInfos.IsInitSystem, "init", false, "Init system and exit")
	flag.StringVar(&common.FlagInfos.AdminUsername, "u", "admin", "User name for admin account")
	flag.StringVar(&common.FlagInfos.AdminPassword, "p", "123456", "Password for admin account")
	flag.Parse()
}

func main() {
	if common.FlagInfos.IsHelp {
		flag.Usage()
		return
	}

	if common.FlagInfos.IsInitSystem {
		err := db.Init(common.FlagInfos.AdminUsername, common.FlagInfos.AdminPassword)
		if err != nil {
			panic(fmt.Sprintf("Init database failed, %v", err))
		}
		fmt.Printf("Init system ok\n")
		return
	}

	_, err := os.Stat(common.FlagInfos.ConfigFileFullPath)
	if os.IsNotExist(err) {
		common.FlagInfos.ConfigFileFullPath = common.CurrDir + "/config/config.json"
		_, err := os.Stat(common.FlagInfos.ConfigFileFullPath)
		if os.IsNotExist(err) {
			common.FlagInfos.ConfigFileFullPath = common.CurrDir + "/config.json"
		}

		_, err = os.Stat(common.FlagInfos.ConfigFileFullPath)
		if os.IsNotExist(err) {
			panic("Please using -f option specifying a configuration file")
		}
	}

	common.JsonConfigs, err = common.LoadConfigFile(common.FlagInfos.ConfigFileFullPath)
	if nil != err {
		panic(err)
	}

	// LogFullPathName为空则输出到标准输出
	log.LoggerInit(common.JsonConfigs.LogLevel, common.JsonConfigs.LogRoll, common.JsonConfigs.LogFullPathName)
	log.System("Log init ok")

	handler.Init()
	router := httprouter.New()

	// // 静态资源
	// staticDir := http.Dir(common.CurrDir + "/" + common.JsonConfigs.StaticDirectory)
	// // html中引用静态资源时使用/static/开始为根目录, 例如/static/xxx.js或者/static/xxx.css
	// router.ServeFiles("/static/*filepath", staticDir)

	// // 页面
	// router.GET("/", handler.LoginPageHandler)
	// router.GET("/index", handler.IndexPageHandler)
	// router.GET("/admin", handler.AdminPageHandler)
	// router.GET("/employee", handler.EmployeePageHandler)

	// API
	router.POST("/api/v1/account/:action", handler.AccountHandler)
	router.POST("/api/v1/index/:action", handler.IndexHandler)
	router.POST("/api/v1/admin/:action", handler.AdminHandler)

	log.System("HRM system listen on %v", common.JsonConfigs.ServerListenHost)
	log.System("Json config from %v:\n%+v\n\n", common.FlagInfos.ConfigFileFullPath, common.JsonConfigs)
	log.Error("%v", http.ListenAndServe(common.JsonConfigs.ServerListenHost, router))
}
