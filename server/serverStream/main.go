package main
// protoc --go_out=plugins=grpc:. ./proto/serverStreamService/serverStreamService.proto

import (
	"net"
	"google.golang.org/grpc"
	"gRPCstream/proto/serverStreamService"
	"math/rand"
	"fmt"
	"time"
)

type Server struct {}

func (s Server) DataDownload(req *serverStreamService.Request, stream serverStreamService.ServerStream_DataDownloadServer) error {
	fmt.Println("Got a request: ", req)
	for i := 0; i <100000; i++ {
		rand.Seed(time.Now().Unix())
		data := fmt.Sprintf("%v: %v",i, rand.Int())
		err := stream.Send(&serverStreamService.Response{Data: data})
		if err != nil {
			fmt.Println("Error sending Response data: ", err)
		}
	}
	return nil
}

const(
	port = "7160"
)

func main() {

	svr := grpc.NewServer()

	serverStreamService.RegisterServerStreamServer(svr, &Server{})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server listening at %s port...\n", port)

	svr.Serve(listener)
}