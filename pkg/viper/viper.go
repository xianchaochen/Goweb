package viper

import (
	"bluebell/config"
	"bluebell/pkg/zaplogger"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)


func Init(filePath string) (err error) {
	//viper.SetConfigName("config.json")
	viper.SetConfigFile(filePath)
	//viper.SetConfigFile("./config/config.yaml") // 配置文件名称(无扩展名) 下面这个配置无效
	//viper.SetConfigType("yaml")   // (专用与从配置中心获取配置指定配置文件的格式)如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath(".")   // 还可以在工作目录中查找配置
	err = viper.ReadInConfig()
	if err != nil {
		return errors.New(fmt.Sprintf("Fatal error config file: %s \n", err))
	}

	// 把读取到的配置反序列化到Conf结构体中
	if err := viper.Unmarshal(config.GlobalConfig); err != nil {
		return errors.New(fmt.Sprintf("viper unmarshal config file failed: %s \n", err))
	}

	watch()

	return nil
}

func watch()  {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		if err := viper.Unmarshal(config.GlobalConfig); err != nil {
			zap.L().Error("config modified and unmarshal config file failed")
		}

		gin.SetMode(config.GlobalConfig.Mode)

		if err := zaplogger.Init(config.GlobalConfig.LogConfig); err != nil {
			zap.L().Error("config modified and reload zaplogger config failed")
		}
		return
	})

	return
}
