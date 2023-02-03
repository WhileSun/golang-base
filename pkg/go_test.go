package pkg

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/whilesun/go-admin/pkg/core/gcasbin"
	"github.com/whilesun/go-admin/pkg/core/gdb"
	"github.com/whilesun/go-admin/pkg/core/glog"
	"github.com/whilesun/go-admin/pkg/core/gredis"
	"github.com/whilesun/go-admin/pkg/helper/gcaptcha"
	"github.com/whilesun/go-admin/pkg/helper/gjwt"
	"github.com/whilesun/go-admin/pkg/helper/gsession"
	"github.com/whilesun/go-admin/pkg/utils/gcrypto"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	log := glog.New("log")
	for i := 0; i <= 100; i++ {
		go func(i int) {
			log.Info(i)
		}(i)
	}
	time.Sleep(10 * time.Second)
}

func TestDb(t *testing.T) {
	db := gdb.New("")
	for j := 0; j <= 500; j++ {
		go func(j int) {
			result := make([]map[string]interface{}, 0)
			db.Raw("Select * from s_user").Scan(&result)
			fmt.Println(j)
		}(j)
	}
	time.Sleep(20 * time.Second)
}

func TestRedis(t *testing.T) {
	gr := gredis.New("redis")
	gr.Set("abc", 100, "1")
	gr.Set("ab1c", 100, "1")
}

func TestCasbin(t *testing.T) {
	gcasbin.New("")
}

func TestSession(t *testing.T) {
	gsession.New("")
}

func TestJwt(t *testing.T) {
	g := gjwt.New("")
	result, _ := g.CreateToken(jwt.MapClaims{"123": 123})
	fmt.Println(result)
	//time.Sleep(12 * time.Second)
	str, err := g.ParseToken(result)
	fmt.Println(err)
	fmt.Printf("%+v", str)
}

func TestCaptcha(t *testing.T) {
	c := gcaptcha.New("")
	id, bs64, _ := c.Generate()
	fmt.Println(id, bs64)
}

func TestToken(t *testing.T) {
	fmt.Println(gcrypto.Md5Encode("1"))
}
