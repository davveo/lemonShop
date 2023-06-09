package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type PaymentMethodMgr struct {
	*_BaseMgr
}

// NewPaymentMethodMgr open func
func NewPaymentMethodMgr(db db.Repo) *PaymentMethodMgr {
	if db == nil {
		panic(fmt.Errorf("NewPaymentMethodMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &PaymentMethodMgr{
		_BaseMgr: &_BaseMgr{
			rdb:       db.GetDbR().Table("es_payment_method"),
			wdb:       db.GetDbW().Table("es_payment_method"),
			isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1},
	}
}

// GetTableName get sql table name.获取数据库名字
func (obj *PaymentMethodMgr) GetTableName() string {
	return "es_payment_method"
}

// Reset 重置gorm会话
func (obj *PaymentMethodMgr) Reset() *PaymentMethodMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *PaymentMethodMgr) Get() (result models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *PaymentMethodMgr) Gets() (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *PaymentMethodMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMethodID method_id获取 支付方式id
func (obj *PaymentMethodMgr) WithMethodID(methodID int) Option {
	return optionFunc(func(o *options) { o.query["method_id"] = methodID })
}

// WithMethodName method_name获取 支付方式名称
func (obj *PaymentMethodMgr) WithMethodName(methodName string) Option {
	return optionFunc(func(o *options) { o.query["method_name"] = methodName })
}

// WithPluginID plugin_id获取 支付插件名称
func (obj *PaymentMethodMgr) WithPluginID(pluginID string) Option {
	return optionFunc(func(o *options) { o.query["plugin_id"] = pluginID })
}

// WithPcConfig pc_config获取 pc是否可用
func (obj *PaymentMethodMgr) WithPcConfig(pcConfig string) Option {
	return optionFunc(func(o *options) { o.query["pc_config"] = pcConfig })
}

// WithWapConfig wap_config获取 wap是否可用
func (obj *PaymentMethodMgr) WithWapConfig(wapConfig string) Option {
	return optionFunc(func(o *options) { o.query["wap_config"] = wapConfig })
}

// WithAppNativeConfig app_native_config获取 app 原生是否可用
func (obj *PaymentMethodMgr) WithAppNativeConfig(appNativeConfig string) Option {
	return optionFunc(func(o *options) { o.query["app_native_config"] = appNativeConfig })
}

// WithImage image获取 支付方式图片
func (obj *PaymentMethodMgr) WithImage(image string) Option {
	return optionFunc(func(o *options) { o.query["image"] = image })
}

// WithIsRetrace is_retrace获取 是否支持原路退回
func (obj *PaymentMethodMgr) WithIsRetrace(isRetrace int) Option {
	return optionFunc(func(o *options) { o.query["is_retrace"] = isRetrace })
}

// WithAppReactConfig app_react_config获取 app RN是否可用
func (obj *PaymentMethodMgr) WithAppReactConfig(appReactConfig string) Option {
	return optionFunc(func(o *options) { o.query["app_react_config"] = appReactConfig })
}

// WithMiniConfig mini_config获取 小程序配置
func (obj *PaymentMethodMgr) WithMiniConfig(miniConfig string) Option {
	return optionFunc(func(o *options) { o.query["mini_config"] = miniConfig })
}

// GetByOption 功能选项模式获取
func (obj *PaymentMethodMgr) GetByOption(opts ...Option) (result models.EsPaymentMethod, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *PaymentMethodMgr) GetByOptions(opts ...Option) (results []*models.EsPaymentMethod, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *PaymentMethodMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsPaymentMethod, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromMethodID 通过method_id获取内容 支付方式id
func (obj *PaymentMethodMgr) GetFromMethodID(methodID int) (result models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`method_id` = ?", methodID).First(&result).Error

	return
}

// GetBatchFromMethodID 批量查找 支付方式id
func (obj *PaymentMethodMgr) GetBatchFromMethodID(methodIDs []int) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`method_id` IN (?)", methodIDs).Find(&results).Error

	return
}

// GetFromMethodName 通过method_name获取内容 支付方式名称
func (obj *PaymentMethodMgr) GetFromMethodName(methodName string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`method_name` = ?", methodName).Find(&results).Error

	return
}

// GetBatchFromMethodName 批量查找 支付方式名称
func (obj *PaymentMethodMgr) GetBatchFromMethodName(methodNames []string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`method_name` IN (?)", methodNames).Find(&results).Error

	return
}

// GetFromPluginID 通过plugin_id获取内容 支付插件名称
func (obj *PaymentMethodMgr) GetFromPluginID(pluginID string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`plugin_id` = ?", pluginID).Find(&results).Error

	return
}

// GetBatchFromPluginID 批量查找 支付插件名称
func (obj *PaymentMethodMgr) GetBatchFromPluginID(pluginIDs []string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`plugin_id` IN (?)", pluginIDs).Find(&results).Error

	return
}

// GetFromPcConfig 通过pc_config获取内容 pc是否可用
func (obj *PaymentMethodMgr) GetFromPcConfig(pcConfig string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`pc_config` = ?", pcConfig).Find(&results).Error

	return
}

// GetBatchFromPcConfig 批量查找 pc是否可用
func (obj *PaymentMethodMgr) GetBatchFromPcConfig(pcConfigs []string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`pc_config` IN (?)", pcConfigs).Find(&results).Error

	return
}

// GetFromWapConfig 通过wap_config获取内容 wap是否可用
func (obj *PaymentMethodMgr) GetFromWapConfig(wapConfig string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`wap_config` = ?", wapConfig).Find(&results).Error

	return
}

// GetBatchFromWapConfig 批量查找 wap是否可用
func (obj *PaymentMethodMgr) GetBatchFromWapConfig(wapConfigs []string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`wap_config` IN (?)", wapConfigs).Find(&results).Error

	return
}

// GetFromAppNativeConfig 通过app_native_config获取内容 app 原生是否可用
func (obj *PaymentMethodMgr) GetFromAppNativeConfig(appNativeConfig string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`app_native_config` = ?", appNativeConfig).Find(&results).Error

	return
}

// GetBatchFromAppNativeConfig 批量查找 app 原生是否可用
func (obj *PaymentMethodMgr) GetBatchFromAppNativeConfig(appNativeConfigs []string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`app_native_config` IN (?)", appNativeConfigs).Find(&results).Error

	return
}

// GetFromImage 通过image获取内容 支付方式图片
func (obj *PaymentMethodMgr) GetFromImage(image string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`image` = ?", image).Find(&results).Error

	return
}

// GetBatchFromImage 批量查找 支付方式图片
func (obj *PaymentMethodMgr) GetBatchFromImage(images []string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`image` IN (?)", images).Find(&results).Error

	return
}

// GetFromIsRetrace 通过is_retrace获取内容 是否支持原路退回
func (obj *PaymentMethodMgr) GetFromIsRetrace(isRetrace int) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`is_retrace` = ?", isRetrace).Find(&results).Error

	return
}

// GetBatchFromIsRetrace 批量查找 是否支持原路退回
func (obj *PaymentMethodMgr) GetBatchFromIsRetrace(isRetraces []int) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`is_retrace` IN (?)", isRetraces).Find(&results).Error

	return
}

// GetFromAppReactConfig 通过app_react_config获取内容 app RN是否可用
func (obj *PaymentMethodMgr) GetFromAppReactConfig(appReactConfig string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`app_react_config` = ?", appReactConfig).Find(&results).Error

	return
}

// GetBatchFromAppReactConfig 批量查找 app RN是否可用
func (obj *PaymentMethodMgr) GetBatchFromAppReactConfig(appReactConfigs []string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`app_react_config` IN (?)", appReactConfigs).Find(&results).Error

	return
}

// GetFromMiniConfig 通过mini_config获取内容 小程序配置
func (obj *PaymentMethodMgr) GetFromMiniConfig(miniConfig string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`mini_config` = ?", miniConfig).Find(&results).Error

	return
}

// GetBatchFromMiniConfig 批量查找 小程序配置
func (obj *PaymentMethodMgr) GetBatchFromMiniConfig(miniConfigs []string) (results []*models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`mini_config` IN (?)", miniConfigs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *PaymentMethodMgr) FetchByPrimaryKey(methodID int) (result models.EsPaymentMethod, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPaymentMethod{}).Where("`method_id` = ?", methodID).First(&result).Error

	return
}
