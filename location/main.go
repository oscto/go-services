package main

import (
	"location.oscto.icu/handler"
	pb "location.oscto.icu/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "location.oscto.icu"
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
	pb.RegisterLocationHandler(srv.Server(), new(handler.Location))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
