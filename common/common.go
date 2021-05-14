package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"golang.org/x/crypto/bcrypt"
)

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func Success(c *gin.Context, httpStatus int, msg string, data interface{})  {
	c.JSON(httpStatus, gin.H{
		"msg": msg,
		"data": data,
		"timestamp":time.Now().Unix(),
	})
}

func GeneratePassword(userPassword string)(pass []byte,err error)  {
	return bcrypt.GenerateFromPassword([]byte(userPassword),bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string,hashed string) (isOK bool,err error)  {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(userPassword));err !=nil {
		return false,errors.New("密码比对错误！")
	}
	return true,nil
}





