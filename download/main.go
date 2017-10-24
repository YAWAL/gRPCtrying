package main

import (
	"flag"
	"net"
	"fmt"
	"log"
	api "github.com/YAWAL/gRPCtrying/api"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"io/ioutil"
)
const (
	downloadServicePort = 50112
)

func main(){

	port := flag.Int("p", downloadServicePort, "port to listen to")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil{
		log.Fatalf("Error has occured %d: %v", *port, err)
	}
	srv := grpc.NewServer()
	api.RegisterDownloadPicServer(srv, server{})
	err = srv.Serve(lis)
	if err != nil{
		log.Fatalf("Error has occured: %v", err)
	}
}

type server struct {}

func (server) Download(ctx context.Context, id *api.Id) (*api.Picture, error) {
	defer ctx.Done()
	pic, err := ioutil.ReadFile(id.Id)
	if err != nil{
		log.Printf("Error has occured: %v", err)
		return &api.Picture{}, err
	}
	picture := &api.Picture{Pic:pic}
	return picture, nil
}
