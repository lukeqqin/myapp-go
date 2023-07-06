package domain

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _WxCosMgr struct {
	*_BaseMgr
}

// WxCosMgr open func
func WxCosMgr(db *gorm.DB) *_WxCosMgr {
	if db == nil {
		panic(fmt.Errorf("WxCosMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxCosMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_cos"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxCosMgr) GetTableName() string {
	return "wx_cos"
}

// Reset 重置gorm会话
func (obj *_WxCosMgr) Reset() *_WxCosMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WxCosMgr) Get() (result WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WxCosMgr) Gets() (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WxCosMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WxCos{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WxCosMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithURL url获取
func (obj *_WxCosMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithCreateBy create_by获取
func (obj *_WxCosMgr) WithCreateBy(createBy int64) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithStatus status获取
func (obj *_WxCosMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithUpdateAt update_at获取
func (obj *_WxCosMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithUpdateTy update_ty获取
func (obj *_WxCosMgr) WithUpdateTy(updateTy int64) Option {
	return optionFunc(func(o *options) { o.query["update_ty"] = updateTy })
}

// WithCreateAt create_at获取
func (obj *_WxCosMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_WxCosMgr) GetByOption(opts ...Option) (result WxCos, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WxCosMgr) GetByOptions(opts ...Option) (results []*WxCos, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WxCosMgr) GetFromID(id int64) (result WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WxCosMgr) GetBatchFromID(ids []int64) (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容
func (obj *_WxCosMgr) GetFromURL(url string) (result WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`url` = ?", url).Find(&result).Error

	return
}

// GetBatchFromURL 批量查找
func (obj *_WxCosMgr) GetBatchFromURL(urls []string) (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`url` IN (?)", urls).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_WxCosMgr) GetFromCreateBy(createBy int64) (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_WxCosMgr) GetBatchFromCreateBy(createBys []int64) (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_WxCosMgr) GetFromStatus(status int8) (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_WxCosMgr) GetBatchFromStatus(statuss []int8) (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_WxCosMgr) GetFromUpdateAt(updateAt time.Time) (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找
func (obj *_WxCosMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromUpdateTy 通过update_ty获取内容
func (obj *_WxCosMgr) GetFromUpdateTy(updateTy int64) (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`update_ty` = ?", updateTy).Find(&results).Error

	return
}

// GetBatchFromUpdateTy 批量查找
func (obj *_WxCosMgr) GetBatchFromUpdateTy(updateTys []int64) (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`update_ty` IN (?)", updateTys).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_WxCosMgr) GetFromCreateAt(createAt time.Time) (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找
func (obj *_WxCosMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WxCosMgr) FetchByPrimaryKey(id int64) (result WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByWxCosURL primary or index 获取唯一内容
func (obj *_WxCosMgr) FetchUniqueByWxCosURL(url string) (result WxCos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxCos{}).Where("`url` = ?", url).Find(&result).Error

	return
}
