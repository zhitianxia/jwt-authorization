package lang

var cn = map[string]string{
	"ok":"正常",
	"Unauthorized":"验证失败：未登录",
	"ValidationErrorExpired":"验证失败：令牌已过期。请重新登录！",
	"ValidationErrorNotValidYet":"验证失败：令牌未生效。请重新登录！",
	"ValidationErrorIssuer":"验证失败：令牌不信任。请重新登录！",
	"ValidationErrorMalformed":"验证失败：令牌格式非法。请重新登录！",
	"ValidationErrorUser":"验证失败：用户名非法",
	"LoginProxyPassError":"登录失败：代理配置有误",
}
