package domain

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _WxGenealogyMembersMgr struct {
	*_BaseMgr
}

// WxGenealogyMembersMgr open func
func WxGenealogyMembersMgr(db *gorm.DB) *_WxGenealogyMembersMgr {
	if db == nil {
		panic(fmt.Errorf("WxGenealogyMembersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxGenealogyMembersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_genealogy_members"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxGenealogyMembersMgr) GetTableName() string {
	return "wx_genealogy_members"
}

// Reset 重置gorm会话
func (obj *_WxGenealogyMembersMgr) Reset() *_WxGenealogyMembersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WxGenealogyMembersMgr) Get() (result WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WxGenealogyMembersMgr) Gets() (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WxGenealogyMembersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WxGenealogyMembersMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_WxGenealogyMembersMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithUserID user_id获取
func (obj *_WxGenealogyMembersMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithAvatar avatar获取
func (obj *_WxGenealogyMembersMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithGenealogyID genealogy_id获取
func (obj *_WxGenealogyMembersMgr) WithGenealogyID(genealogyID int64) Option {
	return optionFunc(func(o *options) { o.query["genealogy_id"] = genealogyID })
}

// WithSex sex获取
func (obj *_WxGenealogyMembersMgr) WithSex(sex int8) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithAddressID address_id获取
func (obj *_WxGenealogyMembersMgr) WithAddressID(addressID int64) Option {
	return optionFunc(func(o *options) { o.query["address_id"] = addressID })
}

// WithPhone phone获取
func (obj *_WxGenealogyMembersMgr) WithPhone(phone int) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithParentID parent_id获取
func (obj *_WxGenealogyMembersMgr) WithParentID(parentID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithBirthday birthday获取
func (obj *_WxGenealogyMembersMgr) WithBirthday(birthday time.Time) Option {
	return optionFunc(func(o *options) { o.query["birthday"] = birthday })
}

// WithAddressDetail address_detail获取
func (obj *_WxGenealogyMembersMgr) WithAddressDetail(addressDetail string) Option {
	return optionFunc(func(o *options) { o.query["address_detail"] = addressDetail })
}

// WithAlive alive获取
func (obj *_WxGenealogyMembersMgr) WithAlive(alive int8) Option {
	return optionFunc(func(o *options) { o.query["alive"] = alive })
}

// WithAge age获取
func (obj *_WxGenealogyMembersMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// WithDeadTime dead_time获取
func (obj *_WxGenealogyMembersMgr) WithDeadTime(deadTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["dead_time"] = deadTime })
}

// GetByOption 功能选项模式获取
func (obj *_WxGenealogyMembersMgr) GetByOption(opts ...Option) (result WxGenealogyMembers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WxGenealogyMembersMgr) GetByOptions(opts ...Option) (results []*WxGenealogyMembers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WxGenealogyMembersMgr) GetFromID(id int64) (result WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromID(ids []int64) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_WxGenealogyMembersMgr) GetFromName(name string) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromName(names []string) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_WxGenealogyMembersMgr) GetFromUserID(userID int64) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromUserID(userIDs []int64) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容
func (obj *_WxGenealogyMembersMgr) GetFromAvatar(avatar string) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromAvatar(avatars []string) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

// GetFromGenealogyID 通过genealogy_id获取内容
func (obj *_WxGenealogyMembersMgr) GetFromGenealogyID(genealogyID int64) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`genealogy_id` = ?", genealogyID).Find(&results).Error

	return
}

// GetBatchFromGenealogyID 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromGenealogyID(genealogyIDs []int64) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`genealogy_id` IN (?)", genealogyIDs).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容
func (obj *_WxGenealogyMembersMgr) GetFromSex(sex int8) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`sex` = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromSex(sexs []int8) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`sex` IN (?)", sexs).Find(&results).Error

	return
}

// GetFromAddressID 通过address_id获取内容
func (obj *_WxGenealogyMembersMgr) GetFromAddressID(addressID int64) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`address_id` = ?", addressID).Find(&results).Error

	return
}

// GetBatchFromAddressID 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromAddressID(addressIDs []int64) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`address_id` IN (?)", addressIDs).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容
func (obj *_WxGenealogyMembersMgr) GetFromPhone(phone int) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromPhone(phones []int) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容
func (obj *_WxGenealogyMembersMgr) GetFromParentID(parentID int64) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromParentID(parentIDs []int64) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromBirthday 通过birthday获取内容
func (obj *_WxGenealogyMembersMgr) GetFromBirthday(birthday time.Time) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`birthday` = ?", birthday).Find(&results).Error

	return
}

// GetBatchFromBirthday 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromBirthday(birthdays []time.Time) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`birthday` IN (?)", birthdays).Find(&results).Error

	return
}

// GetFromAddressDetail 通过address_detail获取内容
func (obj *_WxGenealogyMembersMgr) GetFromAddressDetail(addressDetail string) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`address_detail` = ?", addressDetail).Find(&results).Error

	return
}

// GetBatchFromAddressDetail 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromAddressDetail(addressDetails []string) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`address_detail` IN (?)", addressDetails).Find(&results).Error

	return
}

// GetFromAlive 通过alive获取内容
func (obj *_WxGenealogyMembersMgr) GetFromAlive(alive int8) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`alive` = ?", alive).Find(&results).Error

	return
}

// GetBatchFromAlive 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromAlive(alives []int8) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`alive` IN (?)", alives).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容
func (obj *_WxGenealogyMembersMgr) GetFromAge(age int) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`age` = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromAge(ages []int) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`age` IN (?)", ages).Find(&results).Error

	return
}

// GetFromDeadTime 通过dead_time获取内容
func (obj *_WxGenealogyMembersMgr) GetFromDeadTime(deadTime time.Time) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`dead_time` = ?", deadTime).Find(&results).Error

	return
}

// GetBatchFromDeadTime 批量查找
func (obj *_WxGenealogyMembersMgr) GetBatchFromDeadTime(deadTimes []time.Time) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`dead_time` IN (?)", deadTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WxGenealogyMembersMgr) FetchByPrimaryKey(id int64) (result WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchIndexByWxGenealogyMembersPk  获取多个内容
func (obj *_WxGenealogyMembersMgr) FetchIndexByWxGenealogyMembersPk(genealogyID int64) (results []*WxGenealogyMembers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMembers{}).Where("`genealogy_id` = ?", genealogyID).Find(&results).Error

	return
}
