package service

import (
	"cloudrestaurant/dao"
	"cloudrestaurant/model"
	"cloudrestaurant/param"
	"cloudrestaurant/tool"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) GetUserInfo(userId string) *model.Member {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil
	}
	memberDao := dao.MemberDao{
		tool.DbEngine}
	return memberDao.QueryMemberById(int64(id))
}

func (ms *MemberService) UploadAvatar(userId int64, filename string) string {
	memberdao := dao.MemberDao{tool.DbEngine}
	result := memberdao.UpdateMemberAvatar(userId, filename)
	if result == 0 {
		return ""
	}
	return filename
}

// 实现用户登录
func (ms *MemberService) Login(name string, password string) *model.Member {
	//1.使用用户名 + 密码 查询用户信息，如果存在用户，返回用户信息
	md := dao.MemberDao{tool.DbEngine}
	member := md.Query(name, password)
	if member.Id != 0 {
		return member
	}
	//2.用户信息不存在，作为新用户保存到数据库中
	user := model.Member{}
	user.Username = name
	user.Password = tool.EncoderSha256(password)
	user.RegisterTime = time.Now().Unix()
	result := md.InsertMember(user)
	user.Id = result
	return &user
}

// 用户手机号+验证码的登陆
func (ms *MemberService) Smslogin(loginParam param.SmsLoginParam) *model.Member {
	//1.获取到手机号和验证码
	//2.验证手机号+验证码是否正确
	md := dao.MemberDao{tool.DbEngine}
	sms := md.ValidateSmscode(loginParam.Phone, loginParam.Code)
	if sms.Id == 0 {
		return nil
	}
	//3.根据手机号member表中查询数据
	member := md.QueryByPhone(loginParam.Phone)
	if member.Id != 0 {
		return member
	}
	//4.未查询到，新创建一个member记录，并保存
	user := model.Member{}
	user.Username = loginParam.Phone
	user.Mobile = loginParam.Phone
	user.RegisterTime = time.Now().Unix()
	user.Id = md.InsertMember(user)
	return nil
}

func (ms *MemberService) SendCode(phone string) bool {
	//1.产生验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	//2.调用短信服务sdk 完成发送11
	config := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)
	//3.接受返回结果，并判断发送状态
	response, err := client.SendSms(request)
	fmt.Println(response)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	if response.Code == "OK" {
		smscode := model.SmsCode{Phone: phone, Code: code, BizID: response.BizId, CreateTime: time.Now().Unix()}
		memberdao := dao.MemberDao{tool.DbEngine}
		result := memberdao.InsertCode(smscode)
		if result > 0 {
			return true
		}
	}
	return false
}
