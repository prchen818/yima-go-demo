package main

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"yima_news_app/dao"
	news "yima_news_app/kitex_gen/yima/news/yimanews"
)

func main() {

	dao.DatabaseInit()
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	svr := news.NewServer(
		new(YimaNewsImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: "yima.demo.news",
			Addr:        utils.NewNetAddr("tcp", "127.0.0.1:8888"),
			Weight:      1,
		}),
		// 绑定到内网ip以便consul服务发现
		server.WithServiceAddr(utils.NewNetAddr("tcp", "127.0.0.1:8888")))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
