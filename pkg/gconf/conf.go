package gconf

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"runtime"
)
var Config *viper.Viper

func GetConf(filename string, filetype string) *viper.Viper{
	_, file, _, _ := runtime.Caller(1)
	config := viper.New()
	config.AddConfigPath(filepath.Dir(file)+"/../../config")  //设置读取的文件路径
	config.SetConfigName(filename) //设置读取的文件名
	config.SetConfigType(filetype) //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		log.Fatal("pkg/config read failed: "+err.Error())
	}
	return config
}

func init(){
	//基础配置文件读取
	Config = GetConf("config","yaml")
}