package dto

type PutInWarehouseDTO struct {
	GoodsId    int `json:"goods_id"`
	SkuId      int `json:"sku_id"`
	ReturnNum  int `json:"return_num"`  // 用户退还数量
	StorageNum int `json:"storage_num"` // 实际入库数量
}
