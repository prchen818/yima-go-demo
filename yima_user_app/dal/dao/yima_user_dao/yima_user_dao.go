package yima_user_dao

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"yima_user_app/dal/model"
	"yima_user_app/dal/query"
)

var (
	Dao *YimaUserDao
)

func init() {
	d, err := gorm.Open(mysql.Open("root:mysql.peifang.para.party@(43.248.96.203:30306)/yima_user_db?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err.Error())
	}
	Dao = &YimaUserDao{
		db: d,
	}
}

type YimaUserDao struct {
	db *gorm.DB
}

func (d *YimaUserDao) GetUserInfoByUid(ctx context.Context, id string) (*model.SysUser, error) {
	q := query.Use(d.db).SysUser
	user, err := q.WithContext(ctx).Where(q.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}
