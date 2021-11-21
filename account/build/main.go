package main

import (
	"log"

	"account.oscto.icu/logic"
	account "account.oscto.icu/proto/handler"
	_ "github.com/asim/go-micro/plugins/registry/etcd/v4"
	"go-micro.dev/v4"
)

var (
	serviceName = "account.oscto.icu"
	etcdAddress = "192.168.0.168:2379"
)

func main() {

	service := micro.NewService(
		micro.Name(serviceName),
		//micro.Registry(etcd.NewRegistry(
		//	registry.Addrs(etcdAddress),
		//)),
	)
	service.Init()

	if err := account.RegisterAccountHandler(service.Server(), logic.Register()); err != nil {
		return
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
