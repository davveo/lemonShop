package mgr

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var globalIsRelated bool = true // 全局预加载

// IsNotFound ErrRecordNotFound
func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// prepare for other
type _BaseMgr struct {
	rdb       *gorm.DB
	wdb       *gorm.DB
	ctx       context.Context
	cancel    context.CancelFunc
	timeout   time.Duration
	isRelated bool
}

// SetTimeOut set timeout
func (obj *_BaseMgr) SetTimeOut(timeout time.Duration) {
	obj.ctx, obj.cancel = context.WithTimeout(context.Background(), timeout)
	obj.timeout = timeout
}

// SetCtx set context
func (obj *_BaseMgr) SetCtx(c context.Context) {
	if c != nil {
		obj.ctx = c
	}
}

// GetCtx get context
func (obj *_BaseMgr) GetCtx() context.Context {
	return obj.ctx
}

// Cancel cancel context
func (obj *_BaseMgr) Cancel(c context.Context) {
	obj.cancel()
}

// GetDB get gorm.DB info
func (obj *_BaseMgr) GetDB() *gorm.DB {
	return obj.rdb
}

// UpdateDB update gorm.DB info
func (obj *_BaseMgr) UpdateDB(db *gorm.DB) {
	obj.wdb = db
}

// GetIsRelated Query foreign key Association.获取是否查询外键关联(gorm.Related)
func (obj *_BaseMgr) GetIsRelated() bool {
	return obj.isRelated
}

// SetIsRelated Query foreign key Association.设置是否查询外键关联(gorm.Related)
func (obj *_BaseMgr) SetIsRelated(b bool) {
	obj.isRelated = b
}

// New new gorm.新gorm,重置条件
func (obj *_BaseMgr) New() {
	obj.wdb = obj.NewDB()
}

// NewDB new gorm.新gorm
func (obj *_BaseMgr) NewDB() *gorm.DB {
	return obj.wdb.Session(&gorm.Session{NewDB: true, Context: obj.ctx})
}

// IsNotFound ErrRecordNotFound
func (obj *_BaseMgr) IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

type options struct {
	query map[string]interface{}
}

// Option overrides behavior of Connect.
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

// OpenRelated 打开全局预加载
func OpenRelated() {
	globalIsRelated = true
}

// CloseRelated 关闭全局预加载
func CloseRelated() {
	globalIsRelated = true
}

// 自定义sql查询
type Condition struct {
	list []*conditionInfo
}

func (c *Condition) AndWithCondition(condition bool, column string, cases string, value interface{}) *Condition {
	if condition {
		c.list = append(c.list, &conditionInfo{
			andor:  "and",
			column: column, // 列名
			case_:  cases,  // 条件(and,or,in,>=,<=)
			value:  value,
		})
	}
	return c
}

// And a Condition by and .and 一个条件
func (c *Condition) And(column string, cases string, value interface{}) *Condition {
	return c.AndWithCondition(true, column, cases, value)
}

func (c *Condition) AndEq(column string, value interface{}) *Condition {
	return c.AndWithCondition(true, column, "=", value)
}

func (c *Condition) AndNotEq(column string, value interface{}) *Condition {
	return c.AndWithCondition(true, column, "!=", value)
}

func (c *Condition) AndNotLt(column string, value interface{}) *Condition {
	return c.AndWithCondition(true, column, "<", value)
}

func (c *Condition) AndNotGt(column string, value interface{}) *Condition {
	return c.AndWithCondition(true, column, ">", value)
}

func (c *Condition) AndNotLte(column string, value interface{}) *Condition {
	return c.AndWithCondition(true, column, "<=", value)
}

func (c *Condition) AndNotGte(column string, value interface{}) *Condition {
	return c.AndWithCondition(true, column, ">=", value)
}

func (c *Condition) AndLike(column string, value string) *Condition {
	return c.AndWithCondition(true, column, "like", "%"+value+"%")
}

func (c *Condition) AndIn(column string, value interface{}) *Condition {
	return c.AndWithCondition(true, column, "in", value)
}

func (c *Condition) OrWithCondition(condition bool, column string, cases string, value interface{}) *Condition {
	if condition {
		c.list = append(c.list, &conditionInfo{
			andor:  "or",
			column: column, // 列名
			case_:  cases,  // 条件(and,or,in,>=,<=)
			value:  value,
		})
	}
	return c
}

// Or a Condition by or .or 一个条件
func (c *Condition) Or(column string, cases string, value interface{}) *Condition {
	return c.OrWithCondition(true, column, cases, value)
}

func (c *Condition) Get() (where string, out []interface{}) {
	firstAnd := -1
	for i := 0; i < len(c.list); i++ { // 查找第一个and
		if c.list[i].andor == "and" {
			where = fmt.Sprintf("`%v` %v ?", c.list[i].column, c.list[i].case_)
			out = append(out, c.list[i].value)
			firstAnd = i
			break
		}
	}

	if firstAnd < 0 && len(c.list) > 0 { // 补刀
		where = fmt.Sprintf("`%v` %v ?", c.list[0].column, c.list[0].case_)
		out = append(out, c.list[0].value)
		firstAnd = 0
	}

	for i := 0; i < len(c.list); i++ { // 添加剩余的
		if firstAnd != i {
			where += fmt.Sprintf(" %v `%v` %v ?", c.list[i].andor, c.list[i].column, c.list[i].case_)
			out = append(out, c.list[i].value)
		}
	}

	where = " where" + where
	return
}

type conditionInfo struct {
	andor  string
	column string // 列名
	case_  string // 条件(in,>=,<=)
	value  interface{}
}

func test() {

	//cond := mgr.Condition{}
	//searchCondition, out := cond.
	//	AndNotEq("spec_id", 1).
	//	AndLike("spec_name", "测试").
	//	Get()
	//results, err := s.specificationMgr.Raw(searchCondition, out...)
	//if err != nil {
	//	return err
	//}
	//for _, res := range results {
	//	fmt.Println(res.SpecName)
	//}
	//s.GetCatSpecification(ctx, 1)
}
