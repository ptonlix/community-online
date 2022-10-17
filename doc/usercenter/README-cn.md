# 01.前言
大家好，我是Baird~

之前做了很多项目，基本上每个项目都要实现一遍用户的登陆认证，效率极低，还容易遗漏一些场景。网上有很多登陆认证的例子，但都是很基础的分享，要能做为商业应用，还需要自己进一步改造打磨。

考虑到登陆模块的通用性，我打算做一个可复用的登陆认证模块，方便后续项目快速开发，避免重复造轮子🚗

适合同学： 微服务开发、Go、go-zero、正在开发登陆模块的同学、感兴趣的朋友

# 02.用户中心设计
> 登陆认证模块，主要是用户行为，我们🙆这里统一将模块命名为用户中心

## 2.1. 选择开发语言和框架

考虑到可复用性，我们采用微服务开发的方式，单独打造用户中心, 而微服务这款go-zero目前是不二之选，故⬇️  
**开发语言**：Go  
**框架**: [go-zero](https://go-zero.dev/cn/)  
环境搭建请参考官方文档
## 2.3. 需求设计

用户中心的需求主要有以下几个方面：（这里考虑通用需求）
- 用户注册  
- 登陆认证  
  分为三种：   
  1.手机号+密码登陆   
  2.手机短信验证登陆   
  3.第三方授权登陆（目前只考虑微信小程序授权）
- 找回/修改密码
- 修改用户信息
- 统计在线用户

## 2.2. 数据库设计
数据库设计，考虑设计两张表，一张是用户表，一张是用户授权表  
用户表：保存用户信息  
用户授权表：保存用户登陆类型和登陆信息，如微信登陆，手机认证登陆
```sql
-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  `mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别 0:男 1:女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';

-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  `user_id` bigint NOT NULL DEFAULT '0',
  `auth_key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台唯一id',
  `auth_type` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台类型',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_type_key` (`auth_type`,`auth_key`) USING BTREE,
  UNIQUE KEY `idx_userId_key` (`user_id`,`auth_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户授权表';

```
## 2.3.用户中心总体架构

![用户中心总体架构](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/18b711bd1b954c06843dffcfc2481d3b~tplv-k3u1fbpfcp-zoom-1.image)

目前用户中心一共分为四部分  
1.API模块：  用户中心API  
2.RPC模块： 用户中心 和 短信模块  
3.服务发现模块： 采用etcd，用于RPC模块的注册和API发现RPC模块  
4.数据库： 采用mysql 和 redis

# 2.4.用户中心处理流程
![用户中心处理流程](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/8ec717be88ed471caa9094ccd00f1f0b~tplv-k3u1fbpfcp-zoom-1.image)

总体流程如上图所示
1. 启动基础服务，mysql redis etcd
2. RPC模块启动，注册到etcd
3. API模块启动
4. API请求->API模块->服务发现->RPC请求到RPC模块
5. 短信RPC模块需要对接到短信云服务，如腾讯云短信服务

# 2.5 API和Proto设计
```
syntax = "v1"

info(
	title: "用户中心服务"
	desc: "用户中心服务"
	author: "GoGeekBaird"
	email: "gogeek2022@163.com"
	version: "v1"
)

import (
	"user/user.api"
)

//============================> usercenter v1 <============================
//no need login
@server(
	prefix: usercenter/v1
	group: usercenter
)
service usercenter {
	
	@doc(
		summary: "register"
	)
	@handler register
	post /user/register (RegisterReq) returns (RegisterResp)
	
	@doc(
		summary: "login"
	)
	@handler login
	post /user/login (LoginReq) returns (LoginResp)
	
	@doc(
		summary: "wechat mini auth"
	)
	@handler wxMiniAuth
	post /user/wxmini-auth (WXMiniAuthReq) returns (WXMiniAuthResp)
	
	@doc(
		summary: "get message code"
	)
	@handler messagecode
	post /user/messagecode (MsgCodeReq)
	
	@doc(
		summary: "get captcha code"
	)
	@handler getCaptcha
	get /user/get-captcha returns(GetCaptchaResp)
	
	@doc(
		summary: "verfy captcha code"
	)
	@handler verfyCaptcha
	post /user/verfy-captcha(VerfyCaptchaReq) returns(VerfyCaptchaResp)
	
	@doc(
		summary: "change password"
	)
	@handler changePwd
	post /user/change-pwd(ChangePwdReq)
	
}

//need login
@server(
	prefix: usercenter/v1
	group: usercenter
	jwt: JwtAuth
	middleware: OnlineStatus // 路由中间件声明
)
service usercenter {
	
	@doc(
		summary: "get user info"
	)
	@handler detail
	get /user/detail returns (UserInfoResp)
	
	@doc(
		summary: "update user"
	)
	@handler updateUser
	post /user/update-user(UpdateUserReq)
	
	@doc(
		summary: "update user"
	)
	@handler deleteUser
	post /user/delete-user(DeleteUserReq)
	
	@doc(
		summary: "get online user list"
	)
	@handler getOnlineUser
	get /user/online-user returns(GetOnlineUserResp)
}
```
API设计采用 gozero API语法，使用goctl可以自动生成代码  
Proto设计，同样使用goctl可以自动生成代码 （proto文件太长就不贴啦，可以查看项目）
> goctl 采用1.4版本，模版代码见👇项目地址，可以自行替换  
> goctl 具体使用，请参考gozero官方文档

# 03.详细介绍

好，大体的设计按照上述进行即可，接下来我们分部分来对用户进行的功能开发做详细介绍。

## 3.1. 用户登陆

用户登陆，目前采用三种方式登陆：  
  1.手机号+密码登陆   
  2.手机短信验证登陆   
  3.第三方授权登陆（目前只考虑微信小程序授权）   
 
  登陆成功认证方式同一采用[JWT](https://jwt.io/) 
> JWT （JSON Web Token） 是目前最流行的跨域认证解决方案，是一种基于 Token 的认证授权机制。 从 JWT 的全称可以看出，JWT 本身也是 Token，一种规范化之后的 JSON 结构的 Token。  
  JWT 自身包含了身份验证所需要的所有信息，因此，我们的服务器不需要存储 Session 信息。这显然增加了系统的可用性和伸缩性，大大减轻了服务端的压力。  
  可以看出，**JWT 更符合设计 RESTful API 时的「Stateless（无状态）」原则** 。  
  并且， 使用 JWT 认证可以有效避免 CSRF 攻击，因为 JWT 一般是存在在 localStorage 中，使用 JWT 进行身份验证的过程中是不会涉及到 Cookie 的。
  
第一种：手机号+密码登陆   
    这种方式比较简单，是我们最常见的登陆方式。这种方式需要用户先注册，生成密码后再使用。
    通过手机号找到数据库中密码的HASH和输入的密码HASH比较相等即可认证通过。
    通过后返回生成JWT。   
    这里需要注意一点就是，为避免用户爆破输入，即重试多次用户名和密码。可以增加图形校验码验证。  
    图形校验码生成接口如下：  
    1.getCaptcha  
    2.verfyCaptcha     
    图形校验库采用： github.com/wenlng/go-captcha  
 
第二种：手机号+短信验证码方式  
    这种方式也是我们比较常用的方式：
    需要用短信云服务器，这里采用的是腾讯云，需要可自己开通，个人有100条免费短信。
    通过腾讯云的SDK文档开发接口。  
    短信验证码用redis缓存，记得校验完删除redis缓存即可。  

第三种：第三方授权登陆。  
    常见的有微信授权登陆、支付宝授权登陆等。
    这种方式按需进行开发，最近开发微信小程序比较多，所以选择的是微信小程序授权接口。  
    可以参考微信开发文档：[wx.login](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/login/wx.login.html)  
    
   这里微信开发，采用了比较热门的微信API库：github.com/silenceper/wechat/v2   
   其他登陆方式，大家可以按需自行开发。

除上述三种还有其他的登陆方式，如APP获取手机号，这种方式我接触不多，有待学习补充。

## 3.2 JWT处理流程
上一节介绍了JWT，JWT是无状态的和我们之前用的session、cookie方式有所不同。怎么刷新JWT，怎么判断JWT是否过期，是登陆认证要考虑清楚的，不然登陆的安全性就会受到影响。
> 本项目采用的 github.com/golang-jwt/jwt/v4 库生成和验证JWT

这里我给出我处理JWT的流程思路供大家参考，有可能考虑不全，也请大家指出。

1.首先返回给前端的数据有三个

-   Token String （Token字符串）
-   Token Expire (Token过期时间)
-   Token RefreshAfter （Token重刷新时间点）

前端将数据存储到localStorage

2.前端登陆处理流程

1.  没有Token --> 登陆
1.  有Token，但已过期 -->登陆
1.  有Token，没过期，但超过了RefreshAfter时间 -> 1⃣️ 免登陆 2⃣️获取新Token 3⃣️发送请求带上新Token
1.  有Token，没过期，RefreshAfter时间没到 -> 1⃣️免登陆，使用原来的Token
1.  用户主动注销，前端主动清除本地的Token数据

3.例子：实现登陆后7天免登陆，过程中持续有操作无需重新登陆，7天无操作则强制登陆

-   Token XXX （Token字符串）
-   Token 4233600 (Token过期时间)
-   Token 4233600/2 （Token重刷新时间点）

按照上述前端登陆处理流程，编写逻辑即可。


## 3.3 用户注册
针对手机号+密码登陆方式，需要进行用户自注册。  
用户注册逻辑比较简单，填写相应的用户信息即可。  
如有需要可以增加短信验证码验证手机，通过即可注册成功。  

## 3.4 找回/修改密码
这也是必要的的手段，密码忘记了，需要提供方式让用户能找回密码。  
这里采用就是手机验证码验证，输入新密码，保存成功即可。  
这里可以增加图形校验码验证，起一定防爆破作用。  

## 3.5 修改用户信息
这个比较简单的，提供接口允许用户修改头像和简介等基本信息。

## 3.6 统计在线用户
用户中心一个需求就有要统计在线的用户数量，以作为运营数据，优化业务和相关流程等等  
所以这里需要考虑三点  
1. 用户在线状态需要保存
2. 用户在线状态需要刷新或判断过期
3. 获取用户在线数量

针对第一点，这里可能有朋友已经发现了，这个和Session的功能很像，服务端保存 Session就能实现类似的功能。  
是的，要实现记录用户在线状态，我们就必须得保存一个用户唯一ID，作为用户在线记录。   
所以，我采用用户数据库id，做为Key保存用户在线数据到redis。  


针对第二点，我们可以采用路由中间件的方式，使得每个用户登陆后的操作，去刷新Key

```go
package middleware

import (
	"github/community-online/app/usercenter/cmd/rpc/usercenter"
	"github/community-online/common/ctxdata"
	"net/http"
)

type OnlineStatusMiddleware struct {
	UsercenterRpc usercenter.Usercenter
}

func NewOnlineStatusMiddleware(uRpc usercenter.Usercenter) *OnlineStatusMiddleware {
	return &OnlineStatusMiddleware{UsercenterRpc: uRpc}
}

func (m *OnlineStatusMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 刷新用户在线状态
		userId := ctxdata.GetUidFromCtx(r.Context())
		m.UsercenterRpc.FreshUserOnlineStatus(r.Context(), &usercenter.FreshUserOnlineStatusReq{
			UserId: userId,
		})
		// Passthrough to next handler if need
		next(w, r)
	}
}
```

> gozero路由中间件生成方式，可以参考官方文档


第三点，就比较明显了。我们只要把redis中相关的Key，query出来返回对应用户id列表即可。


以上就是整个用户中心详细介绍，文字较多，想详细了解的同学还需要结合代码过一遍。