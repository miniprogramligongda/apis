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
