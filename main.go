package main

import (
	"fmt"
	"github.com/gin-gonic/gin" // 导入web服务框架
)

func main() {
	err := Start(fmt.Sprintf("%s:%s", "127.0.0.1", "7777"), "Static")
	if err != nil {
		fmt.Printf("web服务错误:%v\n", err)
		return
	}
}

func Start(addr, webDir string) (err error) {
	// 使用gin框架提供的默认web服务引擎
	r := gin.Default()
	// 静态文件服务
	if len(webDir) > 0 {
		// 将一个目录下的静态文件，并注册到web服务器
		r.Static("/web", webDir)
	}

	// api接口服务，定义了路由组
	todo := r.Group("todo")
	{
		// 定义增改查的接口，并注册到web服务器
		todo.GET("", ListTodoHandler)
		todo.POST("", InsertTodoHandler)
		todo.PUT("", UpdateTodoHandler)
		todo.DELETE("", DeleteTodoHandler)
	}

	// 启动web服务
	err = r.Run(addr)
	return err
}
