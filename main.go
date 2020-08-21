package main

import (
	"context"
	"fmt"
	"go-gin-example/pkg/logging"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//router := gin.Default()	// 创建一个路由Handlers,可以后期绑定各类的路由规则和函数、中间件等
	//router.GET("/test", func(c *gin.Context) {		// Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证 JSON 请求、响应 JSON 请求等
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})

	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endpoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//
	//server := endless.NewServer(endpoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	logging.Info("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	logging.Info("server err: %v", err)
	//}

	router := routers.InitRouter()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d",setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logging.Info("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	 <- quit

	 logging.Info("Shutdown Server...")

	 ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	 defer cancel()
	 if err := s.Shutdown(ctx); err != nil {
	 	logging.Info("Server shutdown:", err)
	 }

	 logging.Info("Server exiting")

	//s.ListenAndServe()
}