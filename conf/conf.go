package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//定义server的yaml结构体
type serverYaml struct {
	Ip           string `yaml:"Ip"`
	Port         string `yaml:"Port"`
	ReadTimeout  int    `yaml:"ReadTimeout"`
	WriteTimeout int    `yaml:"WriteTimeout"`
}

//定义VxPush的yaml结构体
type vxPushYaml struct {
	Appid      string `yaml:"Appid"`
	Secret     string `yaml:"Secret"`
	Templateid string `yaml:"Templateid"`
	Url        string `yaml:"Url"`
}

//定义Mysql的yaml结构体
type mysqlYaml struct {
	Addr     string `yaml:"Addr"`
	User     string `yaml:"User"`
	Passowrd string `yaml:"Passowrd"`
	DbName   string `yaml:"DbName"`
}

type ConfYaml struct {
	ServerYaml serverYaml `yaml:"Server"`
	VxPushYaml vxPushYaml `yaml:"VxPush"`
	MysqlYaml  mysqlYaml  `yaml:"Mysql"`
}

var Conf ConfYaml

func init() {
	config, err := ioutil.ReadFile("./conf/conf.yaml")
	if err != nil {
		fmt.Println("获取配置文件失败，请检查配置文件: ", err)
		panic(err)
	}
	// fmt.Println(config)
	err1 := yaml.Unmarshal(config, &Conf)
	if err1 != nil {
		fmt.Println("配置文件导入失败，请检查配置文件是否配置正常: ", err)
		panic(err)
	}
	// fmt.Println(Conf.ServerYaml)
	// fmt.Println(Conf.VxPushYaml)
}
