package main

// protoc --go_out=plugins=grpc:. ./proto/bidirectStreamService/bidirectStreamService.proto

import (
	"google.golang.org/grpc"
	"fmt"
	"io"
	"gRPCstream/proto/bidirectStreamService"
	"math/rand"
	"time"
	"net"
)

type Server struct {}

func (s Server) DataExchange (stream bidirectStreamService.BidirectStream_DataExchangeServer) error {
	for i := 0; i <100000; i++ {
		rand.Seed(time.Now().Unix())
		data := fmt.Sprintf("Data from Server #%v: %v",i, rand.Int())
		err := stream.Send(&bidirectStreamService.Response{Data: data})
		if err != nil {
			fmt.Println("Error sending Response data: ", err)
		}
		
		recvData, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println(recvData)
	}
	return nil
}

const(
	port = "7160"
)

func main() {

	svr := grpc.NewServer()

	bidirectStreamService.RegisterBidirectStreamServer(svr, &Server{})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server listening at %s port...\n", port)

	svr.Serve(listener)
}