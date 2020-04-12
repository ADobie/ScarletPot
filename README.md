# 项目简介

本项目为Scarletwaf 核心功能之一，目的是保护站点免受恶意攻击同时引诱攻击者进入蜜罐，收集攻击者信息与攻击手法，并及时向网站管理员发送预警。



# 基础架构

![image-20200329195610564](http://pic.lwh.red/image-20200329195610564.png)

- 初始化阶段：从`config.ini`中读取蜜罐配置文件，并将配置信息存入`scarlet_config`表中。同时从表中读取配置并缓存，以减少反复查询数据库。
- 蜜罐服务启动阶段：协程启动各个独立的蜜罐服务以及模拟漏洞，同时将服务启动状态写入log
- 服务运行阶段：攻击者访问蜜罐，记录攻击信息并上报



# 项目结构

```ini
├── ./README.md
├── ./conf
│   └── ./conf/config.ini
├── ./db
│   ├── ./db/mysql
│   │   ├── ./db/mysql/honey.sql
│   │   └── ./db/mysql/mysql.go
│   └── ./db/sqlite
│       ├── ./db/sqlite/honey.db
│       └── ./db/sqlite/sqlite.go
├── ./go.mod
├── ./init
│   └── ./init/init.go
├── ./logs
├── ./main.go
├── ./router
│   └── ./router/router.go
├── ./services
│   ├── ./services/elasticsearch
│   ├── ./services/ftp
│   ├── ./services/memcache
│   ├── ./services/mysql
│   ├── ./services/redis
│   ├── ./services/ssh
│   ├── ./services/telnet
│   └── ./services/web
├── ./utils
│   ├── ./utils/cache
│   │   └── ./utils/cache/cache.go
│   └── ./utils/log
└── ./web
    ├── ./web/thinkphp
    ├── ./web/typecho
    └── ./web/wordpress
```



# 说明

- 支持用户自定义服务漏洞 尽可能多的预设漏洞
- 支持用户自定义web框架/页面模板 尽可能多的预设模板
- 数据库层面操作都在panel进行,蜜罐只负责数据上报