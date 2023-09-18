package main

import (
	"context"
	"fmt"
	news "yima_news_app/kitex_gen/yima/news"
	"yima_news_app/service"
)

// YimaNewsImpl implements the last service interface defined in the IDL.
type YimaNewsImpl struct{}

// GetNewsByTitle implements the YimaNewsImpl interface.
func (s *YimaNewsImpl) GetNewsByTitle(ctx context.Context, request *news.GetNewsByTitleRequest) (resp *news.GetNewsByTitleResponse, err error) {
	// TODO: Your code here...
	return
}

// AddNews implements the YimaNewsImpl interface.
func (s *YimaNewsImpl) AddNews(ctx context.Context, request *news.AddNewsRequest) (resp *news.AddNewsResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAllNewsByUserId implements the YimaNewsImpl interface.
func (s *YimaNewsImpl) GetAllNewsByUserId(ctx context.Context, request *news.GetAllNewsByUserIdRequest) (resp *news.GetAllNewsByUserIdResponse, err error) {
	fmt.Print(request)
	id := request.UserId
	sciNews, err := service.NewNewsService().GetAllNewsByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	resp = &news.GetAllNewsByUserIdResponse{}
	resp.NewsList_ = sciNews
	return
}
