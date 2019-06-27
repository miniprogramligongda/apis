#!/bin/bash -v

echo "set proxy"
#nohup ss-local -c ./script/ss.json &
ps -aux | grep ss-local

#export http_proxy=socks5://127.0.0.1:1086
#export https_proxy=socks5://127.0.0.1:1086
echo $http_proxy
echo $https_proxy

echo "/usr/local/go/bin/go  env"
go  env

echo "getting redis"
go  get -u "github.com/garyburd/redigo/redis"
echo "got redis"

echo "getting gorm"
go  get -u "github.com/jinzhu/gorm"
echo "got gorm"

echo "getting mysql"
go  get -u "github.com/go-sql-driver/mysql"
echo "got mysql"

echo "getting echo"
go  get -u "github.com/labstack/echo"
echo "got echo"

echo "getting weapp"
go  get -u "github.com/medivhzhan/weapp"
echo "got weapp"

#killall apis
#echo "kill the last edition apis"

go  build
echo "build OK!"

echo "unset proxy"
unset http_proxy https_proxy
#killall ss-local
