import type { Product } from "./types";

// 商品表格组件的属性类型
interface ProductTableProps {
  products: Product[];
  onEdit: (product: Product) => void;
  onDelete: (id: number) => void;
}

// 格式化日期时间
function formatDate(dateStr: string): string {
  if (!dateStr) return "-";
  const d = new Date(dateStr);
  return d.toLocaleString("zh-CN");
}

// 商品列表表格组件
export default function ProductTable({
  products,
  onEdit,
  onDelete,
}: ProductTableProps) {
  if (products.length === 0) {
    return <div className="empty">暂无商品数据</div>;
  }

  return (
    <div className="table-wrapper">
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>名称</th>
            <th>描述</th>
            <th>价格</th>
            <th>库存</th>
            <th>分类</th>
            <th>创建时间</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          {products.map((p) => (
            <tr key={p.ID}>
              <td>{p.ID}</td>
              <td>{p.Name}</td>
              <td className="desc-cell">{p.Description}</td>
              <td>¥{p.Price.toFixed(2)}</td>
              <td>{p.Stock}</td>
              <td>{p.Category}</td>
              <td>{formatDate(p.CreatedAt)}</td>
              <td className="action-cell">
                <button className="btn btn-edit" onClick={() => onEdit(p)}>
                  编辑
                </button>
                <button
                  className="btn btn-delete"
                  onClick={() => {
                    if (window.confirm(`确定要删除商品「${p.Name}」吗？`)) {
                      onDelete(p.ID);
                    }
                  }}
                >
                  删除
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
