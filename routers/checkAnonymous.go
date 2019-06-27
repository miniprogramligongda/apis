package routers

import "../dao"

// checkAnonymous
// judge if openidB is a friend of openid
// return UserInfo of A, if not friend, the avatar url will be a Anonymous picture
func checkAnonymous(openidA, openidB string) (*dao.UserInfo, int) {
	d := dao.NewDaoUserInfo()
	defer d.Close()
	user, err := d.FindByOpenid(openidA)
	if err == nil && user != nil {
	}

	return nil, 0
}
