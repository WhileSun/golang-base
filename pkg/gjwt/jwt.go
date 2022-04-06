package gjwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/whilesun/go-admin/pkg/gconf"
	"time"
)

type Jwt struct {
	Secret string
	Exp    int64
}

var jwtConfig *Jwt

func init() {
	jwtConfig = &Jwt{
		Secret: "jwt256_go_oms",
		Exp:    int64(3600 * 6),
	}
	gconf.Config.UnmarshalKey("jwt", jwtConfig)
}

func CreateToken(values jwt.MapClaims) (string, error) {
	values["exp"] = time.Now().Unix() + jwtConfig.Exp
	values["iat"] = time.Now().Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, values)
	token, err := at.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	//检测加密方式是否一致
	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		err, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return err, nil
		}
		return []byte(jwtConfig.Secret), nil
	})
	if !token.Valid {
		return nil, tokenErr
	} else {
		claims := token.Claims.(jwt.MapClaims)
		return claims, nil
	}
}
