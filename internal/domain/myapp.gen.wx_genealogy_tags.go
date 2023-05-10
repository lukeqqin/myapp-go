package domain

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WxGenealogyTagsMgr struct {
	*_BaseMgr
}

// WxGenealogyTagsMgr open func
func WxGenealogyTagsMgr(db *gorm.DB) *_WxGenealogyTagsMgr {
	if db == nil {
		panic(fmt.Errorf("WxGenealogyTagsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxGenealogyTagsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_genealogy_tags"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxGenealogyTagsMgr) GetTableName() string {
	return "wx_genealogy_tags"
}

// Reset 重置gorm会话
func (obj *_WxGenealogyTagsMgr) Reset() *_WxGenealogyTagsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WxGenealogyTagsMgr) Get() (result WxGenealogyTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyTags{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WxGenealogyTagsMgr) Gets() (results []*WxGenealogyTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyTags{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WxGenealogyTagsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WxGenealogyTags{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WxGenealogyTagsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGenealogyID genealogy_id获取
func (obj *_WxGenealogyTagsMgr) WithGenealogyID(genealogyID int64) Option {
	return optionFunc(func(o *options) { o.query["genealogy_id"] = genealogyID })
}

// WithName name获取
func (obj *_WxGenealogyTagsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_WxGenealogyTagsMgr) GetByOption(opts ...Option) (result WxGenealogyTags, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyTags{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WxGenealogyTagsMgr) GetByOptions(opts ...Option) (results []*WxGenealogyTags, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyTags{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WxGenealogyTagsMgr) GetFromID(id int64) (result WxGenealogyTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyTags{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WxGenealogyTagsMgr) GetBatchFromID(ids []int64) (results []*WxGenealogyTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyTags{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGenealogyID 通过genealogy_id获取内容
func (obj *_WxGenealogyTagsMgr) GetFromGenealogyID(genealogyID int64) (results []*WxGenealogyTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyTags{}).Where("`genealogy_id` = ?", genealogyID).Find(&results).Error

	return
}

// GetBatchFromGenealogyID 批量查找
func (obj *_WxGenealogyTagsMgr) GetBatchFromGenealogyID(genealogyIDs []int64) (results []*WxGenealogyTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyTags{}).Where("`genealogy_id` IN (?)", genealogyIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_WxGenealogyTagsMgr) GetFromName(name string) (results []*WxGenealogyTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyTags{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_WxGenealogyTagsMgr) GetBatchFromName(names []string) (results []*WxGenealogyTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyTags{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WxGenealogyTagsMgr) FetchByPrimaryKey(id int64) (result WxGenealogyTags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyTags{}).Where("`id` = ?", id).Find(&result).Error

	return
}
