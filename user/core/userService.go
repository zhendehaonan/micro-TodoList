package core

import (
	"context"
	"errors"
	"user/model"
	"user/service"
)

// 登录
func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	var user *model.User
	resp.Code = 200
	var err error
	Db := model.NewDBClient(ctx)
	err = Db.Where("user_name=?", req.UserName).First(&user).Error
	if err != nil {
		if err != nil {
			resp.Code = 400
			return nil
		}
	}
	if user.CheckPassword(req.Password) == false {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}

// 注册
func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	var err error
	if req.Password != req.PasswordConfirm {
		err = errors.New("两次密码输入不一致,请重新输入")
		return err
	}
	Db := model.NewDBClient(ctx)
	var count int64
	if err = Db.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err = errors.New("用户名已存在")
		return err
	}
	user := &model.User{
		UserName: req.UserName,
	}
	//密码加密
	if err = user.SetPassword(req.Password); err != nil {
		return err
	}
	//创建用户
	if err = Db.Model(&model.User{}).Create(&user).Error; err != nil {
		return err
	}
	resp.UserDetail = BuildUser(user)
	resp.Code = 200
	return nil
}

// 序列化
func BuildUser(user *model.User) *service.UserModel {
	return &service.UserModel{
		ID:        uint32(user.ID),
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
}
