package backend

import (
	"flag"
	"net"
	"fmt"
)

func main(){

	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil{

	}
}