package utils

import (
	"clouds.lgb24kcs.cn/global"
	"github.com/spf13/viper"
	"log"
	"os"
)

func SystemInit() *viper.Viper {
	v := viper.New()
	dir, _ := os.Getwd()

	v.SetConfigFile(dir + "/config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	err = v.Unmarshal(&global.APP.Configuration)
	if err != nil {
		log.Println(err)
	}

	return v
}
