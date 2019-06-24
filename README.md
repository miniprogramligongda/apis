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

<div align=center>![ECHO Framework](https://cdn.labstack.com/images/echo-logo.svg)</div>

<div align=center>![八种主流框架对比](http://qwding.github.io/img/golang_framwork_benchmark.png)</div>

原来也有考虑过Beego框架，因为Beego框架是国人开发，并保持100%Go语言编写的『准优良血统』，拥有原生的中文文档和大量的学习例子。但是经过考察，发现Beego框架也只是将PHP语言上那一套MVC思想使用Go语言复现了一遍，并且并没有使用Go语言对比其他语言的优势。所以使用这就是为什么不实用Beego，并且评价它为『准优良血统』的原因。

<div align=center>![Beego Framework](https://beego.me/static/img/beego_purple.png)</div>

在数据库的选择上，我们选择了开源并且使用面广泛的MySQL关系式数据库，因为就互联网上主流的数据库操作来说，对MySQL数据库的操作的兼容性是最高的，所以选择了MySQL。

<div align=center>![MySQL](https://upload.wikimedia.org/wikipedia/zh/thumb/6/62/MySQL.svg/272px-MySQL.svg.png)</div>

原先Redis也是备选之一，因为Radis拥有『世界上最快数据库』的美誉，但是因为它是单线程，并且数据存取关系是『键值对』，对表形数据有些不友好。所以遂放弃之。

<div align=center>![Redis](https://redis.io/images/redis-white.png)</div>

## 框架设计

总共分为两个模块三层结构，分别为main(入口)，router(路由)，sql(数据库控制器)。

入口控制简易HTTP服务，以及整个后端的配置参数设置。入口接到的http操作请求之后，会调用路由模块新开线程进行业务处理。在业务实现代码中会调用相应的数据库控制模块，对数据库做出相应的修改。

<div align=center>![Redis](https://github.com/miniprogramligongda/apis/blob/readme_dev/imgs/framework.png?raw=true)</div>

## Contributor

1. NikoKVCS
2. SmallXeon