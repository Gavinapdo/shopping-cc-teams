import type { Product, ProductFormData } from "./types";

// 后端 API 基础地址
const API_BASE = "http://localhost:8080/api/products";

// 获取所有商品
export async function fetchProducts(): Promise<Product[]> {
  const res = await fetch(API_BASE);
  if (!res.ok) throw new Error("获取商品列表失败");
  return res.json();
}

// 获取单个商品
export async function fetchProduct(id: number): Promise<Product> {
  const res = await fetch(`${API_BASE}/${id}`);
  if (!res.ok) throw new Error("获取商品详情失败");
  return res.json();
}

// 创建商品
export async function createProduct(data: ProductFormData): Promise<Product> {
  const res = await fetch(API_BASE, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("创建商品失败");
  return res.json();
}

// 更新商品
export async function updateProduct(
  id: number,
  data: ProductFormData
): Promise<Product> {
  const res = await fetch(`${API_BASE}/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error("更新商品失败");
  return res.json();
}

// 删除商品
export async function deleteProduct(id: number): Promise<void> {
  const res = await fetch(`${API_BASE}/${id}`, {
    method: "DELETE",
  });
  if (!res.ok) throw new Error("删除商品失败");
}
