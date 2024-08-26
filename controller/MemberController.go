package controller

import (
	"cloudrestaurant/param"
	"cloudrestaurant/service"
	"cloudrestaurant/tool"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmscode)
	engine.POST("/api/login_sms", mc.smSLogin)
	engine.GET("/api/captcha", mc.captcha)
	//login_pwd
	engine.POST("/api/login_pwd", mc.nameLogin)
}

// 用户名+密码、验证码登录
func (mc *MemberController) nameLogin(c *gin.Context) {
	//1.解析用户登录传递参数
	var loginParam param.LoginParam
	err := tool.Decode(c.Request.Body, &loginParam)
	if err != nil {
		tool.Fail(c, "参数解析失败")
		return
	}
	//2.登录
	ms := service.MemberService{}
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member == nil {
	}
}
func (mc *MemberController) captcha(c *gin.Context) {
	tool.GenerateCaptcha()
}

// http://localhost:8091/api/sendcode?phone=12323112321
func (mc *MemberController) sendSmscode(c *gin.Context) {
	// 发送验证码
	phone, exit := c.GetQuery("phone")
	if !exit {
		tool.Fail(c, "参数解析失败")
		return
	}
	ms := service.MemberService{}
	isSend := ms.SendCode(phone)
	if isSend {
		tool.Success(c, "参数解析成功")
	}
	tool.Fail(c, "参数解析失败")
}

// 手机号+短信 登陆的方法
func (mc *MemberController) smSLogin(c *gin.Context) {
	var smsLoginParam param.SmsLoginParam
	err := tool.Decode(c.Request.Body, &smsLoginParam)
	if err != nil {
		tool.Fail(c, "参数解析失败")
		return
	}
	// 完成手机+验证码登录
	us := service.MemberService{}
	member := us.Smslogin(smsLoginParam)
	if member != nil {
		tool.Success(c, member)
	} else {
		tool.Fail(c, "登录失败")
	}
}
