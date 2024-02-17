package main

import (
	"authorization/conf"
	"authorization/handler"
	"authorization/jwt"
	"authorization/logger"
	"authorization/proxy"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	for _, path := range conf.AuthConfig.AuthProxyPrefixPaths {
		r.PathPrefix(path).HandlerFunc(proxy.HttpProxy)
	}
	r.PathPrefix("/").HandlerFunc(handler.Index)
	r.Use(jwt.TokenMiddleware)
	http.Handle("/", r)

	http.HandleFunc(conf.AuthConfig.LoginProxyLocation, proxy.LoginProxy)
	logger.Access.Printf("监听端口：%d\n", conf.AuthConfig.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.AuthConfig.Port), nil)
	if err != nil {
		logger.Error.Fatalln("启动错误", err)
	}
	logger.Access.Printf("HTTP服务启动成功，监听端口：%d\n", conf.AuthConfig.Port)
}
