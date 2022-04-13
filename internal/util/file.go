package util

import (
	"log"
	"os"
	"strconv"
	"time"
)

// 防止文件覆盖，根据时间创建存放目录
func IsMkdirFile(uploadPath string) string {
	timeObj := time.Now()
	var strParent = timeObj.Format("20060102")
	var strSon = strconv.FormatInt(timeObj.UnixNano(), 10)
	uploadPath = uploadPath + "/" + strParent
	_, err := os.Stat(uploadPath)
	if err != nil {
		err := os.Mkdir(uploadPath, 0666)
		if err != nil {
			log.Println("父目录创建失败")
		}
	}

	uploadPath = uploadPath + "/" + strSon
	_, err = os.Stat(uploadPath)
	if err != nil {
		err := os.Mkdir(uploadPath, 0666)
		if err != nil {
			log.Println("子目录创建失败")
		}
	}

	return uploadPath + "/"
}
