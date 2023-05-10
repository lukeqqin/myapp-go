package domain

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _WxUserMgr struct {
	*_BaseMgr
}

// WxUserMgr open func
func WxUserMgr(db *gorm.DB) *_WxUserMgr {
	if db == nil {
		panic(fmt.Errorf("WxUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxUserMgr) GetTableName() string {
	return "wx_user"
}

// Reset 重置gorm会话
func (obj *_WxUserMgr) Reset() *_WxUserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WxUserMgr) Get() (result WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WxUserMgr) Gets() (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WxUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WxUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WxUserMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAppid appid获取
func (obj *_WxUserMgr) WithAppid(appid string) Option {
	return optionFunc(func(o *options) { o.query["appid"] = appid })
}

// WithOpenid openid获取
func (obj *_WxUserMgr) WithOpenid(openid string) Option {
	return optionFunc(func(o *options) { o.query["openid"] = openid })
}

// WithUnionid unionid获取
func (obj *_WxUserMgr) WithUnionid(unionid string) Option {
	return optionFunc(func(o *options) { o.query["unionid"] = unionid })
}

// WithSessionKey session_key获取
func (obj *_WxUserMgr) WithSessionKey(sessionKey string) Option {
	return optionFunc(func(o *options) { o.query["session_key"] = sessionKey })
}

// WithAccessToken access_token获取
func (obj *_WxUserMgr) WithAccessToken(accessToken string) Option {
	return optionFunc(func(o *options) { o.query["access_token"] = accessToken })
}

// WithUserID user_id获取
func (obj *_WxUserMgr) WithUserID(userID uint64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithPhone phone获取 手机号(不带区号)
func (obj *_WxUserMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithSex sex获取 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
func (obj *_WxUserMgr) WithSex(sex bool) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithCity city获取 用户所在城市
func (obj *_WxUserMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithLanguage language获取 用户的语言，简体中文为zh_CN
func (obj *_WxUserMgr) WithLanguage(language string) Option {
	return optionFunc(func(o *options) { o.query["language"] = language })
}

// WithProvince province获取 用户所在省份
func (obj *_WxUserMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithCountry country获取 用户所在国家
func (obj *_WxUserMgr) WithCountry(country string) Option {
	return optionFunc(func(o *options) { o.query["country"] = country })
}

// WithRemark remark获取 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
func (obj *_WxUserMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithNickname nickname获取 昵称
func (obj *_WxUserMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithAvatar avatar获取 头像地址
func (obj *_WxUserMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithCreatedAt created_at获取
func (obj *_WxUserMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_WxUserMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_WxUserMgr) GetByOption(opts ...Option) (result WxUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WxUserMgr) GetByOptions(opts ...Option) (results []*WxUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WxUserMgr) GetFromID(id uint64) (result WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WxUserMgr) GetBatchFromID(ids []uint64) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAppid 通过appid获取内容
func (obj *_WxUserMgr) GetFromAppid(appid string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`appid` = ?", appid).Find(&results).Error

	return
}

// GetBatchFromAppid 批量查找
func (obj *_WxUserMgr) GetBatchFromAppid(appids []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`appid` IN (?)", appids).Find(&results).Error

	return
}

// GetFromOpenid 通过openid获取内容
func (obj *_WxUserMgr) GetFromOpenid(openid string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`openid` = ?", openid).Find(&results).Error

	return
}

// GetBatchFromOpenid 批量查找
func (obj *_WxUserMgr) GetBatchFromOpenid(openids []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`openid` IN (?)", openids).Find(&results).Error

	return
}

// GetFromUnionid 通过unionid获取内容
func (obj *_WxUserMgr) GetFromUnionid(unionid string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`unionid` = ?", unionid).Find(&results).Error

	return
}

// GetBatchFromUnionid 批量查找
func (obj *_WxUserMgr) GetBatchFromUnionid(unionids []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`unionid` IN (?)", unionids).Find(&results).Error

	return
}

// GetFromSessionKey 通过session_key获取内容
func (obj *_WxUserMgr) GetFromSessionKey(sessionKey string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`session_key` = ?", sessionKey).Find(&results).Error

	return
}

// GetBatchFromSessionKey 批量查找
func (obj *_WxUserMgr) GetBatchFromSessionKey(sessionKeys []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`session_key` IN (?)", sessionKeys).Find(&results).Error

	return
}

// GetFromAccessToken 通过access_token获取内容
func (obj *_WxUserMgr) GetFromAccessToken(accessToken string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`access_token` = ?", accessToken).Find(&results).Error

	return
}

// GetBatchFromAccessToken 批量查找
func (obj *_WxUserMgr) GetBatchFromAccessToken(accessTokens []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`access_token` IN (?)", accessTokens).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_WxUserMgr) GetFromUserID(userID uint64) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_WxUserMgr) GetBatchFromUserID(userIDs []uint64) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机号(不带区号)
func (obj *_WxUserMgr) GetFromPhone(phone string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找 手机号(不带区号)
func (obj *_WxUserMgr) GetBatchFromPhone(phones []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
func (obj *_WxUserMgr) GetFromSex(sex bool) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`sex` = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量查找 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
func (obj *_WxUserMgr) GetBatchFromSex(sexs []bool) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`sex` IN (?)", sexs).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 用户所在城市
func (obj *_WxUserMgr) GetFromCity(city string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找 用户所在城市
func (obj *_WxUserMgr) GetBatchFromCity(citys []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromLanguage 通过language获取内容 用户的语言，简体中文为zh_CN
func (obj *_WxUserMgr) GetFromLanguage(language string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`language` = ?", language).Find(&results).Error

	return
}

// GetBatchFromLanguage 批量查找 用户的语言，简体中文为zh_CN
func (obj *_WxUserMgr) GetBatchFromLanguage(languages []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`language` IN (?)", languages).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 用户所在省份
func (obj *_WxUserMgr) GetFromProvince(province string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`province` = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量查找 用户所在省份
func (obj *_WxUserMgr) GetBatchFromProvince(provinces []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`province` IN (?)", provinces).Find(&results).Error

	return
}

// GetFromCountry 通过country获取内容 用户所在国家
func (obj *_WxUserMgr) GetFromCountry(country string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`country` = ?", country).Find(&results).Error

	return
}

// GetBatchFromCountry 批量查找 用户所在国家
func (obj *_WxUserMgr) GetBatchFromCountry(countrys []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`country` IN (?)", countrys).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
func (obj *_WxUserMgr) GetFromRemark(remark string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
func (obj *_WxUserMgr) GetBatchFromRemark(remarks []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 昵称
func (obj *_WxUserMgr) GetFromNickname(nickname string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`nickname` = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量查找 昵称
func (obj *_WxUserMgr) GetBatchFromNickname(nicknames []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像地址
func (obj *_WxUserMgr) GetFromAvatar(avatar string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找 头像地址
func (obj *_WxUserMgr) GetBatchFromAvatar(avatars []string) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_WxUserMgr) GetFromCreatedAt(createdAt time.Time) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_WxUserMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_WxUserMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_WxUserMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WxUserMgr) FetchByPrimaryKey(id uint64) (result WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByOpenid primary or index 获取唯一内容
func (obj *_WxUserMgr) FetchUniqueIndexByOpenid(appid string, openid string) (result WxUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxUser{}).Where("`appid` = ? AND `openid` = ?", appid, openid).Find(&result).Error

	return
}
