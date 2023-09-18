// Code generated by Kitex v0.7.1. DO NOT EDIT.

package yimanews

import (
			"context"
				client "github.com/cloudwego/kitex/client"
				kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
				news "yima_user_app/kitex_gen/yima/news"
)

func serviceInfo() *kitex.ServiceInfo {
	return yimaNewsServiceInfo
 }

var yimaNewsServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "YimaNews"
	handlerType := (*news.YimaNews)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetNewsByTitle":
			kitex.NewMethodInfo(getNewsByTitleHandler, newYimaNewsGetNewsByTitleArgs, newYimaNewsGetNewsByTitleResult, false),
		"AddNews":
			kitex.NewMethodInfo(addNewsHandler, newYimaNewsAddNewsArgs, newYimaNewsAddNewsResult, false),
		"GetAllNewsByUserId":
			kitex.NewMethodInfo(getAllNewsByUserIdHandler, newYimaNewsGetAllNewsByUserIdArgs, newYimaNewsGetAllNewsByUserIdResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":	 "news",
		"ServiceFilePath": "idl/yima_news_app/base.thrift",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName: 	 serviceName,
		HandlerType: 	 handlerType,
		Methods:     	 methods,
		PayloadCodec:  	 kitex.Thrift,
		KiteXGenVersion: "v0.7.1",
		Extra:           extra,
	}
	return svcInfo
}



func getNewsByTitleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*news.YimaNewsGetNewsByTitleArgs)
	realResult := result.(*news.YimaNewsGetNewsByTitleResult)
	success, err := handler.(news.YimaNews).GetNewsByTitle(ctx, realArg.Request)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newYimaNewsGetNewsByTitleArgs() interface{} {
	return news.NewYimaNewsGetNewsByTitleArgs()
}

func newYimaNewsGetNewsByTitleResult() interface{} {
	return news.NewYimaNewsGetNewsByTitleResult()
}


func addNewsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*news.YimaNewsAddNewsArgs)
	realResult := result.(*news.YimaNewsAddNewsResult)
	success, err := handler.(news.YimaNews).AddNews(ctx, realArg.Request)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newYimaNewsAddNewsArgs() interface{} {
	return news.NewYimaNewsAddNewsArgs()
}

func newYimaNewsAddNewsResult() interface{} {
	return news.NewYimaNewsAddNewsResult()
}


func getAllNewsByUserIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error { 
	realArg := arg.(*news.YimaNewsGetAllNewsByUserIdArgs)
	realResult := result.(*news.YimaNewsGetAllNewsByUserIdResult)
	success, err := handler.(news.YimaNews).GetAllNewsByUserId(ctx, realArg.Request)
	if err != nil {
	return err
	}
	realResult.Success = success
	return nil 
}
func newYimaNewsGetAllNewsByUserIdArgs() interface{} {
	return news.NewYimaNewsGetAllNewsByUserIdArgs()
}

func newYimaNewsGetAllNewsByUserIdResult() interface{} {
	return news.NewYimaNewsGetAllNewsByUserIdResult()
}


type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}


func (p *kClient) GetNewsByTitle(ctx context.Context , request *news.GetNewsByTitleRequest) (r *news.GetNewsByTitleResponse, err error) {
	var _args news.YimaNewsGetNewsByTitleArgs
	_args.Request = request
	var _result news.YimaNewsGetNewsByTitleResult
	if err = p.c.Call(ctx, "GetNewsByTitle", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddNews(ctx context.Context , request *news.AddNewsRequest) (r *news.AddNewsResponse, err error) {
	var _args news.YimaNewsAddNewsArgs
	_args.Request = request
	var _result news.YimaNewsAddNewsResult
	if err = p.c.Call(ctx, "AddNews", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAllNewsByUserId(ctx context.Context , request *news.GetAllNewsByUserIdRequest) (r *news.GetAllNewsByUserIdResponse, err error) {
	var _args news.YimaNewsGetAllNewsByUserIdArgs
	_args.Request = request
	var _result news.YimaNewsGetAllNewsByUserIdResult
	if err = p.c.Call(ctx, "GetAllNewsByUserId", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

