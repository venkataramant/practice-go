package main

import (
	"context"
	"fmt"
	"log"
	"net"

	rides "github.com/venkataramant/grpc/pb/rides"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := ":8282"
	nLis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("cannot listen on %s %s", addr, err)
	}
	srv := grpc.NewServer()
	// var loc rides.Location
	var u Rides
	rides.RegisterRidesServer(srv, &u)
	reflection.Register(srv)
	log.Printf("Info server ready on %s", addr)
	if err := srv.Serve(nLis); err != nil {
		log.Fatalf("error: can't server - %s", err)
	}

}

type Rides struct {
	rides.UnimplementedRidesServer
}

func (r *Rides) Start(ctx context.Context, request *rides.StartRequest) (*rides.StartResponse, error) {
	// TODO validate Req
	fmt.Printf("inside Start %+v", request)
	resp := rides.StartResponse{
		Id: request.Id,
	}
	return &resp, nil
}
