// Code generated by Kitex v0.7.1. DO NOT EDIT.
package yimanews

import (
	server "github.com/cloudwego/kitex/server"
	news "yima_news_app/kitex_gen/yima/news"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler news.YimaNews, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
