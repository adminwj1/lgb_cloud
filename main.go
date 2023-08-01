package main

import (
	//"clouds.lgb24kcs.cn/config"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/route"
	"clouds.lgb24kcs.cn/utils"
	"fmt"
)

func main() {
	global.APP.Viper = utils.SystemInit()
	global.APP.Log = utils.InitLogger()
	global.APP.DB = utils.InitDB()

	fmt.Println(global.APP.DB)
	route.Start()
}
