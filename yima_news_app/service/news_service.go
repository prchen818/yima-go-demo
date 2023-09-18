package service

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"yima_news_app/dao"
	"yima_news_app/dao/query"
	"yima_news_app/kitex_gen/yima/news"
)

type NewsService struct {
	db *gorm.DB
}

func NewNewsService() *NewsService {
	return &NewsService{
		db: dao.Db,
	}
}

func (s *NewsService) GetAllNewsByUserId(ctx context.Context, id string) ([]*news.SciNews, error) {
	q := query.Use(s.db).SciNews
	newsList, err := q.WithContext(ctx).Where(q.UserID.Eq(id)).Find()
	if err != nil {
		return nil, err
	}
	var ret []*news.SciNews
	for _, n := range newsList {
		tmp := &news.SciNews{}
		err := copier.Copy(tmp, n)
		if err != nil {
			return nil, err
		}
		ret = append(ret, tmp)
	}
	return ret, nil
}
