package mdb

import "errors"

var (
	ErrorUserExists      = errors.New("用户已存在")
	ErrorPasswordInValid = errors.New("密码错误")
)
