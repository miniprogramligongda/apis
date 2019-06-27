package dao

import (
	"../conf"
	//"../util"
	"github.com/jinzhu/gorm"
	//"time"
)

// Friend structure is define and hold table friend
type Friend struct {
	ObjectOpenid  string `gorm:"type:varchar(28);not null"`
	SubjectOpenid string `gorm:"type:varchar(28);not null"`
	Status        int    `gorm:"type:tinyint(1);not null; default:0"`
	Notes         string `gorm:"type:varchar(64);"`
}

// DaoFriend struct contains a db of object of gorm.DB
type DaoFriend struct {
	db *gorm.DB
}

// NewDaoFriend function
func NewDaoFriend() (d *DaoFriend) {
	d = &DaoFriend{}
	d.db, _ = gorm.Open("mysql", conf.Conf.DbAddr)

	d.db.LogMode(true)
	return
}

// Close function
func (this *DaoFriend) Close() {
	if this.db != nil {
		this.db.Close()
	}
}

// Insert function
func (this *DaoFriend) Insert(u *Friend) error {
	if !this.db.HasTable(&Friend{}) {
		err := this.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf-8").CreateTable(&Friend{}).Error
		if err != nil {
			return err
		}
	}
	if err := this.db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

// FindByObjectAndStatus function
func (this *DaoFriend) FindByObjectAndStatus(objectOpenid string, status int) ([]*Friend, error) {
	subjectList := make([]*Friend, 0)
	err := this.db.Model(&Friend{}).Where("object_openid = ? AND status = ?", objectOpenid, status).Find(&subjectList).Error
	return subjectList, err
}

// UpdateStatusByOSbject function
func (this *DaoFriend) UpdateStatusByOSbject(objectOpenid, subjectOpenid string, stauts int) error {
	err := this.db.Model(&Friend{}).Where("object_openid = ? AND subject_openid = ?", objectOpenid, subjectOpenid).Update("status", stauts).Error
	return err
}
