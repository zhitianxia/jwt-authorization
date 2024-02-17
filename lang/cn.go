package lang

var cn = map[string]string{
	"ok":                         "验证通过：但是当前请求没有匹配的代理服务",
	"Unauthorized":               "验证失败：未登录",
	"ValidationErrorExpired":     "验证失败：令牌已过期。请重新登录！",
	"ValidationErrorNotValidYet": "验证失败：令牌未生效。请重新登录！",
	"ValidationErrorIssuer":      "验证失败：令牌不信任。请重新登录！",
	"ValidationErrorMalformed":   "验证失败：令牌格式非法。请重新登录！",
	"ValidationErrorUser":        "验证失败：用户名非法",
	"LoginProxyPassError":        "登录失败：代理配置有误",
}
