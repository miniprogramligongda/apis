# apis | WXminiProgram: 分答 back-end part

The back-end of WXminiprogram `分答`. It is wrote by Golang and depends on Echo web framework and MySQL database.

微信小程序『分答』的后端。使用Go语言进行编写，基于Echo框架和MySQL数据库。

## TODO 

- [x] Router infrastructure
- [ ] MySQL infrastructure

- [ ] Indexly seperated cooperation
  - [ ] 『收听』
  - [ ] 『心声』
  - [ ] 『我的』

- [ ] Main project entrance

## Why doing like this?

因为在主流的8种Go语言的web框架中，Echo和Gin都有着别的框架无法媲美的有点：速度快 & 0内存占用，而项目编写人员中有使用Echo框架的经验，所以最后选择了使用Echo框架作为本项目的后端web框架。
![ECHO Framework](http://go-echo.org/images/echo.svg)

![八种主流框架对比](http://qwding.github.io/img/golang_framwork_benchmark.png)

原来也有考虑过Beego框架，因为Beego框架是国人开发，并保持100%Go语言编写的『准优良血统』，拥有原生的中文文档和大量的学习例子。但是经过考察，发现Beego框架也只是将PHP语言上那一套MVC思想使用Go语言复现了一遍，并且并没有使用Go语言对比其他语言的优势。所以使用这就是为什么不实用Beego，并且评价它为『准优良血统』的原因。

![Beego Framework](https://beego.me/static/img/beego_purple.png)

在数据库的选择上，我们选择了开源并且使用面广泛的MySQL关系式数据库，因为就互联网上主流的数据库操作来说，对MySQL数据库的操作的兼容性是最高的，所以选择了MySQL。

![MySQL](https://upload.wikimedia.org/wikipedia/zh/thumb/6/62/MySQL.svg/272px-MySQL.svg.png)

原先Redis也是备选之一，因为Radis拥有『世界上最快数据库』的美誉，但是因为它是单线程，并且数据存取关系是『键值对』，对表形数据有些不友好。所以遂放弃之。

![Redis](https://redis.io/images/redis-white.png)

框架设计

## Contributor

1. NikoKVCS
2. SmallXeon
