package domain

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _WxGenealogyMgr struct {
	*_BaseMgr
}

// WxGenealogyMgr open func
func WxGenealogyMgr(db *gorm.DB) *_WxGenealogyMgr {
	if db == nil {
		panic(fmt.Errorf("WxGenealogyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxGenealogyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_genealogy"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxGenealogyMgr) GetTableName() string {
	return "wx_genealogy"
}

// Reset 重置gorm会话
func (obj *_WxGenealogyMgr) Reset() *_WxGenealogyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WxGenealogyMgr) Get() (result WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WxGenealogyMgr) Gets() (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WxGenealogyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WxGenealogyMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取
func (obj *_WxGenealogyMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithCover cover获取
func (obj *_WxGenealogyMgr) WithCover(cover string) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

// WithCreateAt create_at获取
func (obj *_WxGenealogyMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_WxGenealogyMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithCreateBy create_by获取
func (obj *_WxGenealogyMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_WxGenealogyMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_WxGenealogyMgr) GetByOption(opts ...Option) (result WxGenealogy, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WxGenealogyMgr) GetByOptions(opts ...Option) (results []*WxGenealogy, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WxGenealogyMgr) GetFromID(id int64) (result WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WxGenealogyMgr) GetBatchFromID(ids []int64) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_WxGenealogyMgr) GetFromTitle(title string) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找
func (obj *_WxGenealogyMgr) GetBatchFromTitle(titles []string) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromCover 通过cover获取内容
func (obj *_WxGenealogyMgr) GetFromCover(cover string) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`cover` = ?", cover).Find(&results).Error

	return
}

// GetBatchFromCover 批量查找
func (obj *_WxGenealogyMgr) GetBatchFromCover(covers []string) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`cover` IN (?)", covers).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_WxGenealogyMgr) GetFromCreateAt(createAt time.Time) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找
func (obj *_WxGenealogyMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_WxGenealogyMgr) GetFromUpdateAt(updateAt time.Time) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找
func (obj *_WxGenealogyMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_WxGenealogyMgr) GetFromCreateBy(createBy string) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_WxGenealogyMgr) GetBatchFromCreateBy(createBys []string) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_WxGenealogyMgr) GetFromUpdateBy(updateBy string) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_WxGenealogyMgr) GetBatchFromUpdateBy(updateBys []string) (results []*WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WxGenealogyMgr) FetchByPrimaryKey(id int64) (result WxGenealogy, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogy{}).Where("`id` = ?", id).Find(&result).Error

	return
}
