package rpc

import (
	"context"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	consul "github.com/kitex-contrib/registry-consul"
	"time"
	"yima_user_app/kitex_gen/yima/news"
	"yima_user_app/kitex_gen/yima/news/yimanews"
)

var (
	cli yimanews.Client
)

func GetYimaNewsClient() yimanews.Client {
	if cli != nil {
		return cli
	}
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cli = yimanews.MustNewClient("yima.demo.news", client.WithResolver(r))
	return cli
}

func GetAllNewsByUid(ctx context.Context, id string) ([]*news.SciNews, error) {
	defer func() {
		if err := recover(); err != nil {
			logger.CtxErrorf(ctx, "panic!! err:%v", err)
			return
		}
	}()
	// 微服务之间不应该存在强依赖关系，因此到调用时再去寻找服务，而不是像数据库写到初始化里
	resp, err := GetYimaNewsClient().GetAllNewsByUserId(ctx, &news.GetAllNewsByUserIdRequest{UserId: id}, callopt.WithConnectTimeout(5*time.Second), callopt.WithRPCTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return resp.NewsList_, nil
	}
	return nil, nil

}
