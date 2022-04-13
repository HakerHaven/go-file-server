package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	APP  AppConfig  `mapstructure:"app"`
	File FileConfig `mapstructure:"file"`
}

type AppConfig struct {
	Dir string `mapstructure:"dir"`
}

type FileConfig struct {
	UploadPath string `mapstructure:"uploadPath"`
}

func (conf *Config) String() string {
	b, err := json.Marshal(*conf)
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	return out.String()
}

func InitConfig(filename string) {
	viper.SetConfigType("toml")
	viper.SetConfigFile(filename)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("配置文件未找到, %s\n", filename)
		os.Exit(-1)
	}
	viper.Unmarshal(&Conf)
}
