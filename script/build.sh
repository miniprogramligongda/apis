#!/bin/bash

echo "set proxy"
nohup ss-local -c ./script/ss.json &
export http_proxy=socks5://127.0.0.1:1080
export https_proxy=socks5://127.0.0.1:1080

go get -u "github.com/garyburd/redigo/redis"
echo "got redis"

go get -u "github.com/jinzhu/gorm"
echo "got gorm"

go get -u "github.com/go-sql-driver/mysql"
echo "got mysql"

go get -u "github.com/labstack/echo"
echo "got echo"

go get -u "github.com/medivhzhan/weapp"
echo "got weapp"

go build
echo "build OK!"

echo "unset proxy"
unset http_proxy https_proxy
killall ss-local
killall ss-local
killall ss-local
