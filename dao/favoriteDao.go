package dao

import (
	"encoding/json"
	"fmt"

	"../conf"
	"github.com/garyburd/redigo/redis"
)

// DaoIdea struct contains a db object of grom.DB
type DaoFavorite struct {
	redis *redis.Pool
}

// NewDaoIdea create a db connect to table Idea.
func NewDaoFavorite() (d *DaoFavorite) {
	d = &DaoFavorite{}
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

func (this *DaoFavorite) AddFavToUser(Openid string, Iid int64) {
	c := this.redis.Get()
	defer c.Close()

	r, err := redis.String(c.Do("get", Openid))
	list := make([]int64, 0)
	jsons := make([]byte, 0)
	if err != nil {
		list = append(list, Iid)
		jsons, err = json.Marshal(list)
		if err != nil {
			fmt.Printf("JSON Marshal failed : %s", err.Error())
		}
	} else {
		err = json.Unmarshal([]byte(r), &list)
		if err != nil {
			fmt.Printf("JSON Unmarshal failed : %s", err.Error())
		}
		list = append(list, Iid)
		jsons, err = json.Marshal(list)
		if err != nil {
			fmt.Printf("JSON Marshal failed : %s", err.Error())
		}
	}

	_, err = c.Do("set", Openid, string(jsons))
	if err != nil {
		fmt.Println("c.Do err ", err)
		return
	}

}

func (this *DaoFavorite) DelFavOfUser(Openid string, Iid int64) {
	c := this.redis.Get()
	defer c.Close()

	r, err := redis.String(c.Do("get", Openid))
	list := make([]int64, 0)

	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(r), &list)
	if err != nil {
		fmt.Printf("JSON Unmarshal failed : %s", err.Error())
	}
	if len(list) == 0 {
		return
	}
	for i, v := range list {
		if v == Iid {
			// need testing
			list = append(list[:i], list[i+1:]...)
			break
		}
	}
	jsons, err := json.Marshal(list)
	if err != nil {
		fmt.Printf("JSON Marshal failed : %s", err.Error())
	}

	_, err = c.Do("set", Openid, string(jsons))
	if err != nil {
		fmt.Println("c.Do err ", err)
		return
	}
}

func (this *DaoFavorite) GetFavListOfUser(Openid string) []int64 {
	c := this.redis.Get()
	defer c.Close()

	r, err := redis.String(c.Do("get", Openid))
	list := make([]int64, 0)

	if err != nil {
		return list
	}

	err = json.Unmarshal([]byte(r), &list)
	if err != nil {
		fmt.Printf("JSON Unmarshal failed : %s", err.Error())
	}
	return list
}

func (this *DaoFavorite) Close() {
	if this.redis != nil {
		this.redis.Close()
	}
}
