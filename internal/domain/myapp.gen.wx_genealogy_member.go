package domain

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _WxGenealogyMemberMgr struct {
	*_BaseMgr
}

// WxGenealogyMemberMgr open func
func WxGenealogyMemberMgr(db *gorm.DB) *_WxGenealogyMemberMgr {
	if db == nil {
		panic(fmt.Errorf("WxGenealogyMemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxGenealogyMemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_genealogy_member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxGenealogyMemberMgr) GetTableName() string {
	return "wx_genealogy_member"
}

// Reset 重置gorm会话
func (obj *_WxGenealogyMemberMgr) Reset() *_WxGenealogyMemberMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WxGenealogyMemberMgr) Get() (result WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WxGenealogyMemberMgr) Gets() (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WxGenealogyMemberMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WxGenealogyMemberMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_WxGenealogyMemberMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithUserID user_id获取
func (obj *_WxGenealogyMemberMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithAvatar avatar获取
func (obj *_WxGenealogyMemberMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithGenealogyID genealogy_id获取
func (obj *_WxGenealogyMemberMgr) WithGenealogyID(genealogyID int64) Option {
	return optionFunc(func(o *options) { o.query["genealogy_id"] = genealogyID })
}

// WithSex sex获取
func (obj *_WxGenealogyMemberMgr) WithSex(sex int8) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithAddressID address_id获取
func (obj *_WxGenealogyMemberMgr) WithAddressID(addressID int64) Option {
	return optionFunc(func(o *options) { o.query["address_id"] = addressID })
}

// WithPhone phone获取
func (obj *_WxGenealogyMemberMgr) WithPhone(phone int) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithParentID parent_id获取
func (obj *_WxGenealogyMemberMgr) WithParentID(parentID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithBirthday birthday获取
func (obj *_WxGenealogyMemberMgr) WithBirthday(birthday time.Time) Option {
	return optionFunc(func(o *options) { o.query["birthday"] = birthday })
}

// WithAddressDetail address_detail获取
func (obj *_WxGenealogyMemberMgr) WithAddressDetail(addressDetail string) Option {
	return optionFunc(func(o *options) { o.query["address_detail"] = addressDetail })
}

// WithAlive alive获取
func (obj *_WxGenealogyMemberMgr) WithAlive(alive int8) Option {
	return optionFunc(func(o *options) { o.query["alive"] = alive })
}

// WithAge age获取
func (obj *_WxGenealogyMemberMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// WithDeadTime dead_time获取
func (obj *_WxGenealogyMemberMgr) WithDeadTime(deadTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["dead_time"] = deadTime })
}

// WithRole role获取
func (obj *_WxGenealogyMemberMgr) WithRole(role int8) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// GetByOption 功能选项模式获取
func (obj *_WxGenealogyMemberMgr) GetByOption(opts ...Option) (result WxGenealogyMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WxGenealogyMemberMgr) GetByOptions(opts ...Option) (results []*WxGenealogyMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WxGenealogyMemberMgr) GetFromID(id int64) (result WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromID(ids []int64) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_WxGenealogyMemberMgr) GetFromName(name string) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromName(names []string) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_WxGenealogyMemberMgr) GetFromUserID(userID int64) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromUserID(userIDs []int64) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容
func (obj *_WxGenealogyMemberMgr) GetFromAvatar(avatar string) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromAvatar(avatars []string) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

// GetFromGenealogyID 通过genealogy_id获取内容
func (obj *_WxGenealogyMemberMgr) GetFromGenealogyID(genealogyID int64) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`genealogy_id` = ?", genealogyID).Find(&results).Error

	return
}

// GetBatchFromGenealogyID 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromGenealogyID(genealogyIDs []int64) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`genealogy_id` IN (?)", genealogyIDs).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容
func (obj *_WxGenealogyMemberMgr) GetFromSex(sex int8) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`sex` = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromSex(sexs []int8) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`sex` IN (?)", sexs).Find(&results).Error

	return
}

// GetFromAddressID 通过address_id获取内容
func (obj *_WxGenealogyMemberMgr) GetFromAddressID(addressID int64) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`address_id` = ?", addressID).Find(&results).Error

	return
}

// GetBatchFromAddressID 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromAddressID(addressIDs []int64) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`address_id` IN (?)", addressIDs).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容
func (obj *_WxGenealogyMemberMgr) GetFromPhone(phone int) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromPhone(phones []int) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容
func (obj *_WxGenealogyMemberMgr) GetFromParentID(parentID int64) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromParentID(parentIDs []int64) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromBirthday 通过birthday获取内容
func (obj *_WxGenealogyMemberMgr) GetFromBirthday(birthday time.Time) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`birthday` = ?", birthday).Find(&results).Error

	return
}

// GetBatchFromBirthday 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromBirthday(birthdays []time.Time) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`birthday` IN (?)", birthdays).Find(&results).Error

	return
}

// GetFromAddressDetail 通过address_detail获取内容
func (obj *_WxGenealogyMemberMgr) GetFromAddressDetail(addressDetail string) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`address_detail` = ?", addressDetail).Find(&results).Error

	return
}

// GetBatchFromAddressDetail 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromAddressDetail(addressDetails []string) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`address_detail` IN (?)", addressDetails).Find(&results).Error

	return
}

// GetFromAlive 通过alive获取内容
func (obj *_WxGenealogyMemberMgr) GetFromAlive(alive int8) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`alive` = ?", alive).Find(&results).Error

	return
}

// GetBatchFromAlive 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromAlive(alives []int8) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`alive` IN (?)", alives).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容
func (obj *_WxGenealogyMemberMgr) GetFromAge(age int) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`age` = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromAge(ages []int) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`age` IN (?)", ages).Find(&results).Error

	return
}

// GetFromDeadTime 通过dead_time获取内容
func (obj *_WxGenealogyMemberMgr) GetFromDeadTime(deadTime time.Time) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`dead_time` = ?", deadTime).Find(&results).Error

	return
}

// GetBatchFromDeadTime 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromDeadTime(deadTimes []time.Time) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`dead_time` IN (?)", deadTimes).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_WxGenealogyMemberMgr) GetFromRole(role int8) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_WxGenealogyMemberMgr) GetBatchFromRole(roles []int8) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WxGenealogyMemberMgr) FetchByPrimaryKey(id int64) (result WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchIndexByWxGenealogyMembersPk  获取多个内容
func (obj *_WxGenealogyMemberMgr) FetchIndexByWxGenealogyMembersPk(genealogyID int64) (results []*WxGenealogyMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxGenealogyMember{}).Where("`genealogy_id` = ?", genealogyID).Find(&results).Error

	return
}
