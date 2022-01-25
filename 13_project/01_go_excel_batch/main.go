package main

import (
	"GO_Learning/13_project/01_go_excel_batch/consumer"
	"GO_Learning/13_project/01_go_excel_batch/utils/modelUtils"
	"github.com/spf13/viper"
	"log"
)

func InitConf() *viper.Viper {
	config := viper.New()
	config.AddConfigPath("./conf")
	config.SetConfigName("conf")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		log.Println(err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(" Config file not found")
		} else {
			panic(" Config file was found but another error was produced")
		}
	}

	return config
}

func init() {
	conf := InitConf()
	err := modelUtils.InitDB(conf)
	if err != nil {
		log.Println("initDB err: ", err)
		panic(err)
	}
}

func main() {
	conf := InitConf()
	consumer.Consumer(conf)
}
