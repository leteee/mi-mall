package main

import (
	"context"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"time"

	pb "captcha-client/proto/captcha"

	log "go-micro.dev/v4/logger"
)

var (
	service = "captcha"
	version = "latest"
)

func main() {
	//配置consul
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	// Create service
	srv := micro.NewService(
		micro.Registry(consulReg),
	)
	srv.Init()

	// Create client
	c := pb.NewCaptchaService(service, srv.Client())

	for {
		// Call service
		rsp, err := c.MakeCaptcha(context.Background(), &pb.MakeCaptchaRequest{
			Height: 20,
			Width:  60,
			Length: 2,
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Info(rsp)

		time.Sleep(1 * time.Second)
	}
}
