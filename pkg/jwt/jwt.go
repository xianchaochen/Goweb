package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2 *365

var MySecret = []byte("夏天夏天悄悄过去")

func Generate(username string, userID int64) (aToken string, rToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userID,
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(MySecret)

	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(86700*7).Unix(), // 过期时间
		Issuer:    "bluebell",                              // 签发人
	}).SignedString(MySecret)
	return
}

func Parse(tokenString string) (*MyClaims, error) {
	var myClaims = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, myClaims, secret())
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return myClaims, nil
	}
	return nil, errors.New("invalid token")
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	}
}

func Refresh(aToken, rToken string) (newAccessToken string, newRefreshToken string, err error) {
	if _, err := jwt.Parse(rToken, secret()); err != nil {
		return "", "", err
	}

	claims := new(MyClaims)
	_, err = jwt.ParseWithClaims(aToken, claims, secret())
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		return Generate(claims.Username, claims.UserID)
	}

	return
}
