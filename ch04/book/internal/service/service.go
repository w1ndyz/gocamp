package service

import (
	pb "book/api/greeter/v1"
	"book/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService)

type GreeterService struct {
	pb.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}
