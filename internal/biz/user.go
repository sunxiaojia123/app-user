package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// 定义返回数据结构体
type User struct {
	ID       int64
	Mobile   string
	Password string
	NickName string
	Birthday int64
	Gender   int32
	Role     int
}

type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}
