package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	CurrDir     string
	JsonConfigs *JsonConfigStruct
	FlagInfos   FlagInfoStruct
)

func init() {
	currFullPath, err := exec.LookPath(os.Args[0])
	if nil != err {
		panic(err)
	}

	absFullPath, err := filepath.Abs(currFullPath)
	if nil != err {
		panic(err)
	}
	CurrDir = filepath.Dir(absFullPath)
}

type JsonConfigStruct struct {
	Version                string `json:"Version"`
	LogLevel               int    `json:"LogLevel"`
	LogRoll                int    `json:"LogRoll"`
	ServerListenHost       string `json:"ServerListenHost"`
	CookieName             string `json:"CookieName"`
	PostgresqlHost         string `json:"PostgresqlHost"`
	PostgresqlPort         uint16 `json:"PostgresqlPort"`
	PostgresqlUser         string `json:"PostgresqlUser"`
	PostgresqlPassword     string `json:"PostgresqlPassword"`
	PostgresqlDatabaseName string `json:"PostgresqlDatabaseName"`
}

type FlagInfoStruct struct {
	IsHelp             bool
	ConfigFileFullPath string
}

// 读取json配置文件
func LoadConfigFile(filePath string) (*JsonConfigStruct, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	c := &JsonConfigStruct{}
	if err := json.Unmarshal(cData, c); nil != err {
		return nil, err
	}

	return c, nil
}
