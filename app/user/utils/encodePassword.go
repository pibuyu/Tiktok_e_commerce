package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/consts"
)

func EncodePassword(email, password string) string {
	salt := fmt.Sprintf("%s%s", consts.SECRET_SALT, email)
	saltPwd := []byte(fmt.Sprintf("%s%s", salt, password))
	passwordMd5 := fmt.Sprintf("%x", md5.Sum(saltPwd))

	return passwordMd5
}
