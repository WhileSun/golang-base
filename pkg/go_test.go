package pkg

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/whilesun/go-admin/pkg/gcasbin"
	"github.com/whilesun/go-admin/pkg/gcrypto"
	"github.com/whilesun/go-admin/pkg/gdb"
	"github.com/whilesun/go-admin/pkg/gjwt"
	"github.com/whilesun/go-admin/pkg/glog"
	"github.com/whilesun/go-admin/pkg/gredis"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	log := glog.Get()
	log.Error("sdf")
}

func TestDb(t *testing.T){
	gdb.Run()
	db := gdb.Get()
	for j:=0;j<=100;j++ {
		go func(j int) {
			result := make([]map[string]interface{},0)
			db.Raw("Select * from s_user").Scan(&result)
			fmt.Println(j)
		}(j)
	}
	time.Sleep(20*time.Second)
}

func TestCasbin(t *testing.T){
	gcasbin.Get()
}

func TestJwt(t *testing.T){
	result,_:=gjwt.CreateToken(jwt.MapClaims{"123":123})
	fmt.Println(result)
	//time.Sleep(2*time.Second)
	str,err := gjwt.ParseToken(result)
	fmt.Println(err)
	fmt.Printf("%+v",str)
}


func TestRedis(t *testing.T){
	gredis.Run()
	redis := gredis.GetPool()
	for j:=0;j<=100;j++ {
		go func(j int) {
			conn := redis.Get()
			defer conn.Close()
			res, _ := conn.Do("set", fmt.Sprintf("abc%d",j), 100,"EX",10)
			fmt.Println(res, j)
		}(j)
	}
	time.Sleep(20*time.Second)
}

func TestToken(t *testing.T){
	fmt.Println(gcrypto.Md5Encode("1"))
}
