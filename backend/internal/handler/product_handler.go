package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"shopping-demo/backend/internal/model"
	"shopping-demo/backend/internal/store"
)

// ProductHandler 商品HTTP处理器
type ProductHandler struct {
	store *store.ProductStore // 商品存储实例
}

// NewProductHandler 创建商品处理器实例
func NewProductHandler(s *store.ProductStore) *ProductHandler {
	return &ProductHandler{store: s}
}

// List 获取商品列表
// GET /api/products
func (h *ProductHandler) List(c *gin.Context) {
	products := h.store.List()
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取商品列表成功",
		"data":    products,
	})
}

// GetByID 获取单个商品
// GET /api/products/:id
func (h *ProductHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的商品ID",
		})
		return
	}

	product, err := h.store.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取商品成功",
		"data":    product,
	})
}

// Create 创建商品
// POST /api/products
func (h *ProductHandler) Create(c *gin.Context) {
	var req model.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误: " + err.Error(),
		})
		return
	}

	product := h.store.Create(req)
	c.JSON(http.StatusCreated, gin.H{
		"code":    0,
		"message": "创建商品成功",
		"data":    product,
	})
}

// Update 更新商品
// PUT /api/products/:id
func (h *ProductHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的商品ID",
		})
		return
	}

	var req model.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误: " + err.Error(),
		})
		return
	}

	product, err := h.store.Update(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "更新商品成功",
		"data":    product,
	})
}

// Delete 删除商品
// DELETE /api/products/:id
func (h *ProductHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的商品ID",
		})
		return
	}

	if err := h.store.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除商品成功",
	})
}
