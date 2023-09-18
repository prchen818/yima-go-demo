package handler

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"yima_user_app/kitex_gen/yima/news"
	"yima_user_app/rpc"
	"yima_user_app/service"
)

type GetAllNewsByUserResponse struct {
	Id       string          `json:"uid"`
	Nickname string          `json:"nickname"`
	Phone    string          `json:"phone"`
	NewsList []*news.SciNews `json:"news_list"`
}

func GetAllNewsByUser(ctx context.Context, c *app.RequestContext) {
	uid := c.Query("uid")

	newss, err := rpc.GetAllNewsByUid(ctx, uid)
	if err != nil {
		logger.CtxErrorf(ctx, "[GetAllNewsByUser]err: %v", err)
		return
	}
	sysUser, err := service.NewUserService().GetUserInfoByUid(ctx, uid)
	if err != nil {
		logger.CtxErrorf(ctx, "[GetAllNewsByUser]err: %v", err)
		return
	}
	resp := &GetAllNewsByUserResponse{
		Id:       sysUser.ID,
		Nickname: sysUser.Nickname,
		Phone:    sysUser.Phone,
		NewsList: newss,
	}
	c.JSON(consts.StatusOK, resp)
	return
}
