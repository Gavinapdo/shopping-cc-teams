// 商品数据类型定义
export interface Product {
  ID: number;
  Name: string;
  Description: string;
  Price: number;
  Stock: number;
  Category: string;
  CreatedAt: string;
  UpdatedAt: string;
}

// 创建/更新商品时的表单数据类型
export interface ProductFormData {
  Name: string;
  Description: string;
  Price: number;
  Stock: number;
  Category: string;
}
