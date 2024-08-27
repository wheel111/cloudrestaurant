package controller

import (
	"cloudrestaurant/model"
	"cloudrestaurant/param"
	"cloudrestaurant/service"
	"cloudrestaurant/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmscode)
	engine.POST("/api/login_sms", mc.smSLogin)
	engine.GET("/api/captcha", mc.captcha)
	//login_pwd
	engine.POST("/api/login_pwd", mc.nameLogin)

	//头像上传
	engine.POST("/api/upload/avator", mc.uploadAvator)
}

// 头像上传
func (mc *MemberController) uploadAvator(c *gin.Context) {
	//1.解析上传的参数：file、user_id
	userId := c.PostForm("userId")
	fmt.Println(userId)
	file, err := c.FormFile("avatar")
	if err != nil || userId == "" {
		tool.Fail(c, "参数解析失败")
		return
	}
	//2.判断user_id对应的用户是否已经登陆
	session := tool.GetSession(c, "user_"+userId)
	if session == nil {
		tool.Fail(c, "参数不合法")
		return
	}
	var member model.Member
	json.Unmarshal(session.([]byte), &member)
	//3. file保存到本地
	filename := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = c.SaveUploadedFile(file, filename)
	if err != nil {
		tool.Fail(c, "头像更新失败")
		return
	}
	//4.将保存后的文件本地路径，保存到用户表中的头像字段
	memberService := service.MemberService{}
	path := memberService.UploadAvatar(member.Id, filename[1:])
	if path != "" {
		tool.Success(c, "http://localhost:8091"+path)
		return
	}
	//5.返回结果
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
	if member.Id != 0 {
		// 用户信息保存到session
		sess, _ := json.Marshal(member)
		err = tool.SetSession(c, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Fail(c, "登陆失败")
			return
		}
		tool.Success(c, &member)
		return
	}
	tool.Fail(c, "登陆失败")
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
		sess, _ := json.Marshal(member)
		err = tool.SetSession(c, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Fail(c, "登陆失败")
			return
		}
		tool.Success(c, member)
	} else {
		tool.Fail(c, "登录失败")
	}
}
