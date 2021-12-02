// Package e
/**
  * @Author:manguanghui
  * @Date: 2021/12/1
  * @Desc: 状态码
**/
package e

import "net/http"

const (
	SUCCESS              = http.StatusOK
	FAILED               = http.StatusBadRequest
	ERROR                = http.StatusInternalServerError
	INVALID              = http.StatusUnauthorized
	PermissionDenied     = http.StatusForbidden
	NOTFOUND             = http.StatusNotFound
	ErrSuccess           = 100001
	ErrUnknown           = 100002
	ErrBind              = 100003
	ErrValidation        = 100004
	ErrTokenInvalid      = 100005
	ErrAdminNotFound     = 100007
	ErrTokenExpired      = 100203
	ErrDatabase          = 100101
	ErrInvalidAuthHeader = 100204
	ErrPasswordIncorrect = 100206
	ErrMissingHeader     = 100205
	ErrPermissionDenied  = 100207
	ErrEncodingJSON      = 100303
	ErrDecodingJSON      = 100305
	ErrStrToInt64        = 100306
	ErrUserNotFound      = 110001
	ErrUserAlreadyExist  = 110002
)
