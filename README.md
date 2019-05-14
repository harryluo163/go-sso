# go-sso

这个一个使用Go语言的sso使用账户密码授权例子
数据库使用mysql 保存用户信息

## 涉及到包
* 利用Beego创建登录页面和简单的api
* "/login"  登录页面
*  "/Success" 登录成功页面
*   "/serviceValidate" 验证界面
*  使用beego.InsertFilter 进行授权拦截
  
* jwt-go 创建签名
  
  涉及到创建Token,和检验Token

## 编译
 go build main.go

## 使用
第三方直接点击跳转到
http://127.0.0.1:8080/login?service=http://www.xxxx.com.cn/ceo_data 

账户密码验证成功后，会跳转到service中，并带上token

第三方系统只需要通过接口”/serviceValidate“ 验证就可以判断是否登录，并且能得到用户信息哦。

数据库密码账户都有 可以用作测试
