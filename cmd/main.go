package main

import (
	"go-file-server/internal/config"
	"go-file-server/internal/file"

	"github.com/gin-gonic/gin"
)

var VERSION string = "unknow"

var GoVersion string = "unknow"

var BuildTime string = "unknow"

func init() {
	config.InitParse(VERSION, GoVersion, BuildTime)
}

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", file.Upload)
	router.Run()
}
