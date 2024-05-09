package service

import (
	"bluebell/dao/mdb"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SingUp(p *models.ParamSignUp) (err error) {
	if err := mdb.CheckUserExists(p.Username); err != nil {
		return err
	}
	userId := snowflake.GenId()
	newUser := &models.User{
		UserId:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	return mdb.InsertUser(newUser)
}

func Login(p *models.ParamLogin) (u *models.User, err error) {
	u, err = mdb.GetUserByUsername(p.Username)
	if err != nil {
		return
	}
	if u.Password != mdb.EncryptPassword(p.Password) {
		return nil, mdb.ErrorPasswordInValid
	}
	u.Token, err = jwt.GenToken(u.UserId, u.Username)
	return
}
