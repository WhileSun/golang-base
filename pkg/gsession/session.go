package gsession

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var store cookie.Store
var secret = "oms"
var sessionName = "SessionID"

func init() {
	store = cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{
		MaxAge: 60 * 30, //30min
		Path:   "/",
	})
	fmt.Println("session init")
}

func GinUse() gin.HandlerFunc {
	return sessions.Sessions(sessionName, store)
}

func Get(c *gin.Context, sessionName string) interface{} {
	return sessions.Default(c).Get(sessionName)
}

func Set(c *gin.Context, sessionName string, value interface{}) error {
	session := sessions.Default(c)
	session.Set(sessionName, value)
	//sessions.
	return session.Save()
}

func Delete(c *gin.Context, sessionName string) error {
	session := sessions.Default(c)
	session.Delete(sessionName)
	return session.Save()
}

func Clear(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	return session.Save()
}

func GetID(c *gin.Context) string {
	session := sessions.Default(c)
	return session.ID()
}
