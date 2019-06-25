package dao

import (
	"../conf"
	"../util"
	"github.com/jinzhu/gorm"
	"time"
)

// Idea struct is define and hold table idea
type Idea struct {
	Iid      string    `gorm:"primary_key;type:int(11);not null;index:iid_idx"`
	Openid   string    `gorm:"type:varchar(28);not null"`
	Time     time.Time `gorm:"type:timestamp;not null; default now()"`
	Content  string    `gorm:"type:varchar(1024);not null"`
	Like     string    `gorm:"type:smallint(6);not null;default 0"`
	Favorite string    `gorm:"type:smallint(6);not null;default 0"`
}

// DaoIdea struct contains a db object of grom.DB
type DaoIdea struct {
	db *gorm.DB
}

// NewDaoIdea create a db connect to table Idea
func NewDaoIdea() (d *DaoIdea) {
	d = &DaoIdea{}
	d.db, _ = gorm.Open("mysql", conf.Conf.DbAddr)

	d.db.LogMode(true)
	return
}

// Insert insert data into table Idea
func (this *DaoIdea) Insert(u *Idea) {
	if !this.db.HasTable(&Idea{}) {
		err := this.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Idea{}).Error
		util.CheckErr(err)
	}
	util.CheckErr(this.db.Create(u).Error)
}

func (this *DaoIdea) FindByIid(iid string) *Idea {
	result := &Idea{}
	err := this.db.Model(&Idea{}).Where("iid=" + iid).First(result).Error
	util.CheckErr(err)
	return result
}

func (this *DaoIdea) Close() {
	if this.db != nil {
		this.db.Close()
	}
}
