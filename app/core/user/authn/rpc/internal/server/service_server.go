// Code generated by goctl. DO NOT EDIT.
// Source: main.proto

package server

import (
	"context"

	"user/authn/rpc/internal/logic"
	"user/authn/rpc/internal/svc"
	"user/authn/rpc/pb"
)

type ServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedServiceServer
}

func NewServiceServer(svcCtx *svc.ServiceContext) *ServiceServer {
	return &ServiceServer{
		svcCtx: svcCtx,
	}
}

// Register 用户注册:
func (s *ServiceServer) Register(ctx context.Context, in *pb.UserRegisterReq) (*pb.UserRegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

// Login 用户登录:
func (s *ServiceServer) Login(ctx context.Context, in *pb.UserLoginReq) (*pb.UserLoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

// Logout 用户退出:
func (s *ServiceServer) Logout(ctx context.Context, in *pb.UserLogoutReq) (*pb.UserLogoutResp, error) {
	l := logic.NewLogoutLogic(ctx, s.svcCtx)
	return l.Logout(in)
}

// AuthMobile 移动端: 鉴权 + token
func (s *ServiceServer) AuthMobile(ctx context.Context, in *pb.AuthMobileReq) (*pb.AuthMobileResp, error) {
	l := logic.NewAuthMobileLogic(ctx, s.svcCtx)
	return l.AuthMobile(in)
}

// AuthWeb web端: 鉴权 + cookie
func (s *ServiceServer) AuthWeb(ctx context.Context, in *pb.AuthWebReq) (*pb.AuthWebResp, error) {
	l := logic.NewAuthWebLogic(ctx, s.svcCtx)
	return l.AuthWeb(in)
}
