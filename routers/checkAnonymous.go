package routers

import "../dao"

// checkAnonymous
// judge if openidB is a friend of openid
// return UserInfo of A, if not friend, the avatar url will be a Anonymous picture
func checkAnonymous(openid string, friends []*dao.Friend) (*dao.UserInfo, int) {
	d := dao.NewDaoUserInfo()
	defer d.Close()
	user, err := d.FindByOpenid(openid)
	if err != nil || user == nil {
		user = &dao.UserInfo{}
		user.Openid = openid
		user.AvatarUrl = "https://gss0.baidu.com/-vo3dSag_xI4khGko9WTAnF6hhy/zhidao/wh%3D600%2C800/sign=e746333c8cd4b31cf0699cbdb7e60b47/d788d43f8794a4c2d0b5cf5409f41bd5ad6e393e.jpg"
		user.Nickname = "匿名用户"
		user.Gender = 1
		return user, 0
	}

	if friends != nil && len(friends) > 0 {
		for _, v := range friends {
			if v.ObjectOpenid == openid {
				return user, 1
			}
		}
	}

	user.AvatarUrl = "https://gss0.baidu.com/-vo3dSag_xI4khGko9WTAnF6hhy/zhidao/wh%3D600%2C800/sign=e746333c8cd4b31cf0699cbdb7e60b47/d788d43f8794a4c2d0b5cf5409f41bd5ad6e393e.jpg"
	user.Nickname = "匿名用户"

	return user, 0
}
