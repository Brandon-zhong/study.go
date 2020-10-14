package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const configFile = "/data/ftp/server/config.yaml"

var Config ServerConfig

func init() {
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("read config file error, file path: ", configFile)
	}
	if err := v.Unmarshal(&Config); err != nil {
		fmt.Println("unmarshal config error")
	}
}

type ServerConfig struct {
	Mysql Mysql `mapstruecture:"mysql"`
}

type Mysql struct {
	Url          string `mapstructure:"url"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"db-name"`
	MaxIdleConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
	LogMod       bool   `mapstructure:"log-mod"`
}
