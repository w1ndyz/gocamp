package data

import (
	"book/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper("data/greeter", logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) (int64, error) {
	gr, err := r.data.db.Greeter.Create().SetName(g.Name).SetDescription(g.Description).Save(ctx)
	if err != nil {
		return -1, err
	}

	return gr.ID, nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, id int64, g *biz.Greeter) error {
	greeter, err := r.data.db.Greeter.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = greeter.Update().SetName(g.Name).SetDescription(g.Description).SetUpdatedAt(time.Now()).Save(ctx)
	return err
}

func(r *greeterRepo) GetGreeter(ctx context.Context, id int64) (*biz.Greeter, error) {
	greeter, err := r.data.db.Greeter.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Greeter{
		ID: greeter.ID,
		Name: greeter.Name,
		Description: greeter.Description,
		CreatedAt: greeter.CreatedAt,
	}, nil
}

func (r *greeterRepo) DeleteGreeter(ctx context.Context, id int64) error {
	return r.data.db.Greeter.DeleteOneID(id).Exec(ctx)
}