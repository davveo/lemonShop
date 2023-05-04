package entity

type Category struct {
	CategoryID    int    `json:"category_id"`    // 主键
	Name          string `json:"name"`           // 分类名称
	ParentID      int    `json:"parent_id"`      // 分类父id
	CategoryPath  string `json:"category_path"`  // 分类父子路径
	GoodsCount    int    `json:"goods_count"`    // 该分类下商品数量
	CategoryOrder int    `json:"category_order"` // 分类排序
	Image         string `json:"image"`          // 分类图标
}
