package main

import (
	"flag"
	"log"
	"context"
	"time"
	"google.golang.org/grpc"
	api "github.com/YAWAL/gRPCtrying/api"

	"io/ioutil"
)

func main() {
	backend := flag.String("backend", "172.17.0.3:50111", "port for connection to download server")
	flag.Parse()

	cntx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(cntx, *backend, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		log.Fatalf("Error has occured %v:", err)
	}
	client := api.NewSearchPicClient(conn)
	name := &api.Name{Name: "gopher"}
	picture, err := client.Search(context.Background(), name)
	if err != nil {
		log.Fatalf("Error has occured %v:", err)
	}
	if err = ioutil.WriteFile("gopher1.png", picture.Pic, 0666); err != nil {
		log.Fatalf("Error has occured %v:", err)
		return
	}
	log.Println("File has been created")
}
