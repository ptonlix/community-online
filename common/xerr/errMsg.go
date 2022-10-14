package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	// 用户模块
	message[USER_MESSAGE_ERROR] = "短信验证码验证失败,请检查重试"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[USER_LOGIN_VERIFY_ERROR] = "账号或密码验证失败"
	message[CAPTCHA_GENERATE_ERROR] = "图形校验码生成失败"
	message[CAPTCHA_VERIFY_ERROR] = "图形校验码验证失败"
	message[USER_NOT_EXIST_ERROR] = "该用户不存在"
	message[USER_ALREADY_REGISTER_ERROR] = "该用户已注册"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
