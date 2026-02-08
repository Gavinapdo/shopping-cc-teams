package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"shopping-demo/backend/internal/model"
	"shopping-demo/backend/internal/store"
)

// setupTestRouter 创建测试用的路由和处理器
func setupTestRouter() (*gin.Engine, *store.ProductStore) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	s := store.NewProductStore()
	h := NewProductHandler(s)

	api := r.Group("/api")
	products := api.Group("/products")
	{
		products.GET("", h.List)
		products.GET("/:id", h.GetByID)
		products.POST("", h.Create)
		products.PUT("/:id", h.Update)
		products.DELETE("/:id", h.Delete)
	}

	return r, s
}

// TestListProducts 测试获取商品列表接口
func TestListProducts(t *testing.T) {
	r, _ := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 %d，实际为 %d", http.StatusOK, w.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("解析响应失败: %v", err)
	}

	if resp["code"].(float64) != 0 {
		t.Errorf("期望 code=0，实际为 %v", resp["code"])
	}

	data, ok := resp["data"].([]interface{})
	if !ok {
		t.Fatal("响应中缺少 data 字段或类型错误")
	}

	// 初始 mock 数据应有5个商品
	if len(data) != 5 {
		t.Errorf("期望5个商品，实际为 %d", len(data))
	}
}

// TestGetProductByID 测试获取单个商品接口
func TestGetProductByID(t *testing.T) {
	r, _ := setupTestRouter()

	// 测试获取存在的商品
	req, _ := http.NewRequest("GET", "/api/products/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 %d，实际为 %d", http.StatusOK, w.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("解析响应失败: %v", err)
	}

	data, ok := resp["data"].(map[string]interface{})
	if !ok {
		t.Fatal("响应中缺少 data 字段或类型错误")
	}

	if data["id"].(float64) != 1 {
		t.Errorf("期望商品ID=1，实际为 %v", data["id"])
	}
}

// TestGetProductByID_NotFound 测试获取不存在的商品
func TestGetProductByID_NotFound(t *testing.T) {
	r, _ := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/products/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("期望状态码 %d，实际为 %d", http.StatusNotFound, w.Code)
	}
}

// TestGetProductByID_InvalidID 测试无效的商品ID
func TestGetProductByID_InvalidID(t *testing.T) {
	r, _ := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/products/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("期望状态码 %d，实际为 %d", http.StatusBadRequest, w.Code)
	}
}

// TestCreateProduct 测试创建商品接口
func TestCreateProduct(t *testing.T) {
	r, _ := setupTestRouter()

	newProduct := model.CreateProductRequest{
		Name:        "测试商品",
		Description: "这是一个测试商品",
		Price:       99.99,
		Stock:       50,
		Category:    "测试分类",
	}

	body, _ := json.Marshal(newProduct)
	req, _ := http.NewRequest("POST", "/api/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("期望状态码 %d，实际为 %d", http.StatusCreated, w.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("解析响应失败: %v", err)
	}

	data, ok := resp["data"].(map[string]interface{})
	if !ok {
		t.Fatal("响应中缺少 data 字段或类型错误")
	}

	if data["name"] != "测试商品" {
		t.Errorf("期望商品名称为 '测试商品'，实际为 '%v'", data["name"])
	}

	// 新创建的商品ID应为6（mock数据有5个）
	if data["id"].(float64) != 6 {
		t.Errorf("期望商品ID=6，实际为 %v", data["id"])
	}
}

// TestCreateProduct_InvalidRequest 测试创建商品时参数校验
func TestCreateProduct_InvalidRequest(t *testing.T) {
	r, _ := setupTestRouter()

	// 缺少必填字段
	invalidProduct := map[string]interface{}{
		"description": "缺少名称和价格",
	}

	body, _ := json.Marshal(invalidProduct)
	req, _ := http.NewRequest("POST", "/api/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("期望状态码 %d，实际为 %d", http.StatusBadRequest, w.Code)
	}
}

// TestUpdateProduct 测试更新商品接口
func TestUpdateProduct(t *testing.T) {
	r, _ := setupTestRouter()

	newName := "更新后的名称"
	newPrice := 599.00
	updateReq := model.UpdateProductRequest{
		Name:  &newName,
		Price: &newPrice,
	}

	body, _ := json.Marshal(updateReq)
	req, _ := http.NewRequest("PUT", "/api/products/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 %d，实际为 %d", http.StatusOK, w.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("解析响应失败: %v", err)
	}

	data, ok := resp["data"].(map[string]interface{})
	if !ok {
		t.Fatal("响应中缺少 data 字段或类型错误")
	}

	if data["name"] != "更新后的名称" {
		t.Errorf("期望商品名称为 '更新后的名称'，实际为 '%v'", data["name"])
	}

	if data["price"].(float64) != 599.00 {
		t.Errorf("期望商品价格为 599.00，实际为 %v", data["price"])
	}
}

// TestUpdateProduct_NotFound 测试更新不存在的商品
func TestUpdateProduct_NotFound(t *testing.T) {
	r, _ := setupTestRouter()

	newName := "不存在的商品"
	updateReq := model.UpdateProductRequest{
		Name: &newName,
	}

	body, _ := json.Marshal(updateReq)
	req, _ := http.NewRequest("PUT", "/api/products/999", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("期望状态码 %d，实际为 %d", http.StatusNotFound, w.Code)
	}
}

// TestDeleteProduct 测试删除商品接口
func TestDeleteProduct(t *testing.T) {
	r, _ := setupTestRouter()

	req, _ := http.NewRequest("DELETE", "/api/products/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 %d，实际为 %d", http.StatusOK, w.Code)
	}

	// 验证商品已被删除
	req2, _ := http.NewRequest("GET", "/api/products/1", nil)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)

	if w2.Code != http.StatusNotFound {
		t.Errorf("删除后期望状态码 %d，实际为 %d", http.StatusNotFound, w2.Code)
	}
}

// TestDeleteProduct_NotFound 测试删除不存在的商品
func TestDeleteProduct_NotFound(t *testing.T) {
	r, _ := setupTestRouter()

	req, _ := http.NewRequest("DELETE", "/api/products/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("期望状态码 %d，实际为 %d", http.StatusNotFound, w.Code)
	}
}

// TestDeleteProduct_InvalidID 测试删除时使用无效ID
func TestDeleteProduct_InvalidID(t *testing.T) {
	r, _ := setupTestRouter()

	req, _ := http.NewRequest("DELETE", "/api/products/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("期望状态码 %d，实际为 %d", http.StatusBadRequest, w.Code)
	}
}
