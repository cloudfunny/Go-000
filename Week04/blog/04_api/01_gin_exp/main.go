package main

import "github.com/gin-gonic/gin"

func handler(ctx *gin.Context) {
	// get params
	params := struct {
		Msg string `json:"msg"`
	}{}
	ctx.BindQuery(&params)

	// 业务逻辑

	// 返回数据
	ctx.JSON(200, gin.H{
		"message": params.Msg,
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", handler)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
