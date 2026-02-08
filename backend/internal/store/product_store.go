package store

import (
	"fmt"
	"sync"
	"time"

	"shopping-demo/backend/internal/model"
)

// ProductStore 商品内存存储
type ProductStore struct {
	mu       sync.RWMutex       // 读写锁，保证并发安全
	products map[int]*model.Product // 商品数据映射
	nextID   int                    // 下一个可用的商品ID
}

// NewProductStore 创建商品存储实例并加载初始 mock 数据
func NewProductStore() *ProductStore {
	s := &ProductStore{
		products: make(map[int]*model.Product),
		nextID:   1,
	}
	s.loadMockData()
	return s
}

// loadMockData 加载初始 mock 商品数据
func (s *ProductStore) loadMockData() {
	now := time.Now()
	mockProducts := []model.Product{
		{
			Name:        "机械键盘",
			Description: "Cherry MX 红轴机械键盘，87键紧凑布局",
			Price:       399.00,
			Stock:       150,
			Category:    "电脑外设",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			Name:        "无线鼠标",
			Description: "人体工学设计，2.4G无线连接，续航持久",
			Price:       129.00,
			Stock:       300,
			Category:    "电脑外设",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			Name:        "显示器支架",
			Description: "铝合金材质，支持17-32寸显示器，可旋转升降",
			Price:       259.00,
			Stock:       80,
			Category:    "办公用品",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			Name:        "降噪耳机",
			Description: "主动降噪，蓝牙5.0，40小时续航",
			Price:       699.00,
			Stock:       200,
			Category:    "音频设备",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			Name:        "USB-C 扩展坞",
			Description: "Type-C接口，支持HDMI/USB3.0/SD卡/网口扩展",
			Price:       189.00,
			Stock:       120,
			Category:    "电脑配件",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	for _, p := range mockProducts {
		p.ID = s.nextID
		s.nextID++
		product := p // 避免闭包引用问题
		s.products[product.ID] = &product
	}
}

// List 获取所有商品列表
func (s *ProductStore) List() []*model.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	products := make([]*model.Product, 0, len(s.products))
	for _, p := range s.products {
		products = append(products, p)
	}
	return products
}

// GetByID 根据ID获取商品
func (s *ProductStore) GetByID(id int) (*model.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, ok := s.products[id]
	if !ok {
		return nil, fmt.Errorf("商品不存在: id=%d", id)
	}
	return p, nil
}

// Create 创建新商品
func (s *ProductStore) Create(req model.CreateProductRequest) *model.Product {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	p := &model.Product{
		ID:          s.nextID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	s.nextID++
	s.products[p.ID] = p
	return p
}

// Update 更新商品信息
func (s *ProductStore) Update(id int, req model.UpdateProductRequest) (*model.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p, ok := s.products[id]
	if !ok {
		return nil, fmt.Errorf("商品不存在: id=%d", id)
	}

	// 仅更新非空字段
	if req.Name != nil {
		p.Name = *req.Name
	}
	if req.Description != nil {
		p.Description = *req.Description
	}
	if req.Price != nil {
		p.Price = *req.Price
	}
	if req.Stock != nil {
		p.Stock = *req.Stock
	}
	if req.Category != nil {
		p.Category = *req.Category
	}
	p.UpdatedAt = time.Now()

	return p, nil
}

// Delete 删除商品
func (s *ProductStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.products[id]; !ok {
		return fmt.Errorf("商品不存在: id=%d", id)
	}
	delete(s.products, id)
	return nil
}
