package proxy

import (
	"authorization/conf"
	"authorization/jwt"
	"authorization/logger"
	"authorization/util"
	"bytes"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/thedevsaddam/gojsonq/v2"
)

func HttpProxy(w http.ResponseWriter, r *http.Request) {
	proxyPass := r.Header.Get("X-Proxy-Pass")
	if proxyPass == "" {
		proxyPass = conf.AuthConfig.AuthProxyPass
	}
	u, _ := url.Parse(proxyPass)
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
	logger.Trace.Printf("代理请求完成：%s", r.RequestURI)
}

func LoginProxy(w http.ResponseWriter, r *http.Request) {
	loginProxyPass := conf.AuthConfig.LoginProxyPass
	usernameNode := conf.AuthConfig.LoginProxyUsernameJsonNode
	logger.Trace.Printf("loginProxyPass:%s;；usernameNode:%s", loginProxyPass, usernameNode)
	if loginProxyPass == "" {
		logger.Error.Fatalf("未配置登录代理地址。url[%s]\n", r.RequestURI)
		util.ResponseWithProblem(w, 400, "LoginProxyPassError")
		return
	}

	u, _ := url.Parse(loginProxyPass)
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ErrorLog = logger.Error
	proxy.ModifyResponse = func(response *http.Response) error {
		var username string
		if usernameNode != "" {
			content, err := io.ReadAll(response.Body)
			if err != nil {
				logger.Error.Println("读取代理返回出错:", err)
			}
			response.Body = io.NopCloser(bytes.NewBuffer(content))
			usernameNodeValue := gojsonq.New().FromString(string(content)).Find(usernameNode)
			if usernameNodeValue == nil {
				logger.Error.Printf("登录不成功，代理返回错误:%s。url[%s]\n", string(content), r.RequestURI)
				return nil
			}
			username = usernameNodeValue.(string)
			logger.Access.Printf("登录用户名为：%s，url[%s]\n", username, r.RequestURI)
		}
		jwt.SetAuthorizationHeader(response, username)
		return nil
	}
	proxy.ServeHTTP(w, r)
}
