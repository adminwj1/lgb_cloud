package global

import (
	"clouds.lgb24kcs.cn/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Config struct {
	Configuration config.Configuration
	Viper         *viper.Viper
	DB            *gorm.DB
	Log           *zap.Logger
}

var APP = new(Config)
