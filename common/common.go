package common

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func FormatTranslateMsg(fields map[string]string) string {
	var builder strings.Builder
	for _, err := range fields {
		fmt.Println(fmt.Sprintf("%+v\n", err))
		err = strings.Trim(err, "\"")
		if builder.Len() == 0 {
			builder.WriteString(err)
		} else {
			builder.WriteString(","+err)
		}
	}

	return 	strings.Trim(builder.String(), "\"")
}

func GeneratePassword(userPassword string) (pass []byte, err error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil
}


