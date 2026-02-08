# 商品管理 API 文档

## 基本信息

- **基础URL**: `http://localhost:8080`
- **数据格式**: JSON
- **字符编码**: UTF-8

## 通用响应格式

所有接口返回统一的 JSON 格式：

```json
{
  "code": 0,
  "message": "操作描述",
  "data": {}
}
```

- `code`: 0 表示成功，非0表示错误
- `message`: 操作结果描述
- `data`: 返回的数据（部分接口无此字段）

---

## 接口列表

### 1. 获取商品列表

**请求**

```
GET /api/products
```

**响应示例**

```json
{
  "code": 0,
  "message": "获取商品列表成功",
  "data": [
    {
      "id": 1,
      "name": "机械键盘",
      "description": "Cherry MX 红轴机械键盘，87键紧凑布局",
      "price": 399.00,
      "stock": 150,
      "category": "电脑外设",
      "created_at": "2025-01-01T10:00:00Z",
      "updated_at": "2025-01-01T10:00:00Z"
    },
    {
      "id": 2,
      "name": "无线鼠标",
      "description": "人体工学设计，2.4G无线连接，续航持久",
      "price": 129.00,
      "stock": 300,
      "category": "电脑外设",
      "created_at": "2025-01-01T10:00:00Z",
      "updated_at": "2025-01-01T10:00:00Z"
    }
  ]
}
```

---

### 2. 获取单个商品

**请求**

```
GET /api/products/:id
```

**路径参数**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id   | int  | 是   | 商品ID |

**成功响应 (200)**

```json
{
  "code": 0,
  "message": "获取商品成功",
  "data": {
    "id": 1,
    "name": "机械键盘",
    "description": "Cherry MX 红轴机械键盘，87键紧凑布局",
    "price": 399.00,
    "stock": 150,
    "category": "电脑外设",
    "created_at": "2025-01-01T10:00:00Z",
    "updated_at": "2025-01-01T10:00:00Z"
  }
}
```

**失败响应 (404)**

```json
{
  "code": 404,
  "message": "商品不存在: id=999"
}
```

**失败响应 (400)**

```json
{
  "code": 400,
  "message": "无效的商品ID"
}
```

---

### 3. 创建商品

**请求**

```
POST /api/products
Content-Type: application/json
```

**请求体**

| 字段        | 类型    | 必填 | 说明                 |
|-------------|---------|------|----------------------|
| name        | string  | 是   | 商品名称             |
| description | string  | 否   | 商品描述             |
| price       | float64 | 是   | 商品价格（必须大于0）|
| stock       | int     | 是   | 库存数量（必须>=0）  |
| category    | string  | 是   | 商品分类             |

**请求示例**

```json
{
  "name": "蓝牙音箱",
  "description": "便携式蓝牙音箱，IPX7防水",
  "price": 299.00,
  "stock": 100,
  "category": "音频设备"
}
```

**成功响应 (201)**

```json
{
  "code": 0,
  "message": "创建商品成功",
  "data": {
    "id": 6,
    "name": "蓝牙音箱",
    "description": "便携式蓝牙音箱，IPX7防水",
    "price": 299.00,
    "stock": 100,
    "category": "音频设备",
    "created_at": "2025-01-01T12:00:00Z",
    "updated_at": "2025-01-01T12:00:00Z"
  }
}
```

**失败响应 (400)**

```json
{
  "code": 400,
  "message": "请求参数错误: Key: 'CreateProductRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
}
```

---

### 4. 更新商品

**请求**

```
PUT /api/products/:id
Content-Type: application/json
```

**路径参数**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id   | int  | 是   | 商品ID |

**请求体**（所有字段均为可选，仅更新提供的字段）

| 字段        | 类型    | 必填 | 说明     |
|-------------|---------|------|----------|
| name        | string  | 否   | 商品名称 |
| description | string  | 否   | 商品描述 |
| price       | float64 | 否   | 商品价格 |
| stock       | int     | 否   | 库存数量 |
| category    | string  | 否   | 商品分类 |

**请求示例**

```json
{
  "name": "机械键盘（升级版）",
  "price": 459.00
}
```

**成功响应 (200)**

```json
{
  "code": 0,
  "message": "更新商品成功",
  "data": {
    "id": 1,
    "name": "机械键盘（升级版）",
    "description": "Cherry MX 红轴机械键盘，87键紧凑布局",
    "price": 459.00,
    "stock": 150,
    "category": "电脑外设",
    "created_at": "2025-01-01T10:00:00Z",
    "updated_at": "2025-01-01T14:00:00Z"
  }
}
```

**失败响应 (404)**

```json
{
  "code": 404,
  "message": "商品不存在: id=999"
}
```

---

### 5. 删除商品

**请求**

```
DELETE /api/products/:id
```

**路径参数**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id   | int  | 是   | 商品ID |

**成功响应 (200)**

```json
{
  "code": 0,
  "message": "删除商品成功"
}
```

**失败响应 (404)**

```json
{
  "code": 404,
  "message": "商品不存在: id=999"
}
```

**失败响应 (400)**

```json
{
  "code": 400,
  "message": "无效的商品ID"
}
```

---

## 错误码说明

| HTTP状态码 | 说明           |
|------------|----------------|
| 200        | 请求成功       |
| 201        | 创建成功       |
| 400        | 请求参数错误   |
| 404        | 资源不存在     |

## CORS 配置

API 已配置 CORS 中间件，支持以下设置：

- **允许来源**: 所有来源（`*`）
- **允许方法**: GET, POST, PUT, DELETE, OPTIONS
- **允许请求头**: Origin, Content-Type, Accept, Authorization
- **预检缓存**: 12小时
