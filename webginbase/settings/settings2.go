package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量用来保存程序的所有配置
var Conf = new(AppConfig)

// viper 库的结构体tag 是 mapstructure
type (
	AppConfig struct {
		Name         string `mapstructure:"name"`
		Mode         string `mapstructure:"mod"`
		Version      string `mapstructure:"version"`
		Port         int    `mapstructure:"port"`
		*LogConfig   `mapstructure:"log"`
		*MySQLConfig `mapstructure:"mysql"`
		*RedisConfig `mapstructure:"redis"`
	}

	LogConfig struct {
		Level      string `mapstructure:"level"`
		FileName   string `mapstructure:"filename"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxAge     int    `mapstructure:"max_age"`
		MaxBackups int    `mapstructure:"max_backups"`
	}

	MySQLConfig struct {
		Host         string `mapstructure:"host"`
		Port         int    `mapstructure:"port"`
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		DBName       int    `mapstructure:"db_name"`
		MaxOpenConns int    `mapstructure:"max_open_conns"`
		MaxIdleConns int    `mapstructure:"max_idle_conns"`
	}

	RedisConfig struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		DBName   int    `mapstructure:"db_name"`
		PoolSize int    `mapstructure:"pool_size"`
	}
)

func Init2() (err error) {
	viper.SetConfigName("config")         // 指定配置文件
	viper.SetConfigType("yaml")           // 指定配置文件
	viper.AddConfigPath("../")            // 指定查找配置文件的路径
	if viper.ReadInConfig(); err != nil { // 读取配置信息
		return
	}
	// 配置信息反序列化到 Conf 对象中
	if err = viper.Unmarshal(Conf); err != nil {
		return
	}

	// 监控配置文件变化
	viper.WatchConfig()
	// 配置文件发生变化后要同步到全局变量 Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已经被修改.")
		if err = viper.Unmarshal(Conf); err != nil {
			return
		}
	})

	return nil
}
