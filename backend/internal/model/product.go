package model

import "time"

// Product 商品数据模型
type Product struct {
	ID          int       `json:"id"`          // 商品ID
	Name        string    `json:"name"`        // 商品名称
	Description string    `json:"description"` // 商品描述
	Price       float64   `json:"price"`       // 商品价格
	Stock       int       `json:"stock"`       // 库存数量
	Category    string    `json:"category"`    // 商品分类
	CreatedAt   time.Time `json:"created_at"`  // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`  // 更新时间
}

// CreateProductRequest 创建商品请求
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`        // 商品名称（必填）
	Description string  `json:"description"`                    // 商品描述
	Price       float64 `json:"price" binding:"required,gt=0"`  // 商品价格（必填，大于0）
	Stock       int     `json:"stock" binding:"required,gte=0"` // 库存数量（必填，大于等于0）
	Category    string  `json:"category" binding:"required"`    // 商品分类（必填）
}

// UpdateProductRequest 更新商品请求
type UpdateProductRequest struct {
	Name        *string  `json:"name"`        // 商品名称
	Description *string  `json:"description"` // 商品描述
	Price       *float64 `json:"price"`       // 商品价格
	Stock       *int     `json:"stock"`       // 库存数量
	Category    *string  `json:"category"`    // 商品分类
}
