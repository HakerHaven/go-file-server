package file

import (
	"fmt"
	"go-file-server/internal/config"
	"go-file-server/internal/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	fileInfo := config.Conf.File
	// Single file
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusNotFound, "文件上传出错")
	}
	log.Println("文件名称：" + file.Filename)
	log.Println("文件大小：" + strconv.FormatInt(file.Size, 10))

	fmt.Print("------------" + fileInfo.UploadPath)
	uploadPath := util.IsMkdirFile(fileInfo.UploadPath)
	fmt.Print(uploadPath)
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, uploadPath+file.Filename)

	// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	c.JSON(http.StatusOK, gin.H{"fileDownUrl": "https://127.0.0.1:8080/info/lisang"})
}
