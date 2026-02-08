import { useState, useEffect, useCallback } from "react";
import type { Product, ProductFormData } from "./types";
import {
  fetchProducts,
  createProduct,
  updateProduct,
  deleteProduct,
} from "./api";
import ProductTable from "./ProductTable";
import ProductForm from "./ProductForm";
import "./App.css";

// 主应用组件
function App() {
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // 模态框状态
  const [formVisible, setFormVisible] = useState(false);
  const [editingProduct, setEditingProduct] = useState<Product | null>(null);

  // 加载商品列表
  const loadProducts = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await fetchProducts();
      setProducts(data ?? []);
    } catch (err) {
      setError((err as Error).message);
    } finally {
      setLoading(false);
    }
  }, []);

  // 初始加载
  useEffect(() => {
    loadProducts();
  }, [loadProducts]);

  // 打开新增表单
  const handleAdd = () => {
    setEditingProduct(null);
    setFormVisible(true);
  };

  // 打开编辑表单
  const handleEdit = (product: Product) => {
    setEditingProduct(product);
    setFormVisible(true);
  };

  // 关闭表单
  const handleCancel = () => {
    setFormVisible(false);
    setEditingProduct(null);
  };

  // 提交表单（新增或编辑）
  const handleSubmit = async (data: ProductFormData) => {
    try {
      if (editingProduct) {
        await updateProduct(editingProduct.ID, data);
      } else {
        await createProduct(data);
      }
      setFormVisible(false);
      setEditingProduct(null);
      await loadProducts();
    } catch (err) {
      alert((err as Error).message);
    }
  };

  // 删除商品
  const handleDelete = async (id: number) => {
    try {
      await deleteProduct(id);
      await loadProducts();
    } catch (err) {
      alert((err as Error).message);
    }
  };

  return (
    <div className="app">
      <header className="header">
        <h1>商品管理系统</h1>
      </header>
      <main className="main">
        <div className="toolbar">
          <button className="btn btn-primary" onClick={handleAdd}>
            + 新增商品
          </button>
          <button className="btn btn-refresh" onClick={loadProducts}>
            刷新
          </button>
        </div>

        {error && <div className="error">错误：{error}</div>}
        {loading ? (
          <div className="loading">加载中...</div>
        ) : (
          <ProductTable
            products={products}
            onEdit={handleEdit}
            onDelete={handleDelete}
          />
        )}
      </main>

      <ProductForm
        visible={formVisible}
        initialData={editingProduct}
        onSubmit={handleSubmit}
        onCancel={handleCancel}
      />
    </div>
  );
}

export default App;
