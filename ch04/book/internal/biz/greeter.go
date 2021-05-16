package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	ID int64
	Name string
	Description string
	CreatedAt time.Time
	DeletedAt time.Time
}

type GreeterRepo interface {
	CreateGreeter(context.Context, *Greeter) (int64, error)
	UpdateGreeter(context.Context, int64, *Greeter) error
	GetGreeter(context.Context, int64) (*Greeter, error)
	DeleteGreeter(context.Context, int64) error
}

type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper("usecase/greeter", logger)}
}

func (uc *GreeterUsecase) Create(ctx context.Context, g *Greeter) (int64, error) {
	return uc.repo.CreateGreeter(ctx, g)
}

func (uc *GreeterUsecase) Update(ctx context.Context, id int64, g *Greeter) error {
	return uc.repo.UpdateGreeter(ctx, id, g)
}

func (uc *GreeterUsecase) GetGreeter(ctx context.Context, id int64) (*Greeter, error) {
	return uc.repo.GetGreeter(ctx, id)
}

func (uc *GreeterUsecase) DeleteGreeter(ctx context.Context, id int64) error {
	return uc.repo.DeleteGreeter(ctx, id)
}
