package yima_user_service

import (
	"context"
	"yima_user_app/dal/dao/yima_user_dao"
	"yima_user_app/dal/model"
)

var (
	Service *YimaUserService
)

type YimaUserService struct {
	yimaUserDao *yima_user_dao.YimaUserDao
}

func init() {
	Service = &YimaUserService{
		yimaUserDao: yima_user_dao.Dao,
	}
}

func (s *YimaUserService) GetUserInfo(ctx context.Context, id string) (*model.SysUser, error) {
	return s.yimaUserDao.GetUserInfoByUid(ctx, id)
}
