package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"shopping-demo/backend/internal/handler"
	"shopping-demo/backend/internal/store"
)

// Setup 初始化并返回配置好的 Gin 引擎
func Setup() *gin.Engine {
	r := gin.Default()

	// 配置 CORS 中间件，允许所有来源的跨域请求
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// 初始化存储和处理器
	productStore := store.NewProductStore()
	productHandler := handler.NewProductHandler(productStore)

	// 注册商品相关路由
	api := r.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", productHandler.List)        // 获取商品列表
			products.GET("/:id", productHandler.GetByID) // 获取单个商品
			products.POST("", productHandler.Create)     // 创建商品
			products.PUT("/:id", productHandler.Update)  // 更新商品
			products.DELETE("/:id", productHandler.Delete) // 删除商品
		}
	}

	return r
}
