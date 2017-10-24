package main

import (
	"flag"
	"net"
	"fmt"
	"log"
	"google.golang.org/grpc"
	api "github.com/YAWAL/gRPCtrying/api"
	"golang.org/x/net/context"
	"io/ioutil"
	"strings"
	"path/filepath"
	"time"
)

const (
	target            = "download:50112"
	searchServicePort = 50111
	resourceFolder    = "/home/vya/Pictures"
)

func main() {

	port := flag.Int("p", searchServicePort, "port to listen to")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Error has occured %d: %v", *port, err)
	}
	srv := grpc.NewServer()
	api.RegisterSearchPicServer(srv, server{})
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("Error has occured: %v", err)
	}
}

type server struct{}

func (server) Search(ctx context.Context, name *api.Name) (*api.Picture, error) {
	pics, err := ioutil.ReadDir(resourceFolder)
	if err != nil {
		log.Printf("Error has occured %v", err)
		return new(api.Picture), err
	}

	var path string

	for _, p := range pics {
		if !p.IsDir() {
			if strings.Contains(p.Name(), name.Name){
				path = filepath.Join(resourceFolder, p.Name())
				break
			}
		}
	}
fmt.Println(path)
	var conn *grpc.ClientConn

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err = grpc.DialContext(ctx, target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error has occured: %v", err)
	}

	client := api.NewDownloadPicClient(conn)
	id := &api.Id{Id: path}
	pictures, err := client.Download(context.Background(), id)
	if err != nil {
		log.Printf("Error has occured: %v", err)
		return new(api.Picture), err
	}
	return pictures, nil
}
