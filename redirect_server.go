package main

import (
	"github.com/gin-gonic/gin"
)

// CodeCh 用于传输服务器返回的client-code
var CodeCh = make(chan *string)

func StartServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/auth", authMs)
	err := router.Run(":11451")
	if err != nil {
		return
	}
}

func authMs(c *gin.Context) {
	code := c.Query("code")
	if len(code) != 0 {
		CodeCh <- &code
		c.String(200, "获取client-code成功，请回到工具执行下一步操作\r\n\r\n%s" , code)
	}else{
		CodeCh <- nil
		c.String(200, "获取client-code失败，请检查网络设置并再次运行")
	}
}
