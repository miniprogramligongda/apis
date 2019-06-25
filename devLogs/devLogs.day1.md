# day 1 (2019.6.24 Mon)

version：0.0.1

## 确认开发因素

- 语言：Golang
  - web框架：Echo(Golang)

- 数据库：MySQL
- 服务端架构：VPS(1u1g25g100mbps)+Ubuntu16.04_x86_64

## 确定设计框架

- main(top) -> router -> sql

main 调用 router 开线程处理 http 请求，业务代码调用 sql controller 对数据库进行增删改查

- router -> routers

router function 集中设置 routers 中的业务代码的路由路径

## 今日完成

- 总入口编写

- router 基础设施

  - 将路由router放在单独一个包内，并且每个路由函数一个代码文件，从而有利于避免多人同时修改在同一个文件修改router导致merge conflicts

- readme 文档

- mysql 基础设施

  - 将mysql的操作进行封装，以类函数接口的方式提供增删改查操作。

- 服务器环境建设

  - 分配两个云服务器，一个用于做数据库服务器，一个用于运行后端代码

    
