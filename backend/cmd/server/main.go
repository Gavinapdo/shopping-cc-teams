package main

import (
	"log"

	"shopping-demo/backend/internal/router"
)

func main() {
	// 初始化路由
	r := router.Setup()

	// 启动HTTP服务，监听8080端口
	log.Println("商品管理API服务启动，监听端口 :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
