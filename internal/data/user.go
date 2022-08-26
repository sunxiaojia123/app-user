package data

import (
	"context"
	"time"
	"user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// 定义数据表结构体
type User struct {
	ID         int64
	Mobile     string
	Password   string
	NickName   string
	Birthday   string
	Gender     int32
	Role       int
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoCreateTime"`
	DeleteTime gorm.DeletedAt
}

func (m *User) TableName() string {
	return "user"
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo . 这里需要注意，上面 data 文件 wire 注入的是此方法，方法名不要写错了
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUser .
func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	var user User
	// 验证是否已经创建
	result := r.data.db.Where(&biz.User{Mobile: u.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.Mobile = u.Mobile
	user.NickName = u.NickName
	user.Password = u.Password // 密码加密
	// user.Password = encrypt(u.Password) // 密码加密
	res := r.data.db.Create(&user)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return &biz.User{
		ID:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     user.Role,
	}, nil
}

// // Password encryption
// func encrypt(psd string) string {
// 	options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
// 	salt, encodedPwd := password.Encode(psd, options)
// 	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
// }
