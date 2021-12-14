# 适合创业的golang微服务框架go-zero + 金光灿灿的gorm V2实践

# 为什么使用go-zero
## 可以有第三个选择
+ golang圈子不大,微服务框架框架屈指可数:除了go-micro、go-kit,几乎没有其他选择。go-zero为此提供第三个可能。
+ go-micro 对webapi支持并不友好，需要运行micro指令,难以定制
## 创业产品需要一款支持业务成长的框架
我们到底需要什么样的框架?我们需要一款可以支持业务成长的框架!产品运营初期,比如需求验证阶段，我们并不需要采取微服务架构，因为运营成本太高。一款单体应用足以应付。随着业务发展,微服务成为必要,此时我们希望不进行太多的代码工作量，柔性升级。这正是go-zero价值所在

## go-zero是什么
以下安利是copy的的，具体参考`https://github.com/tal-tech/go-zero`
go-zero是一个集成了各种工程实践的包含web和rpc框架，有如下主要特点：

+ 强大的工具支持，尽可能少的代码编写
+ 极简的接口
+ 完全兼容net/http
+ 支持中间件，方便扩展
+ 高性能
+ 面向故障编程，弹性设计
+ 内建服务发现、负载均衡
+ 内建限流、熔断、降载，且自动触发，自动恢复
+ API参数自动校验
+ 超时级联控制
+ 自动缓存控制
+ 链路跟踪、统计报警等
+ 高并发支撑，稳定保障了晓黑板疫情期间每天的流量洪峰

# 怎么用
在阅读本文档前,请将golang 升级到`go14`及以上版本,并开启go module支持,GO14以上只是为了支持Gorm

```sh
export GOPROXY=https://goproxy.io,direct
export GO111MODULE=on 
```
## 安装goctl
goctl是go-zero配套的代码生成器,偷懒神器,毕竟`写代码大多时间是体力活`
如何安装呢?先把源代码下载下来吧!

```bash
git clone https://github.com/tal-tech/go-zero
cd go-zero/tools/goctl
go build goctl.go
```

最后生成goctl.exe  复制到`$gopath/bin`下

## goctl指令说明
自行浏览文档吧`https://github.com/tal-tech/go-zero/blob/master/tools/goctl/goctl.md`

本文用到指令如下
```bash
goctl api      go       -api             open.api            -dir                     .
```

```bash
#代码说明如下
goctl api      go       -api             open.api            -dir                     .
 |      |        |         |                 |                  |                      | 
      生成api  go语言     指定api模板文件   模板文件名称         指定生成代码存放路径     当前文件夹
```

# 创建项目
## 生成go.mod文件
以如下指令创建项目
```bash
mkdir hello
cd hello
go mod init  hello
```
## 定义hello.api
本文设计API如下
|描述|格式|方法|参数|返回|
|----|----|----|----|----|
|用户注册|/open/register|post|Email:手机号,passwd:密码,code:图片验证码|id:用户ID,token:用户token|
|用户登录|/open/authorization|post|Email:手机号,passwd:密码,code:图片验证码|id:用户ID,token:用户token|
|图片验证码请求|/open/verify|get|ticket:图片验证码的id|data:base64格式的图片|

根据以上描述,书写api的模板文件如下

```golang

type (
	UserOptReq struct {
		Email string `json:"Email"`
		passwd string `json:"passwd"`
		code   string `json:"code"`
	}

	UserOptResp struct {
		id    uint   `json:"id"`
		token string `json:"token"`
	}
	//图片验证码支持
	VerifyReq struct {
		ticket string `json:"ticket"`
	}
	//图片验证码支持
	VerifyResp struct {
		data string `json:"data"`
	}
)

service open-api {
	@doc(
        summary: 公开的api函数
        desc: >
        register 用户注册,
        authorization 用户登录,
        verify 图片验证码接口
    )
	@server(
		handler: registerHandler
		folder: open
	)
	post /open/register(UserOptReq) returns(UserOptResp)
	
	
	@server(
		handler: authorizationHandler
		folder: open
	)
	post /open/authorization(UserOptReq) returns(UserOptResp)

	@server(
		handler: verifyHandler
		folder: open
	)
	post /open/verify(VerifyReq) returns(VerifyResp)
	
}

```
注意
+ 一个文件里面只能有一个service
+ 工具最后会以type里面模型为样板生成各种结构体,所以参数和结构体保持一致即可
+ 如果我们需要分文件夹管理业务, 可以用folder属性来定义
## 生成代码
采用如下指令生成代码
```bash
goctl api  go   -api   open.api   -dir  .
```
最后代码结构如下
```bash
#tree /F /A
|   go.mod
|   go.sum
|   hello.api
|   open.go
|
+---etc
|       open-api.yaml
|
\---internal
    +---config
    |       config.go
    |
    +---handler
    |   |   routes.go
    |   |
    |   \---open
    |           authorizationhandler.go
    |           registerhandler.go
    |           verifyhandler.go
    |
    +---logic
    |   \---open
    |           authorizationlogic.go
    |           registerlogic.go
    |           verifylogic.go
    |
    +---svc
    |       servicecontext.go
    |
    \---types
            types.go                                          
```
运行一下
```bash
go run open.go
```
测试一下
```bash
curl http://127.0.0.1:8888/open/register -X POST -H "Content-Type: application/json" -d {\"Email\":\"15367151352\",\"passwd\":\"testpwd\",\"code\":\"asdf\"}
{"id":0,"token":""}
```

# 集成明星产品Gorm V2
金珠大佬升级了Gorm V2集成测试一下吧

## 配置文件
配置文件在`etc/open-api.yaml`
```
Name: open-api
Host: 0.0.0.0
Port: 8888
DataSourceName: root:1D007648b4f8@(127.0.0.1:3306)/gozero?charset=utf8
```
在`etc/open-api.yaml`中添加参数DataSourceName,
在`internal/config/config.go`中添加DataSourceName
```golang
type Config struct {
	rest.RestConf
	DataSourceName string
}

```
关于配置文件,系统内置了一部分关键字 如Cache,资料不多;基本上可以随便配置,然后在Conf中定义同名变量即可。

## 启动Gorm支持
修改`svc/servicecontext.go`代码如下
```golang
package svc

import (
	"hello/internal/config"
	"hello/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
    //启动Gorm支持
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tech_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,    // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
    })
    //如果出错就GameOver了
	if err != nil {
		panic(err)
    }
    //自动同步更新表结构,不要建表了O(∩_∩)O哈哈~
	db.AutoMigrate(&models.User{})

	return &ServiceContext{Config: c, DbEngin: db}
}

```
## 新建模型文件
新建`models\models.go`文件
```golang
//models\models.go文件
package models

import (
	"errors"
	"hello/internal/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"index:Email;type:varchar(13)"`
	Passwd string `gorm:"type:varchar(64)"`
}
//在创建前检验验证一下密码的有效性
func (u *User) BeforeCreate(db *gorm.DB) error {
	if len(u.Passwd) < 6 {
		return errors.New("密码太简单了")
    }
    //对密码进行加密存储
	u.Passwd = utils.Password(u.Passwd)
	return nil
}
```
utils.Password是我们编写的工具包,代码如下

```golang
package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//密码加密
func Password(plainpwd string) string {
    //谷歌的加密包
	hash, err := bcrypt.GenerateFromPassword([]byte(plainpwd), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	return encodePWD
}
//密码校验
func CheckPassword(plainpwd, cryptedpwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(cryptedpwd), []byte(plainpwd)) //验证（对比）
	return err == nil
}

```
## 实现业务逻辑
在`logic\open\registerlogic.go`中修改代码如下
```golang
package logic

import (
	"context"

	"hello/internal/models"
	"hello/internal/svc"
	"hello/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx context.Context
	logx.Logger
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.UserOptReq) (*types.UserOptResp, error) {
	user := models.User{
		Email: req.Email,
		Passwd: req.Passwd,
	}
	result := l.svcCtx.DbEngin.Create(&user)
	return &types.UserOptResp{
		Id: user.ID,
	}, result.Error
}

```
+ RegisterLogic中添加svcCtx *svc.ServiceContext,因为需要用到里面的DbEngin
+ NewRegisterLogic 配置svcCtx
+ 在Register函数中实现逻辑`result := l.svcCtx.DbEngin.Create(&user)`

## 测试一下
```bash
>curl http://127.0.0.1:8888/open/register -X POST -H "Content-Type: application/json" -d {\"Email\":\"15367151352\",\"passwd\":\"testpwd\"}
{"id":3,"token":""}
```

# 期待更新的功能点
## go-zero
###  接口定义希望支持多种content-type
    UserOptReq struct {
		Email string `json:"Email" form:"Email" xml:"Email"`
		passwd string `json:"passwd" form:"passwd" xml:"passwd"`
		code   string `json:"code" form:"code" xml:"code"`
	}
一种可能的解决方法是
修改`github.com/tal-tech/go-zero/rest/httpx/requests.go`中的`Parse`成如下模型
```golang
func Parse(r *http.Request, v interface{}) error {
	if err := ParsePath(r, v); err != nil {
		return err
	}
	if strings.Contains(r.Header.Get(ContentType), multipartFormData) {
		return ParseForm(r, v)
	} else if strings.Contains(r.Header.Get(ContentType), urlencodeformdata) {
		return ParseForm(r, v)
	} else if strings.Contains(r.Header.Get(ContentType), applicationjson) {
		return ParseJsonBody(r, v)
	} else {
		return errors.New("不支持的请求类型")
	}
}
```
### 支持一个文件多个方法
比如如下写法,则生成俩个方法在verifyHandler.go文件中
```
	@server(
		handler: verifyHandler
		folder: open
	)
	post /open/verify(VerifyReq) returns(VerifyResp)
    post /open/authorization(UserOptReq) returns(UserOptResp)
```

## gorm v2
###  建议默认`SingularTable`属性为`true`
```golang
NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tech_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,    // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
```
### 建议增强缓存功能
建议提供缓存如redis/memcache/内存缓存支持

# 本文代码获取
关注公众号`betaidea` 输入gozero或者gormv2即可获得