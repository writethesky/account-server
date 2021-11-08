package internal

import (
	"flag"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type configEntity struct {
	TokenServer serverEntity `yaml:"token_server"`
	UserServer  serverEntity `yaml:"user_server"`
	MySQL       mysqlEntity  `yaml:"mysql"`
	Server      serverEntity `yaml:"server"`
}

type mysqlEntity struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

type serverEntity struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

var Config *configEntity

func init() {
	var configFileName string
	flag.StringVar(&configFileName, "config", "", "配置文件位置")
	flag.Parse()

	if "" == configFileName {
		flag.Usage()
		panic(nil)
	}

	file, err := os.Open(configFileName)
	if nil != err {
		panic(err)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if nil != err {
		panic(err)
	}
	Config = new(configEntity)
	err = yaml.Unmarshal(fileBytes, Config)
	if nil != err {
		panic(err)
	}

}
