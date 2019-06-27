#!/bin/bash

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
