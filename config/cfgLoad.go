package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
)

func init() {
	var cfgPath string
	var env string
	flag.StringVar(&env, "env", "dev", "--init--")
	flag.Parse()
	if env == "prod" {
		cfgPath = "./config-prod.yaml"
	} else {
		cfgPath = "./config.yaml"
	}
	vip := viper.New()
	vip.SetConfigFile(cfgPath)
	_, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		panic(fmt.Errorf("配置文件不存在！: %s \n", err))
	}
	err = vip.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s \n", err))
	}
	Config = vip
	fmt.Printf("---配置文件init---\n")

}
