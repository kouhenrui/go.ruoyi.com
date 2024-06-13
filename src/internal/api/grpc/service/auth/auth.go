package auth

import (
	"context"
	pb "go.ruoyi.com/src/internal/api/proto/auth"
	//"go-microservice/api/proto/auth"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

func (a *AuthService) Login(context.Context, *pb.LoginReq) (*pb.LoginRes, error) {
	return &pb.LoginRes{Message: "login success"}, nil
}
