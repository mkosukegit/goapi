package repository

import (
	"context"
)

// 外部から使える
type UserRepository interface {
	SelectAllUser(ctx context.Context)
}

type userRepository struct {
}

// コンストラクタ
func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (userRepository *userRepository) SelectAllUser(ctx context.Context) {
	Db.Query("SELECT * FROM user")
}
