package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Database       string `yaml:"database" json:"database"`
	Table          string `yaml:"table" json:"table"`
	ServiceName    string `yaml:"serviceName" json:"serviceName"`
	GoPackageName  string `yaml:"goPackageName" json:"goPackageName"`
	Port           int    `yaml:"port" json:"port"`
	Password       string `yaml:"password" json:"password"`
	User           string `yaml:"user" json:"user"`
	PackageName    string `yaml:"packageName" json:"packageName"`
	IgnoreTableStr string `yaml:"ignoreTableStr" json:"ignoreTableStr"`
	DbType         string `yaml:"dbType" json:"dbType"`
	Host           string `yaml:"host" json:"host"`
	FilePath       string `yaml:"filePath" json:"filePath"`
}

var config Config

func InitConfig() error {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error("配置文件读取err", err)
		return err
	}
	var configData Config
	if err = viper.Unmarshal(&configData); err != nil {
		logrus.Error("配置文件 Unmarshal err", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Info("配置文件发生更新")
		if err := viper.Unmarshal(&configData); err != nil {
			logrus.Infof("配置文件更新解析失败,err:%v", err)
		}
	})
	SetConfig(&configData)
	return nil
}

func SetConfig(cfg *Config) {
	config = *cfg
}

func GetConfig() *Config {
	return &config
}
