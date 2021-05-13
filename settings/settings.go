package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}



func Init(filePath string) (err error) {
	//viper.SetConfigName("config.json")
	viper.SetConfigFile(filePath)
	//viper.SetConfigFile("./conf/config.yaml") // 配置文件名称(无扩展名) 下面这个配置无效
	//viper.SetConfigType("yaml")   // (专用与从配置中心获取配置指定配置文件的格式)如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath(".")   // 还可以在工作目录中查找配置
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
		return
	}

	// 把读取到的配置反序列化到Conf结构体中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper unmarshal config file failed: %s \n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了!")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper unmarshal config file failed: %s \n", err)
		}
	})

	return
}
