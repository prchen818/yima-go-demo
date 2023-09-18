package service

import (
	"context"
	"gorm.io/gorm"
	"yima_user_app/dao"
	"yima_user_app/dao/model"
	"yima_user_app/dao/query"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		db: dao.Db,
	}
}

func (s *UserService) GetUserInfoByUid(ctx context.Context, id string) (*model.SysUser, error) {
	q := query.Use(s.db).SysUser
	user, err := q.WithContext(ctx).Where(q.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}
