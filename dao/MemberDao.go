package dao

import (
	"cloudrestaurant/model"
	"cloudrestaurant/tool"
	"fmt"
	"log"
)

type MemberDao struct {
	*tool.Orm
}

// 根据用户名和密码查询
func (md *MemberDao) Query(name string, password string) *model.Member {
	var member model.Member
	_, err := md.Where("username = ? and password = ?", name, password).Get(&member)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return &member
}

// 确认手机号与验证码正确
func (md *MemberDao) ValidateSmscode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode
	if _, err := md.Where("phone = ? and code = ?", phone, code).Get(&sms); err != nil {
		fmt.Println(err.Error())
	}
	return &sms
}

// 根据手机号查询表记录
func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member
	if _, err := md.Where("mobile = ?", phone).Get(&member); err != nil {
		fmt.Println(err.Error())
	}
	return &member

}

// 新用户插入数据库
func (md *MemberDao) InsertMember(member model.Member) int64 {
	result, err := md.InsertOne(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

// 根据手机号插入验证码
func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
