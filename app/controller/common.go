package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Common struct {
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(62)]
	}
	return string(b)
}

func NewCommon() *Common {
	return &Common{}
}

func (c *Common) UploadPics(req *gin.Context, pathName string, filename string) ([]string, error) {
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// 单文件
	urls := make([]string, 0)
	form, err := req.MultipartForm()
	if err != nil {
		return urls, err
	}
	now := time.Now()
	fileDir := fmt.Sprintf("uploads/pics/%s/%d-%d/%d", pathName, now.Year(), now.Month(), now.Day())
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		return urls, err
	}

	files := form.File[filename]
	for _, file := range files {
		//判断后缀为图片的文件，如果是图片我们才存入到数据库中
		fileType := "other"
		fileExt := filepath.Ext(file.Filename)
		if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
			fileType = "img"
		}
		if fileType == "other" {
			continue
		}
		timeStamp := time.Now().Unix()
		fileName := fmt.Sprintf("%d-%s%s", timeStamp, randSeq(5), fileExt)
		filePathStr := filepath.Join(fileDir, fileName)
		req.SaveUploadedFile(file, filePathStr)
		urls = append(urls, filePathStr)
	}
	return urls, err
}

func (c *Common) UploadFile(req *gin.Context, pathName string, filename string) (string, error) {
	url := ""
	file, err := req.FormFile(filename)
	if err != nil {
		return url, err
	}
	now := time.Now()
	fileDir := fmt.Sprintf("uploads/files/%s/%d-%d/%d", pathName, now.Year(), now.Month(), now.Day())
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		return url, err
	}
	fileExt := filepath.Ext(file.Filename)
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s%s", timeStamp, randSeq(5), fileExt)
	filePathStr := filepath.Join(fileDir, fileName)
	req.SaveUploadedFile(file, filePathStr)
	url = filePathStr
	return url, err
}

func (c *Common) UploadFiles(req *gin.Context, pathName string, filename string) ([]string, error) {
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// 单文件
	urls := make([]string, 0)
	form, err := req.MultipartForm()
	if err != nil {
		return urls, err
	}
	now := time.Now()
	fileDir := fmt.Sprintf("uploads/files/%s/%d-%d/%d", pathName, now.Year(), now.Month(), now.Day())
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		return urls, err
	}

	files := form.File[filename]
	for _, file := range files {
		//判断后缀为图片的文件，如果是图片我们才存入到数据库中
		//fileType := "other"
		fileExt := filepath.Ext(file.Filename)
		//if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		//	fileType = "img"
		//}
		//if fileType == "other" {
		//	continue
		//}
		timeStamp := time.Now().Unix()
		fileName := fmt.Sprintf("%d-%s%s", timeStamp, randSeq(5), fileExt)
		filePathStr := filepath.Join(fileDir, fileName)
		req.SaveUploadedFile(file, filePathStr)
		urls = append(urls, filePathStr)
	}
	return urls, err
}
