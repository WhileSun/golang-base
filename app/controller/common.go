package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/pkg/e"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type CommonController struct {

}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(62)]
	}
	return string(b)
}

func NewCommon() *CommonController{
	return &CommonController{}
}

func (c *CommonController) UploadPics(req *gin.Context,pathName string){
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// 单文件
	form,err := req.MultipartForm()
	if err != nil {
		fmt.Println(err.Error())
		e.New(req).MsgDetail(e.FAILED,"上传图片异常！")
		return
	}
	now := time.Now()
	fileDir := fmt.Sprintf("uploads/pics/%s/%d-%d/%d",pathName, now.Year(), now.Month(), now.Day())
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED,"上传图片异常！")
		return
	}
	files := form.File["image"]
	urls := make([]string,0)
	for _,file := range files{
		//判断后缀为图片的文件，如果是图片我们才存入到数据库中
		fileType := "other"
		fileExt := filepath.Ext(file.Filename)
		if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
			fileType = "img"
		}
		if fileType == "other"{
			continue
		}
		timeStamp := time.Now().Unix()
		fileName := fmt.Sprintf("%d.%s-%s", timeStamp,randSeq(3),file.Filename)
		filePathStr := filepath.Join(fileDir, fileName)
		req.SaveUploadedFile(file,filePathStr)
		urls=append(urls,filePathStr)
	}
	e.New(req).Data(e.SUCCESS,map[string]interface{}{"url":urls})
}
