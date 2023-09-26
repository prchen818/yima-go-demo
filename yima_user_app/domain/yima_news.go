package domain

import (
	"context"
	"git.horsecoder.com/chenpeiran/yima.demo.news/kitex_gen/yima/news"
	"git.horsecoder.com/chenpeiran/yima.demo.news/yima_demo_news"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/client/callopt"
	"time"
)

func GetAllNewsByUid(ctx context.Context, id string) ([]*news.SciNews, error) {
	defer func() {
		if err := recover(); err != nil {
			logger.CtxErrorf(ctx, "panic!! err:%v", err)
			return
		}
	}()
	resp, err := yima_demo_news.RawCall.GetAllNewsByUserId(ctx, &news.GetAllNewsByUserIdRequest{UserId: id}, callopt.WithConnectTimeout(5*time.Second), callopt.WithRPCTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return resp.NewsList_, nil
	}
	return nil, nil

}
