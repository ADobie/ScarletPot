# 项目简介

本项目为Scarletwaf 核心功能之一，目的是保护站点免受恶意攻击同时引诱攻击者进入蜜罐，收集攻击者信息与攻击手法，并及时向网站管理员发送预警。


# 使用说明

帮助界面
```bash
$ ./scarletpot help   


███████╗  ██████╗  █████╗  ██████╗  ██╗     ███████╗ ████████╗   ██████╗  ██████╗ ████████╗
██╔════╝ ██╔════╝ ██╔══██╗ ██╔══██╗ ██║     ██╔════╝ ╚══██╔══╝   ██╔══██╗██╔═══██╗╚══██╔══╝
███████╗ ██║      ███████║ ██████╔╝ ██║     █████╗      ██║  *** ██████╔╝██║   ██║   ██║   
╚════██║ ██║      ██╔══██║ ██╔══██╗ ██║     ██╔══╝      ██║  *** ██╔═══╝ ██║   ██║   ██║   
███████║ ╚██████╗ ██║  ██║ ██║  ██║ ███████╗███████╗    ██║  *** ██║     ╚██████╔╝   ██║   
╚══════╝  ╚═════╝ ╚═╝  ╚═╝ ╚═╝  ╚═╝ ╚══════╝╚══════╝    ╚═╝  *** ╚═╝      ╚═════╝    ╚═╝

   run           Start all scarlet service
   install       Start install program
   version       Show scarletPot Version
   help          Show help

```

版本号
```bash
$ ./scarletpot version
ScarletPot v0.1 2020.4.19
By Annevi
```

初始化安装
```bash
$ ./scarletpot install


███████╗  ██████╗  █████╗  ██████╗  ██╗     ███████╗ ████████╗   ██████╗  ██████╗ ████████╗
██╔════╝ ██╔════╝ ██╔══██╗ ██╔══██╗ ██║     ██╔════╝ ╚══██╔══╝   ██╔══██╗██╔═══██╗╚══██╔══╝
███████╗ ██║      ███████║ ██████╔╝ ██║     █████╗      ██║  *** ██████╔╝██║   ██║   ██║   
╚════██║ ██║      ██╔══██║ ██╔══██╗ ██║     ██╔══╝      ██║  *** ██╔═══╝ ██║   ██║   ██║   
███████║ ╚██████╗ ██║  ██║ ██║  ██║ ███████╗███████╗    ██║  *** ██║     ╚██████╔╝   ██║   
╚══════╝  ╚═════╝ ╚═╝  ╚═╝ ╚═╝  ╚═╝ ╚══════╝╚══════╝    ╚═╝  *** ╚═╝      ╚═════╝    ╚═╝

------------------------------------- ABOUT ---------------------------------------
|                 author: Annevi                                                  |
|                 github: https://github.com/ScarletWaf/ScarletPot                |
-----------------------------------------------------------------------------------

----------------------------- Scarlet Pot installer -------------------------------

请选择基础语言 (1. zh-CN 2. en-US): 
....
```

启动服务
```bash
$ ./scarletpot run 
     
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> scarletpot/protocol/web.initJsonp.func1 (3 handlers)
[GIN-debug] POST   /api/a                    --> scarletpot/protocol/web.initJsonp.func2 (3 handlers)
[GIN-debug] Listening and serving HTTP on 0.0.0.0:8888
[INFO] Redis listening 0.0.0.0:63790
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> scarletpot/panel.(*Service).initRouter.func1 (3 handlers)
[GIN-debug] POST   /api/report               --> scarletpot/panel.(*Service).initRouter.func2 (3 handlers)
[GIN-debug] POST   /api/token/gen            --> scarletpot/panel.(*Service).initRouter.func3 (3 handlers)
[GIN-debug] GET    /api/token/sec            --> github.com/gin-gonic/gin.RecoveryWithWriter.func1 (2 handlers)
[GIN-debug] Listening and serving HTTP on 0.0.0.0:9000

```
# 基础架构

![image-20200329195610564](http://pic.lwh.red/image-20200329195610564.png)

- 初始化阶段：从`config.ini`中读取蜜罐配置文件，并将配置信息存入`scarlet_config`表中。同时从表中读取配置并缓存，以减少反复查询数据库。
- 蜜罐服务启动阶段：协程启动各个独立的蜜罐服务以及模拟漏洞，同时将服务启动状态写入log
- 服务运行阶段：攻击者访问蜜罐，记录攻击信息并上报



# 项目结构



# 模块说明
## Web钓鱼
### jsonp劫持：
> 攻击者 -> 钓鱼页面 -> 获取用户信息 -> 前端请求蜜罐返回用户数据

# 说明

- 支持用户自定义服务漏洞 尽可能多的预设漏洞
- 支持用户自定义web框架/页面模板 尽可能多的预设模板
- 数据库层面操作都在panel进行,蜜罐只负责数据上报