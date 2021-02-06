package admin

import (
	"blog/blog_api/utils"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//上传文件
func Upload(c *gin.Context)  {
	file,err := c.FormFile("file")
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "文件上传错误"}
		res.Json(c)
		return
	}

	now := time.Now()
	ext := path.Ext(file.Filename)
	fileName := strconv.Itoa(now.Nanosecond()) + ext
	filePath := fmt.Sprintf("%s%s%s%s",
		"./static/upload",
		fmt.Sprintf("%04d",now.Year()),
		fmt.Sprintf("%02d",now.Month()),
		fmt.Sprintf("%02d",now.Day()),
		)
	utils.CreateDir(filePath)
	fullPath := filePath + "/" +fileName
	c.SaveUploadedFile(file,fullPath)
	url := fullPath[1:len(fullPath)]
	res := &utils.Response{Code: 0,Msg: "OK",Data: url}
	res.Json(c)
}

//上传文件编辑
func UploadEdit(c *gin.Context) {
	file, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errno": 1000,
		})
		return
	}

	now := time.Now()
	ext := path.Ext(file.Filename)
	fileName := strconv.Itoa(now.Nanosecond()) + ext
	filePath := fmt.Sprintf("%s%s%s%s",
		"./static/upload/",
		fmt.Sprintf("%04d", now.Year()),
		fmt.Sprintf("%02d", now.Month()),
		fmt.Sprintf("%02d", now.Day()),
		)
	fmt.Println(filePath)
	utils.CreateDir(filePath)

	fullPath := filePath + "/" + fileName
	c.SaveUploadedFile(file, fullPath)
	url := fullPath[1:len(fullPath)]
	c.JSON(http.StatusOK, gin.H{
		"errno": 0,
		"data":  []string{url},
	})
	return
}

