package dao

import (
	"fmt"
	"testing"
)

func TestFriendDao(t *testing.T) {
	f := &Friend{}
	f.ObjectOpenid = "oYKEK43t1HNAljVQRGGMibzjRIeQ"
	f.SubjectOpenid = "oYKEK4xa2MoGBu1CITy7WTJBFqT8"
	f.Status = 0
	f.Notes = "I really want to friend with you!"

	d := NewDaoFriend()

	//TEST PASS
	fmt.Println("insert: ", f)
	errI := d.Insert(f)
	if errI != nil {
		fmt.Println(errI)
	}

	fmt.Printf("find by O:%s & St:%d\n", f.ObjectOpenid, f.Status)
	fList, errF := d.FindByObjectAndStatus(f.ObjectOpenid, f.Status)
	if errF != nil {
		fmt.Println(errF)
	} else {
		fmt.Println("result:")
		for _, item := range fList {
			fmt.Printf("O:%s\nS:%s\nST:%d\nN:%s\n", item.ObjectOpenid, item.SubjectOpenid, item.Status, item.Notes)
		}
	}

	fmt.Printf("find by S:%s & St:%d\n", f.SubjectOpenid, f.Status)
	fList, errF = d.FindBySubjectAndStatus(f.SubjectOpenid, f.Status)
	if errF != nil {
		fmt.Println(errF)
	} else {
		fmt.Println("result:")
		fmt.Println(fList)
		for _, item := range fList {
			fmt.Printf("O:%s\nS:%s\nST:%d\nN:%s\n", item.ObjectOpenid, item.SubjectOpenid, item.Status, item.Notes)
		}
	}

	var s int
	if f.Status == 1 {
		s = 0
	} else {
		s = 1
	}
	fmt.Printf("update status by O:%s & S%s\n", f.ObjectOpenid, f.SubjectOpenid)
	errU := d.UpdateStatusByOSbject(f.ObjectOpenid, f.SubjectOpenid, s)
	if errU != nil {
		fmt.Println(errU)
	} else {
		fmt.Println("update success.")
		fmt.Printf("find by O:%s & St:%d\n", f.ObjectOpenid, s)
		fList, errFF := d.FindByObjectAndStatus(f.ObjectOpenid, s)
		if errFF != nil {
			fmt.Println(errFF)
		} else {
			fmt.Println("result:")
			for _, item := range fList {
				fmt.Printf("O:%s\nS:%s\nST:%d\nN:%s\n", item.ObjectOpenid, item.SubjectOpenid, item.Status, item.Notes)
			}
		}
	}
}
