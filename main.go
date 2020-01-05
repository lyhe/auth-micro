package main

import (
	"github.com/lyhe/auth-micro/router"
	"github.com/micro/go-web"
	"github.com/prometheus/common/log"
	"net/http"
	"time"
)

func main() {
	// 创建 micro 服务
	service := web.NewService(
		web.Name("go.micro.api.lyhe"),
	)
	_ = service.Init()

	// 创建gin路由
	r := router.Router()
	s := http.Server{
		Addr:           ":8888",
		Handler:        r,
		ReadTimeout:    time.Second * 30,
		WriteTimeout:   time.Second * 30,
		MaxHeaderBytes: 1 << 20,
	}

	// 用gin注册go-micro handler
	service.Handle("/", r)

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("run gin error: %s\n", err)
	}
	if err := service.Run(); err != nil {
		log.Fatalf("run go-micro error: %s\n", err)
	}
}
