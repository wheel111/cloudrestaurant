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
	"time"
)

type MemberService struct {
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
