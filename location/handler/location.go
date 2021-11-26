package handler

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	log "go-micro.dev/v4/logger"

	"github.com/oschwald/geoip2-golang"
	"github.com/oscto/ky3k"
	pb "location.oscto.icu/proto"
)

type Location struct{}

func (e *Location) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {

	fmt.Println("client ip", req.ClientIp)
	location, err := e.ParseIP(req.ClientIp)
	if err != nil {
		return err
	}

	rsp.Location = ky3k.JsonToString(location)

	return nil
}

func (e *Location) ClientStream(ctx context.Context, stream pb.Location_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *Location) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Location_ServerStreamStream) error {
	log.Infof("Received Location.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *Location) BidiStream(ctx context.Context, stream pb.Location_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func (e *Location) ParseIP(ipAddr string) (*geoip2.City, error) {

	db, err := geoip2.Open("/data/services/GeoLite2-City.mmdb")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	ip := net.ParseIP(ipAddr)
	record, err := db.City(ip)
	if err != nil {
		return nil, err
	}

	return record, nil
}
