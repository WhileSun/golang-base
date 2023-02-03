package gsession

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/pkg/core/gconfig"
)

type SessionConfig struct {
	Secret      string
	SessionName string
	Store       cookie.Store
}

func New(sessionKey string) *SessionConfig {
	// 默认log配置
	if sessionKey == "" {
		sessionKey = "session"
	}
	sessionConfig := &SessionConfig{
		Secret:      "UlhknW5WBBL0lqOgLzEDp1FlqvqvFVBW",
		SessionName: "SessionID",
	}
	gconfig.Config.UnmarshalKey(sessionKey, sessionConfig)
	store := sessionConfig.initSession()
	sessionConfig.Store = store
	return sessionConfig
}

func (sessionConfig *SessionConfig) initSession() cookie.Store {
	store := cookie.NewStore([]byte(sessionConfig.Secret))
	store.Options(sessions.Options{
		MaxAge: 60 * 30, //30min
		Path:   "/",
	})
	return store
}

func (sessionConfig *SessionConfig) Use() gin.HandlerFunc {
	return sessions.Sessions(sessionConfig.SessionName, sessionConfig.Store)
}

func (sessionConfig *SessionConfig) Get(c *gin.Context, sessionName string) interface{} {
	return sessions.Default(c).Get(sessionName)
}

func (sessionConfig *SessionConfig) Set(c *gin.Context, sessionName string, value interface{}) error {
	session := sessions.Default(c)
	session.Set(sessionName, value)
	return session.Save()
}

func (sessionConfig *SessionConfig) Delete(c *gin.Context, sessionName string) error {
	session := sessions.Default(c)
	session.Delete(sessionName)
	return session.Save()
}

func (sessionConfig *SessionConfig) Clear(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	return session.Save()
}

func (sessionConfig *SessionConfig) GetID(c *gin.Context) string {
	session := sessions.Default(c)
	return session.ID()
}
