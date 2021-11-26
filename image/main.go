package main

import (
	"image.oscto.icu/logic"
	pb "image.oscto.icu/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "image.oscto.icu"
	version = "v0.0.1"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	_ = pb.RegisterImageHandler(srv.Server(), logic.Register())

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
