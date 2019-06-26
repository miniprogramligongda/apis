package dao

import (
	"../conf"
	"../util"
	"github.com/jinzhu/gorm"
)

// UserInfo struct
// 需要通过tag 规定primary_key，通过index规定索引
type UserInfo struct {
	Openid    string `gorm:"primary_key;type:varchar(28);not null;index:openid_idx"`
	AvatarUrl string `gorm:"type:varchar(128);not null"`
	Gender    int    `gorm:"type:tinyint(1)"`
	Nickname  string `gorm:"type:varchar(32)"`
	//City      string `gorm:"type:varchar(64)"`
	//Country   string `gorm:"type:varchar(64)"`
	//Language  string `gorm:"type:varchar(64)"`
	//Province  string `gorm:"type:varchar(32)"`
}

type DaoUserInfo struct {
	db *gorm.DB
}

// NewDaoUserInfo
func NewDaoUserInfo() (d *DaoUserInfo) {
	d = &DaoUserInfo{}
	d.db, _ = gorm.Open("mysql", conf.Conf.DbAddr)

	//db.DB().SetMaxIdleConns(c.Idle)
	//db.DB().SetMaxOpenConns(c.Active)
	//db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	d.db.LogMode(true)
	return
}

func (this *DaoUserInfo) Insert(u *UserInfo) {
	if !this.db.HasTable(&UserInfo{}) {
		err := this.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserInfo{}).Error
		util.CheckErr(err)
	}
	util.CheckErr(this.db.Create(u).Error)
}

func (this *DaoUserInfo) FindByOpenid(openid string) (*UserInfo, error) {
	result := &UserInfo{}
	err := this.db.Model(&UserInfo{}).Where("openid = ?", openid).First(result).Error
	return result, err
}

func (this *DaoUserInfo) Close() {
	if this.db != nil {
		this.db.Close()
	}
}
