package setting

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// AppSection app.yaml server 配置
type AppSection struct {
	JwtSecret string `yaml:"jwt_secret"`
}

// ServerSection app.yaml server 配置
type ServerSection struct {
	HTTPPort int `yaml:"http_port"`
}

// 所有的配置项
var (
	App    AppSection
	Server ServerSection
)

// app.yaml 对应结构体
var config struct {
	App    AppSection
	Server ServerSection
}

// init 初始化加载配置文件
func init() {
	// 解析 app.yml
	file, err := ioutil.ReadFile("config/app.yml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal([]byte(file), &config)
	if err != nil {
		log.Fatalln(err)
	}
	App = config.App
	Server = config.Server
}
