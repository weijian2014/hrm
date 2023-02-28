package main

import (
	"flag"
	"net/http"
	"os"

	"hrm/common"
	"hrm/handler"
	"hrm/log"

	"github.com/julienschmidt/httprouter"
)

func init() {
	// common option
	flag.BoolVar(&common.FlagInfos.IsHelp, "h", false, "Show help")
	flag.StringVar(&common.FlagInfos.ConfigFileFullPath, "f", common.CurrDir+"/config.json", "The path of config.json file, support for absolute and relative paths")
}

func main() {
	if common.FlagInfos.IsHelp {
		flag.Usage()
		return
	}

	var logLevel int = 0
	var logRoll int = 50000

	_, err := os.Stat(common.FlagInfos.ConfigFileFullPath)
	if os.IsNotExist(err) {
		common.FlagInfos.ConfigFileFullPath = common.CurrDir + "/config/config.json"
		_, err := os.Stat(common.FlagInfos.ConfigFileFullPath)
		if os.IsNotExist(err) {
			common.FlagInfos.ConfigFileFullPath = common.CurrDir + "/../config/config.json"
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

	logLevel = common.JsonConfigs.LogLevel
	logRoll = common.JsonConfigs.LogRoll

	// The third parameter "" mean the log output to stdout
	log.LoggerInit(logLevel, logRoll, "")

	log.System("Log init ok.")
	handler.Init()
	router := httprouter.New()

	// 静态资源
	staticDir := http.Dir(common.CurrDir + "/web/static")
	router.ServeFiles("/static/*filepath", staticDir)

	// 页面
	router.GET("/", handler.LoginPageHandler)
	router.GET("/index", handler.IndexPageHandler)
	router.GET("/admin", handler.AdminPageHandler)

	// API
	router.POST("/api/v1/account/:action", handler.AccountHandler)
	router.POST("/api/v1/index/:action", handler.IndexHandler)
	router.POST("/api/v1/admin/:action", handler.AdminHandler)

	log.System("HRM system listen on %v", common.JsonConfigs.ServerListenHost)
	http.ListenAndServe(common.JsonConfigs.ServerListenHost, router)
}
