// Package e
/**
  * @Author:manguanghui
  * @Date: 2021/12/1
  * @Desc: 状态码对应信息
**/
package e

var MsgFlags = map[int32]string{
	ErrSuccess:           "成功",
	FAILED:               "前端传入参数不完整",
	ERROR:                "Fail",
	ErrUserNotFound:      "用户不存在",
	ErrUserAlreadyExist:  "用户已存在",
	ErrUnknown:           "内部服务器错误",
	ErrBind:              "在将请求体绑定到结构体时发生错误",
	ErrValidation:        "验证失败",
	ErrTokenInvalid:      "Token 无效",
	ErrAdminNotFound:     "admin 不存在",
	ErrTokenExpired:      "Token 过期",
	ErrDatabase:          "数据库错误",
	ErrInvalidAuthHeader: "无效的授权头",
	ErrPasswordIncorrect: "密码不正确",
	ErrMissingHeader:     "授权头为空",
	ErrPermissionDenied:  "权限不够,拒绝访问",
	ErrEncodingJSON:      "JSON数据无法编码",
	ErrDecodingJSON:      "JSON数据无法解码",
	ErrStrToInt64:        "string转int64错误",
}

func GetMsg(code int32) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
