package consts

type ServiceOperateTypeEnum string

const (
	SELLER_AUDIT        ServiceOperateTypeEnum = "SELLER_AUDIT"        // 商家审核操作
	STOCK_IN            ServiceOperateTypeEnum = "STOCK_IN"            // 审核入库操作
	SELLER_REFUND       ServiceOperateTypeEnum = "SELLER_REFUND"       // 商家退款操作
	ADMIN_REFUND        ServiceOperateTypeEnum = "ADMIN_REFUND"        // 平台退款操作
	FILL_LOGISTICS_INFO ServiceOperateTypeEnum = "FILL_LOGISTICS_INFO" // 买家填写物流信息操作
	CLOSE               ServiceOperateTypeEnum = "CLOSE"               // 关闭操作
	CREATE_NEW_ORDER    ServiceOperateTypeEnum = "CREATE_NEW_ORDER"    // 创建新订单操作
)
