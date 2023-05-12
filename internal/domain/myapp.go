package domain

import (
	"time"
)

// WxGenealogy [...]
type WxGenealogy struct {
	ID       int64     `gorm:"primaryKey;column:id;type:bigint(20);not null"`
	Title    string    `gorm:"column:title;type:varchar(32);not null"`
	Cover    string    `gorm:"column:cover;type:varchar(255)"`
	CreateAt time.Time `gorm:"column:create_at;type:datetime"`
	UpdateAt time.Time `gorm:"column:update_at;type:datetime"`
	CreateBy int64     `gorm:"column:create_by;type:bigint(20)"`
	UpdateBy int64     `gorm:"column:update_by;type:bigint(20)"`
}

// WxGenealogyColumns get sql column name.获取数据库列名
var WxGenealogyColumns = struct {
	ID       string
	Title    string
	Cover    string
	CreateAt string
	UpdateAt string
	CreateBy string
	UpdateBy string
}{
	ID:       "id",
	Title:    "title",
	Cover:    "cover",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	CreateBy: "create_by",
	UpdateBy: "update_by",
}

// WxGenealogyTags [...]
type WxGenealogyTags struct {
	ID          int64  `gorm:"primaryKey;column:id;type:bigint(20);not null"`
	GenealogyID int64  `gorm:"index:wx_genealogy_tags_genealogy_id_index;column:genealogy_id;type:bigint(20);not null"`
	Name        string `gorm:"column:name;type:varchar(32);not null"`
}

// WxGenealogyTagsColumns get sql column name.获取数据库列名
var WxGenealogyTagsColumns = struct {
	ID          string
	GenealogyID string
	Name        string
}{
	ID:          "id",
	GenealogyID: "genealogy_id",
	Name:        "name",
}

// WxUser 小程序-微信openid用户表
type WxUser struct {
	ID          int64     `gorm:"primaryKey;column:id;type:bigint(20);not null"`
	Appid       string    `gorm:"uniqueIndex:openid;column:appid;type:varchar(255);not null;default:''"`
	Openid      string    `gorm:"uniqueIndex:openid;column:openid;type:varchar(255);not null;default:''"`
	Unionid     string    `gorm:"column:unionid;type:varchar(255);not null;default:''"`
	SessionKey  string    `gorm:"column:session_key;type:varchar(255);not null;default:''"`
	AccessToken string    `gorm:"column:access_token;type:varchar(255);not null;default:''"`
	UserID      uint64    `gorm:"column:user_id;type:bigint(20) unsigned"`
	Phone       string    `gorm:"column:phone;type:varchar(255);not null"`               // 手机号(不带区号)
	Sex         bool      `gorm:"column:sex;type:tinyint(1);default:0"`                  // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City        string    `gorm:"column:city;type:varchar(64)"`                          // 用户所在城市
	Language    string    `gorm:"column:language;type:varchar(64)"`                      // 用户的语言，简体中文为zh_CN
	Province    string    `gorm:"column:province;type:varchar(64)"`                      // 用户所在省份
	Country     string    `gorm:"column:country;type:varchar(64)"`                       // 用户所在国家
	Remark      string    `gorm:"column:remark;type:varchar(256)"`                       // 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	Nickname    string    `gorm:"column:nickname;type:varchar(255);not null;default:''"` // 昵称
	Avatar      string    `gorm:"column:avatar;type:varchar(255);not null;default:''"`   // 头像地址
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;not null"`
}

// WxUserColumns get sql column name.获取数据库列名
var WxUserColumns = struct {
	ID          string
	Appid       string
	Openid      string
	Unionid     string
	SessionKey  string
	AccessToken string
	UserID      string
	Phone       string
	Sex         string
	City        string
	Language    string
	Province    string
	Country     string
	Remark      string
	Nickname    string
	Avatar      string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Appid:       "appid",
	Openid:      "openid",
	Unionid:     "unionid",
	SessionKey:  "session_key",
	AccessToken: "access_token",
	UserID:      "user_id",
	Phone:       "phone",
	Sex:         "sex",
	City:        "city",
	Language:    "language",
	Province:    "province",
	Country:     "country",
	Remark:      "remark",
	Nickname:    "nickname",
	Avatar:      "avatar",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}
