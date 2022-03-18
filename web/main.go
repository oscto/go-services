package main

import (
	"fmt"
	"net/http"
	"runtime"

	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"

	"log"

	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

var (
	service = "http"
	version = "v0.0.1"
	address = ":8090"
)

func main() {

	runtime.GOMAXPROCS(10)
	go func() {
		fmt.Println("hello world")
	}()

	go func() {
		for {
			fmt.Println("aaaa")
			//	time.Sleep(time.Duration(1) * time.Second)
		}
	}()

	select {}

	srv := httpServer.NewServer(
		server.Name(service),
		server.Address(address),
	)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/", func(ctx *gin.Context) {
		Hello()
		ctx.JSON(http.StatusOK, "Hello World")
	})

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(registry.NewRegistry()),
	)
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func Hello() {
	// service := micro.NewService()
	// client := ktt.NewBrokerKttService("broker-ktt", service.Client())
	// call, err := client.Call(context.Background(), &ktt.CallRequest{Name: "Channing"})

	// fmt.Println(fmt.Sprintf("call: %v,err: %v", call, err))
}
