package conf

import (
	"fmt"

	"github.com/go-akka/configuration"
)

type configModel struct {
	Port                       int32
	LoginProxyLocation         string
	LoginProxyPass             string
	LoginProxyUsernameJsonNode string
	AuthProxyPrefixPaths       []string
	AuthProxyPass              string
	AuthProxyUsernameJsonNode  string
	SecretKey                  string
	ValidityInMilliseconds     int64
	Issuer                     string
	Language                   string
	AccessLogPath              string
	ErrorLogPath               string
}

var conf = configuration.LoadConfig("auth.conf")

var AuthConfig = configModel{
	Port:                       conf.GetInt32("port", 8000),
	LoginProxyLocation:         conf.GetString("proxy.login.location"),
	LoginProxyPass:             conf.GetString("proxy.login.pass"),
	LoginProxyUsernameJsonNode: conf.GetString("proxy.login.username-json-node"),
	AuthProxyPrefixPaths:       conf.GetStringList("proxy.auth.location-prefixes"),
	AuthProxyPass:              conf.GetString("proxy.auth.pass"),
	AuthProxyUsernameJsonNode:  conf.GetString("proxy.auth.username-json-node"),
	SecretKey:                  conf.GetString("jwt.secretKey", "secretKeyForJWT"),
	ValidityInMilliseconds:     conf.GetInt64("jwt.validityInMilliseconds", 1800000),
	Issuer:                     conf.GetString("jwt.issuer", "issuer"),
	Language:                   conf.GetString("language", "cn"),
	AccessLogPath:              conf.GetString("log.access.path"),
	ErrorLogPath:               conf.GetString("log.error.path", "error.log"),
}

// TODO
func LoadConfig(filename string) *configuration.Config {
	defer func() {
		if info := recover(); info != nil {
			fmt.Println("配置文件错误:", info)
		}
	}()
	return configuration.LoadConfig(filename)
}
