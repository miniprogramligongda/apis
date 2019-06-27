#!/bin/bash

go get -u "github.com/garyburd/redigo/redis"
go get -u "github.com/jinzhu/gorm"
go get -u "github.com/go-sql-driver/mysql"
go get -u "github.com/labstack/echo"
go get -u "github.com/medivhzhan/weapp"

go build
