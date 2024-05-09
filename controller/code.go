package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParams
	CodeServerBusy
	CodeUserNotExists
	CodeUserExists
	CodePasswordInValid
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeServerBusy:      "业务繁忙",
	CodeUserNotExists:   "用户不存在",
	CodeUserExists:      "用户存在",
	CodePasswordInValid: "密码错误",
}

func (r ResCode) Msg() string {
	s, ok := codeMsgMap[r]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return s
}
