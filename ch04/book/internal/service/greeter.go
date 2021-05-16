package service

import (
	pb "book/api/greeter/v1"
	"book/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{
		uc: uc,
		log: log.NewHelper("service/greeter", logger)}
}

func (s *GreeterService) CreateGreeter(ctx context.Context, req *pb.CreateGreeterRequest) (*pb.CreateGreeterReply, error) {
	id, err := s.uc.Create(ctx, &biz.Greeter{
		Name: req.Greeter.Name,
		Description: req.Greeter.Description,
	})
	return &pb.CreateGreeterReply{
		Result: &pb.Result{Code: "0"},
		Id: id,
	}, err
}

func (s *GreeterService) UpdateGreeter(ctx context.Context, req *pb.UpdateGreeterRequest) (*pb.UpdateGreeterReply, error) {
	err := s.uc.Update(ctx, req.Greeter.Id, &biz.Greeter{
		Name: req.Greeter.Name,
		Description: req.Greeter.Description,
	})
	return &pb.UpdateGreeterReply{Result: &pb.Result{Code: "0"}}, err
}

func (s *GreeterService) GetGreeter(ctx context.Context, req *pb.GetGreeterRequest) (*pb.GetGreeterReply, error) {
	g, err := s.uc.GetGreeter(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetGreeterReply{Greeter: &pb.GreeterData{
		Id: g.ID,
		Name: g.Name,
		Description: g.Description,
	}, Result: &pb.Result{Code: "0"}}, nil
}

func (s *GreeterService) DeleteGreeter(ctx context.Context, req *pb.DeleteGreeterRequest) (*pb.DeleteGreeterReply, error)  {
	err := s.uc.DeleteGreeter(ctx, req.Id)
	return &pb.DeleteGreeterReply{Result: &pb.Result{Code: "0"}}, err
}