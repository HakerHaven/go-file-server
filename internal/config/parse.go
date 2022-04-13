package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// parse configure file
func InitParse(VERSION, GoVersion, BuildTime string) {
	var version bool
	var toml bool
	var configFileName string
	flag.BoolVar(&version, "v", false, "版本号")
	flag.BoolVar(&toml, "p", false, "参数信息")
	flag.StringVar(&configFileName, "c", "config.toml", "配置文件")
	flag.Parse()
	if version {
		res := strings.Split(VERSION, "-")
		if len(res) == 3 {
			if len(res[0]) == 0 {
				res[0] = "unknow"
			}
			if len(res[2]) == 0 {
				res[2] = "debug"
			}
			VERSION = fmt.Sprintf("%s-%s-%s", res[0], res[1], res[2])
		}
		fmt.Printf("Version: %s\n", VERSION)
		fmt.Printf("GoVersion: %s\n", GoVersion)
		fmt.Printf("BuildTime: %s\n", BuildTime)
		os.Exit(0)
	}
	// Parse configuration file
	InitConfig(configFileName)
	Conf.APP.Dir = getFileDir(configFileName)
	if toml {
		fmt.Println("Config:", Conf.String())
		os.Exit(0)
	}
}

func getFileDir(pathFile string) string {
	pathFile = pathFile[len(filepath.VolumeName(pathFile)):]
	// Find the last element
	i := len(pathFile) - 1
	for i >= 0 && !os.IsPathSeparator(pathFile[i]) {
		i--
	}
	if i <= 0 {
		i = 0
	}
	if i >= 0 && len(pathFile) > i {
		pathFile = pathFile[:i]
	}
	if pathFile == "" {
		return string(".")
	}
	realPath, _ := filepath.Abs(pathFile)
	return realPath
}
