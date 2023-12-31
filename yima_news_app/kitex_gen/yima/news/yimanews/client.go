// Code generated by Kitex v0.7.1. DO NOT EDIT.

package yimanews

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	news "yima_news_app/kitex_gen/yima/news"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetNewsByTitle(ctx context.Context, request *news.GetNewsByTitleRequest, callOptions ...callopt.Option) (r *news.GetNewsByTitleResponse, err error)
	AddNews(ctx context.Context, request *news.AddNewsRequest, callOptions ...callopt.Option) (r *news.AddNewsResponse, err error)
	GetAllNewsByUserId(ctx context.Context, request *news.GetAllNewsByUserIdRequest, callOptions ...callopt.Option) (r *news.GetAllNewsByUserIdResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kYimaNewsClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kYimaNewsClient struct {
	*kClient
}

func (p *kYimaNewsClient) GetNewsByTitle(ctx context.Context, request *news.GetNewsByTitleRequest, callOptions ...callopt.Option) (r *news.GetNewsByTitleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetNewsByTitle(ctx, request)
}

func (p *kYimaNewsClient) AddNews(ctx context.Context, request *news.AddNewsRequest, callOptions ...callopt.Option) (r *news.AddNewsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddNews(ctx, request)
}

func (p *kYimaNewsClient) GetAllNewsByUserId(ctx context.Context, request *news.GetAllNewsByUserIdRequest, callOptions ...callopt.Option) (r *news.GetAllNewsByUserIdResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllNewsByUserId(ctx, request)
}
