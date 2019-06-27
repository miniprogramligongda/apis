package dao

import (
	"../conf"
	//"../util"
	"fmt"
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
		fmt.Println("no table friend")
		err := this.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Friend{}).Error
		if err != nil {
			return err
		}
	} else {
		fmt.Println("table friend exist")
	}

	if err := this.db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

// HaveFriendByOSbject function
func (this *DaoFriend) HaveFriendByOSbject(objectOpenid, subjectOpenid string) (bool, error) {
	f := &Friend{}
	ff := &Friend{}
	err := this.db.Model(&Friend{}).Where("object_openid = ? AND subject_openid = ?", objectOpenid, subjectOpenid).First(f).Error

	// fmt.Println("f: ", f)
	// fmt.Println("ff: ", ff)
	if fmt.Sprint(f) == fmt.Sprint(ff) {
		fmt.Printf("have no record of O=%s and S=%s\n", objectOpenid, subjectOpenid)
		return false, err
	}

	return true, err
}

// FindByObjectAndStatus function
func (this *DaoFriend) FindByObjectAndStatus(objectOpenid string, status int) ([]*Friend, error) {
	subjectList := make([]*Friend, 0)
	err := this.db.Model(&Friend{}).Where("object_openid = ? AND status = ?", objectOpenid, status).Find(&subjectList).Error
	return subjectList, err
}

// FindBySubjectAndStatus function
func (this *DaoFriend) FindBySubjectAndStatus(subjectOpenid string, status int) ([]*Friend, error) {
	objectList := make([]*Friend, 0)
	err := this.db.Model(&Friend{}).Where("subject_openid = ? AND status = ?", subjectOpenid, status).Find(&objectList).Error
	return objectList, err
}

// UpdateStatusByOSbject function
func (this *DaoFriend) UpdateStatusByOSbject(objectOpenid, subjectOpenid string, stauts int) error {
	err := this.db.Model(&Friend{}).Where("object_openid = ? AND subject_openid = ?", objectOpenid, subjectOpenid).Update("status", stauts).Error
	return err
}
