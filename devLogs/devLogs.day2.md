# day 2

version：0.0.2

## 初步确定服务端接口

1. 收听

   - 数据：avatar, nickname, time, content, like(isLike & like amount), favorite(isFavorite & favorite amount)
   - GET：avatar, nickname, time, content, like(isLike & like amount), favorite(isFavorite & favorite amount)
   - POST：isLike(false <—> true), isFavorite(false <—> true)

2. 心声

   - 数据（POST）：content, time,

3. 我的

   - 数据（fromWX & POST）：wxId, userid, nickname, avatar

## 初步分工（分配接口）：

- DAO (Data Access Object) (include unit test)
  - nikokvcs: userinfo
  - Xeon: idea

- Router

  - nikokvcs:
    1. postIdea
    2. postUserinfo

  - Xeon:
    - loginCode

## 完成数据库建表

1. 用户表

   - openid		varchar(28)
   - avatar_url	varchar(128)
   - gender		tinyint(1)
   - nickname	varchar(32)

2. 心声表

   - iid			int(11)
   - openid		    varchar(28)
   - time			
   - content		varchar(1024)
   - like			smallint(6)
   - favorite		smallint(6)

## 今日完成

   - 服务器环境搭建：Ubuntu16.04+apis & Ubuntu16.04+MySQL5.7.26
   - 接口编写若干
   - 更新README.md
