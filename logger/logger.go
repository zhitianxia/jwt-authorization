package logger

import (
	"authorization/conf"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace *log.Logger //调试日志
	Access *log.Logger //请求日志
	Error *log.Logger //错误日志
)

func init() {
	errorFile,err := os.OpenFile(conf.AuthConfig.ErrorLogPath,os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil{
		log.Fatalln("打开日志文件失败",err)
	}else {
		log.Printf("错误日志位置：%s",conf.AuthConfig.ErrorLogPath)
	}

	if conf.AuthConfig.AccessLogPath == "" {
		Access = log.New(os.Stdout,
			"INFO:",log.Ldate|log.Ltime)
	}else{
		accessFile,err := os.OpenFile(conf.AuthConfig.AccessLogPath,os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil{
			log.Fatalln("打开日志文件失败",err)
		}else{
			log.Printf("访问日志位置：%s",conf.AuthConfig.AccessLogPath)
		}
		Access = log.New(io.MultiWriter(accessFile,os.Stdout),
			"Info:",log.Ldate|log.Ltime)
	}

	Trace = log.New(ioutil.Discard,
		"Trace:",log.Ldate|log.Ltime)

	Error = log.New(io.MultiWriter(errorFile,os.Stderr),
		"Error:",log.Ldate|log.Ltime)
}
