package controller

import "C"
import (
	"bluebell/dao/mdb"
	"bluebell/models"
	"bluebell/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SingUp with invalid params", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ResponseWithFail(c, CodeInvalidParams)
			return
		}
		ResponseWithFailMsg(c, CodeInvalidParams, RemoveTopStruct(errs.Translate(Trans)))
		return
	}
	err := service.SingUp(p)
	if err != nil {
		if errors.Is(err, mdb.ErrorUserExists) {
			ResponseWithFail(c, CodeUserExists)
		} else {
			ResponseWithFail(c, CodeServerBusy)
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SingUp with invalid params", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ResponseWithFail(c, CodeInvalidParams)
		} else {
			ResponseWithFailMsg(c, CodeInvalidParams, RemoveTopStruct(errs.Translate(Trans)))
		}
		return
	}
	u, err := service.Login(p)
	if err != nil {
		zap.L().Error("service.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseWithFail(c, CodeUserNotExists)
			return
		} else {
			ResponseWithFail(c, CodePasswordInValid)
			return
		}
	}
	ResponseWithSuccess(c, gin.H{
		"userId":   u.UserId,
		"username": u.Username,
		"token":    u.Token,
	})
}
