package e

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Empty = make([]string, 0)

type RespData struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time int64       `json:"time"`
}

type RespObj struct {
	ctx *gin.Context
}

func New(req *gin.Context) *RespObj {
	return &RespObj{
		ctx:req,
	}
}

func transMsg(code uint, detail string) (msg string){
	msg = GetMessage(code)
	if detail != ""{
		msg += ", "+detail
	}
	return msg
}


func(resp RespObj) Msg(code uint){
	resp.ctx.JSON(http.StatusOK, &RespData{
		Code: code,
		Msg:  GetMessage(code),
		Data: Empty,
		Time: time.Now().Unix(),
	})
}

// MsgDetail 加自定义报错信息
func(resp RespObj) MsgDetail(code uint , msg string){
	resp.ctx.JSON(http.StatusOK, &RespData{
		Code: code,
		Msg:  transMsg(code,msg),
		Data: Empty,
		Time: time.Now().Unix(),
	})
}

func(resp RespObj) Data(code uint,data interface{}){
	if data == nil{
		fmt.Println("empty")
	}
	resp.ctx.JSON(http.StatusOK, &RespData{
		Code: code,
		Msg:  GetMessage(code),
		Data: data,
		Time: time.Now().Unix(),
	})
}

func(resp RespObj) DataDetail(code uint, msg string, data interface{}){
	resp.ctx.JSON(http.StatusOK, &RespData{
		Code: code,
		Msg:  transMsg(code,msg),
		Data: data,
		Time: time.Now().Unix(),
	})
}