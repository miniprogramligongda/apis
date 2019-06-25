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
- readme 文档
