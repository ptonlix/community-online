# Community-Online
简体中文 | [English](./README-en.md)
<p>
	<p align="center">
		<img src="https://img.gejiba.com/images/77641b7520af20c7e065ffad2a7e5480.png" height=180px>
	</p>
	<p align="center">
		<font size=6 face="宋体">凌逸智慧社区开源平台</font>
	</p>
    <p align="center">
    打造一站式智慧社区平台，服务社区经济圈，让社区没有难做生意
    </p>
</p>
<p align="center">
<img alt="Go" src="https://img.shields.io/badge/Go-1.18%2B-blue">
<img alt="Mysql" src="https://img.shields.io/badge/Mysql-5.7%2B-brightgreen">
<img alt="Redis" src="https://img.shields.io/badge/Redis-6.2%2B-yellowgreen">
<img alt="go-zero" src="https://img.shields.io/badge/go--zero-1.4.1-orange">
<img alt="license" src="https://img.shields.io/badge/license-GPL-lightgrey">
</p>

## 环境依赖

Mysql v5.7+

Redis v6.2.6

etcd v3.5.5

go-zero v1.4.1

Note: 请选择合适的版本进行安装,否则服务启动失败。 go-zero可以参考[官方文档](https://github.com/zeromicro/go-zero)

## 开发部署

```
1.克隆项目
git clone https://github.com/ptonlix/community-online.git

2. 修改配置文件
community-online/app/usercenter/cmd/api/etc/usercenter.yaml
community-online/app/usercenter/cmd/rpc/etc/usercenter.yaml
community-online/app/sms/cmd/rpc/etc/sms.yaml

修改上述三个配置的Mysql Redis etcd 等地址 
还有微信小程序密钥，腾讯短信云密钥等信息

3.运行
分别在各种目录中启动服务，如
cd community-online/app/usercenter/cmd/api/
go run usercenter.go -f  etc/usercenter.yaml

4.测试
可以使用Apifox等工具，进行API请求测试验证。

```
## 目前实现模块
### 1.用户中心

可以跳转查看用户中设计和详细说明->[README](./doc/usercenter/README-cn.md)

### 2.支付中心

待开发～ 敬请期待


<p align="center">
  <b>SPONSORED BY</b>
</p>
<p align="center">
   <a href="https://www.gogeek.com.cn/" title="gogeek" target="_blank">
      <img height="200px" src="https://img.gejiba.com/images/96b6d150bd758b13d66aec66cb18044e.jpg" title="gogeek">
   </a>
</p>