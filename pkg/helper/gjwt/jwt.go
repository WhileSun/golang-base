package gjwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/whilesun/go-admin/pkg/core/gconfig"
	"time"
)

type JwtConfig struct {
	Secret  string
	Exp     int64
	LastExp int64
	Version float64
}

func New(jwtKey string) *JwtConfig {
	if jwtKey == "" {
		jwtKey = "jwt"
	}
	jwtConfig := &JwtConfig{
		Secret:  "jwt256_go_oms",
		Exp:     int64(1800),
		Version: 1.0,
	}
	gconfig.Config.UnmarshalKey(jwtKey, jwtConfig)
	// 默认失效后半天内访问自动续期
	if jwtConfig.LastExp == 0 {
		jwtConfig.LastExp = jwtConfig.Exp + 43200
	}
	return jwtConfig
}

func (jwtConfig *JwtConfig) CreateToken(values jwt.MapClaims) (string, error) {
	values["exp"] = time.Now().Unix() + jwtConfig.Exp
	values["lastexp"] = time.Now().Unix() + jwtConfig.LastExp
	values["iat"] = time.Now().Unix()
	values["version"] = jwtConfig.Version
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, values)
	token, err := at.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (jwtConfig *JwtConfig) ParseToken(tokenString string) (jwt.MapClaims, error) {
	//检测加密方式是否一致
	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		err, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return err, nil
		}
		return []byte(jwtConfig.Secret), nil
	})
	if token == nil {
		return nil, tokenErr
	}
	if !token.Valid {
		if token.Claims == nil {
			return nil, tokenErr
		}
		claims := token.Claims.(jwt.MapClaims)
		return claims, tokenErr
	} else {
		claims := token.Claims.(jwt.MapClaims)
		return claims, nil
	}
}
