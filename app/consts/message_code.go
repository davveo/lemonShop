package consts

type MessageCodeEnum string

const (
	SHOPORDERSNEW             = "SHOPORDERSNEW"             //店铺新订单创建提醒
	SHOPORDERSPAY             = "SHOPORDERSPAY"             //店铺订单支付提醒
	SHOPORDERSRECEIVE         = "SHOPORDERSRECEIVE"         //店铺订单收货提醒
	SHOPORDERSEVALUATE        = "SHOPORDERSEVALUATE"        //店铺订单评价提醒
	SHOPORDERSCANCEL          = "SHOPORDERSCANCEL"          //店铺订单取消提醒
	SHOPREFUND                = "SHOPREFUND"                //店铺退款提醒
	SHOPRETURN                = "SHOPRETURN"                //店铺退货提醒
	SHOPGOODSVIOLATION        = "SHOPGOODSVIOLATION"        //商品违规被禁售提醒（商品下架）
	SHOPGOODSVERIFY           = "SHOPGOODSVERIFY"           //商品审核失败提醒
	MEMBERRETURNUPDATE        = "MEMBERRETURNUPDATE"        //退货提醒
	MEMBERREFUNDUPDATE        = "MEMBERREFUNDUPDATE"        //售后服务退款提醒
	MEMBERORDERSSEND          = "MEMBERORDERSSEND"          //订单发货提醒
	MEMBERORDERSRECEIVE       = "MEMBERORDERSRECEIVE"       //订单收货提醒
	MEMBERORDERSPAY           = "MEMBERORDERSPAY"           //订单支付提醒
	MEMBERORDERSCANCEL        = "MEMBERORDERSCANCEL"        //订单取消提醒
	MOBILECODESEND            = "MOBILECODESEND"            //手机发送验证码
	EMAILCODESEND             = "EMAILCODESEND"             //邮箱发送验证码
	SHOPGOODSMARKETENABLE     = "SHOPGOODSMARKETENABLE"     //商品下架消息提醒
	MEMBERLOGINSUCCESS        = "MEMBERLOGINSUCCESS"        //会员登录成功提醒
	MEMBERREGISTESUCCESS      = "MEMBERREGISTESUCCESS"      //会员注册成功提醒
	MEMBERORDERCREATE         = "MEMBERORDERCREATE"         //订单创建提醒
	SHOPGOODSASK              = "SHOPGOODSASK"              //店铺商品咨询提醒
	SHOPAFTERSALE             = "SHOPAFTERSALE"             //订单申请售后服务提醒
	SHOPAFTERSALEGOODSSHIP    = "SHOPAFTERSALEGOODSSHIP"    //买家退还售后商品提醒
	MEMBERASAUDIT             = "MEMBERASAUDIT"             //售后服务申请审核提醒
	MEMBERASCOMPLETED         = "MEMBERASCOMPLETED"         //售后服务完成提醒
	MEMBERASAUDITRETURNCHANGE = "MEMBERASAUDITRETURNCHANGE" //退货和换货售后服务审核通过提醒
	MEMBERASCLOSED            = "MEMBERASCLOSED"            //售后服务单关闭提醒
	PINTUANORDERAUTOCANCEL    = "PINTUANORDERAUTOCANCEL"    //拼团订单自动取消提醒");
)
