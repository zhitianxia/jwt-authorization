package lang

var langMap = map[string]map[string]string{
	"cn": cn,
	//"en": en,
}

func Get(value string) string{
	langKey := "cn"
	if msg,ok :=langMap[langKey][value];ok {
		return msg
	}
	return value
}
