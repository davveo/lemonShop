package consts

type ServiceStatusEnum string

const (
	APPLY_SERVICE_STATUS           ServiceStatusEnum = "APPLY"           // 待审核
	PASS_SERVICE_STATUS            ServiceStatusEnum = "PASS"            // 审核通过
	REFUSE_SERVICE_STATUS          ServiceStatusEnum = "REFUSE"          // 审核未通过
	FULL_COURIER_SERVICE_STATUS    ServiceStatusEnum = "FULL_COURIER"    // 已退还商品
	WAIT_FOR_MANUAL_SERVICE_STATUS ServiceStatusEnum = "WAIT_FOR_MANUAL" // 待人工处理
	STOCK_IN_SERVICE_STATUS        ServiceStatusEnum = "STOCK_IN"        // 已入库
	REFUNDING_SERVICE_STATUS       ServiceStatusEnum = "REFUNDING"       // 退款中
	REFUNDFAIL_SERVICE_STATUS      ServiceStatusEnum = "REFUNDFAIL"      // 退款失败
	COMPLETED_SERVICE_STATUS       ServiceStatusEnum = "COMPLETED"       // 已完成
	CLOSED_SERVICE_STATUS          ServiceStatusEnum = "CLOSED"          // 已关闭
	ERROR_EXCEPTION_SERVICE_STATUS ServiceStatusEnum = "ERROR_EXCEPTION" // 异常状态
)
