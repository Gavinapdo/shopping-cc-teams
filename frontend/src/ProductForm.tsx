import { useState, useEffect } from "react";
import type { Product, ProductFormData } from "./types";

// 商品表单组件的属性类型
interface ProductFormProps {
  visible: boolean;
  initialData?: Product | null;
  onSubmit: (data: ProductFormData) => void;
  onCancel: () => void;
}

// 空表单初始值
const emptyForm: ProductFormData = {
  Name: "",
  Description: "",
  Price: 0,
  Stock: 0,
  Category: "",
};

// 商品表单模态框组件
export default function ProductForm({
  visible,
  initialData,
  onSubmit,
  onCancel,
}: ProductFormProps) {
  const [form, setForm] = useState<ProductFormData>(emptyForm);

  // 当编辑数据变化时，填充表单
  useEffect(() => {
    if (initialData) {
      setForm({
        Name: initialData.Name,
        Description: initialData.Description,
        Price: initialData.Price,
        Stock: initialData.Stock,
        Category: initialData.Category,
      });
    } else {
      setForm(emptyForm);
    }
  }, [initialData, visible]);

  if (!visible) return null;

  const isEdit = !!initialData;

  // 处理表单提交
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(form);
  };

  return (
    <div className="modal-overlay" onClick={onCancel}>
      <div className="modal" onClick={(e) => e.stopPropagation()}>
        <h2>{isEdit ? "编辑商品" : "新增商品"}</h2>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label>名称</label>
            <input
              type="text"
              required
              value={form.Name}
              onChange={(e) => setForm({ ...form, Name: e.target.value })}
            />
          </div>
          <div className="form-group">
            <label>描述</label>
            <textarea
              value={form.Description}
              onChange={(e) =>
                setForm({ ...form, Description: e.target.value })
              }
            />
          </div>
          <div className="form-row">
            <div className="form-group">
              <label>价格</label>
              <input
                type="number"
                required
                min="0"
                step="0.01"
                value={form.Price}
                onChange={(e) =>
                  setForm({ ...form, Price: parseFloat(e.target.value) || 0 })
                }
              />
            </div>
            <div className="form-group">
              <label>库存</label>
              <input
                type="number"
                required
                min="0"
                value={form.Stock}
                onChange={(e) =>
                  setForm({ ...form, Stock: parseInt(e.target.value) || 0 })
                }
              />
            </div>
          </div>
          <div className="form-group">
            <label>分类</label>
            <input
              type="text"
              value={form.Category}
              onChange={(e) => setForm({ ...form, Category: e.target.value })}
            />
          </div>
          <div className="form-actions">
            <button type="button" className="btn btn-cancel" onClick={onCancel}>
              取消
            </button>
            <button type="submit" className="btn btn-primary">
              {isEdit ? "保存" : "创建"}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
