package mdb

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
)

func CheckUserExists(username string) (err error) {
	var count int64
	result := db.Model(&models.User{}).Where(&models.User{Username: username}).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count > 0 {
		return ErrorUserExists
	}
	return
}

func InsertUser(user *models.User) (err error) {
	user.Password = EncryptPassword(user.Password)
	tx := db.Create(user)
	err = tx.Error
	return
}

const secret = "bluebell"

func EncryptPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum([]byte(password)))
}

func GetUserByUsername(username string) (u *models.User, err error) {
	u = new(models.User)
	tx := db.Where("username = ?", username).First(u)
	err = tx.Error
	return
}
