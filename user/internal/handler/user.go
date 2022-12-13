package handler

import (
	"context"
	"user/internal/repository"
	"user/internal/service"
	"user/pkg/e"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// /请求参数req 去数据库 repository.User 查找； 不存在就是没注册，存在就返回user
func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.SUCCESS
	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = e.ERROR
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

// 注册=> 请求参数req,数据库user 创建一个新的数据；成功就注册成功，失败就返回error
func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.SUCCESS
	err = user.Create(req)
	if err != nil {
		resp.Code = e.ERROR
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

// /登出
func (*UserService) UserLogout(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	resp = new(service.UserDetailResponse)
	return resp, nil
}
