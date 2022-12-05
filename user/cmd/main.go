package main

import (
	"context"
	"git.vonechain.com/vone-bfs/user/handler"
	pb "git.vonechain.com/vone-bfs/user/proto"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
	"sync"
	"time"

	ot "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"

	"github.com/go-micro/cli/debug/trace/jaeger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service = "user"
	version = "latest"
)

func main() {
	// Create tracer
	tracer, closer, err := jaeger.NewTracer(
		jaeger.Name(service),
		jaeger.FromEnv(true),
		jaeger.GlobalTracer(true),
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer closer.Close()

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.BeforeStart(func() error {
			logger.Infof("Starting service %s", service)
			return nil
		}),
		micro.BeforeStop(func() error {
			logger.Infof("Shutting down service %s", service)
			cancel()
			return nil
		}),
		micro.AfterStop(func() error {
			wg.Wait()
			return nil
		}),
		micro.WrapCall(ot.NewCallWrapper(tracer)),
		micro.WrapClient(ot.NewClientWrapper(tracer)),
		micro.WrapHandler(ot.NewHandlerWrapper(tracer)),
		micro.WrapSubscriber(ot.NewSubscriberWrapper(tracer)),
		// 服务发现注册
		micro.RegisterTTL(10*time.Second),
		micro.RegisterInterval(30*time.Second),
		micro.Registry(etcd.NewRegistry(registry.Addrs("43.138.199.52:2379"))),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Server().Init(
		server.Wait(&wg),
	)

	ctx = server.NewContext(ctx, srv.Server())

	// Register handler
	if err := pb.RegisterUserHandler(srv.Server(), new(handler.User)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
