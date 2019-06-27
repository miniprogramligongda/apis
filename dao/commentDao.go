package dao

import (
	"encoding/json"
	"fmt"
	"time"

	"../conf"
	"github.com/garyburd/redigo/redis"
)

type Comment struct {
	Iid     int64
	Openid  string
	Content string
	Time    time.Time
}

// DaoIdea struct contains a db object of grom.DB
type DaoComment struct {
	redis *redis.Pool
}

// NewDaoIdea create a db connect to table Idea.
func NewDaoComment() (d *DaoComment) {
	d = &DaoComment{}
	d.redis = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.Conf.RedisAddr)
		},
	}
	return
}

func (this *DaoComment) GetCommentsOfIdea(iid int64) []*Comment {
	c := this.redis.Get()
	defer c.Close()

	list := make([]*Comment, 0)
	r, err := redis.String(c.Do("get", iid))
	if err != nil { // 空的
		return list
	} else {
		err = json.Unmarshal([]byte(r), &list)
		if err != nil {
			fmt.Printf("JSON Unmarshal failed : %s", err.Error())
			return nil
		}
		return list
	}
}

func (this *DaoComment) AddCommentToIdea(iid int64, openid string, content string) {
	c := this.redis.Get()
	defer c.Close()

	r, err := redis.String(c.Do("get", iid))
	list := make([]*Comment, 0)
	jsons := make([]byte, 0)
	item := &Comment{iid, openid, content, time.Now()}
	if err != nil { // 无评论
		list = append(list, item)
		jsons, err = json.Marshal(list)
		if err != nil {
			fmt.Printf("JSON Marshal failed : %s", err.Error())
			return
		}
	} else { // 有评论
		err = json.Unmarshal([]byte(r), &list)
		if err != nil {
			fmt.Printf("JSON Unmarshal failed : %s", err.Error())
			return
		}
		list = append(list, item)
		jsons, err = json.Marshal(list)
		if err != nil {
			fmt.Printf("JSON Marshal failed : %s", err.Error())
			return
		}
	}

	_, err = c.Do("set", iid, string(jsons))
	if err != nil {
		fmt.Println("c.Do err ", err)
		return
	}
}

func (this *DaoComment) Close() {
	if this.redis != nil {
		this.redis.Close()
	}
}
