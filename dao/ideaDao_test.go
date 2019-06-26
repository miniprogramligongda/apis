package dao

import (
	"fmt"
	"testing"
	"time"
)

// func TestIdeaDao
//
func TestIdeaDao(t *testing.T) {
	i := &Idea{}
	//i.Iid = 2
	i.Openid = fmt.Sprintf("%028d", 1)
	i.Time = time.Now()
	i.Content = "a test idea"
	i.Favorite = 0
	i.Like = 0

	dao := NewDaoIdea()
	fmt.Printf("Prepare to insert : %v\n", i)
	dao.Insert(i)
	result := dao.FindByIid(i.Iid)
	fmt.Printf("result : %v", result)
}

// type Idea struct {
// 	Iid      string    `gorm:"primary_key;type:int(11);not null;index:iid_idx"`
// 	Openid   string    `gorm:"type:varchar(28);not null"`
// 	Time     time.Time `gorm:"type:timestamp;not null; default now()"`
// 	Content  string    `gorm:"type:varchar(1024);not null"`
// 	Like     string    `gorm:"type:smallint(6);not null;default 0"`
// 	Favorite string    `gorm:"type:smallint(6);not null;default 0"`
// }
