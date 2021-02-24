package main

import (
	"io"
	"google.golang.org/grpc"
	"gRPCstream/proto/serverStreamService"
	"fmt"
	"context"
)

func main() {
	conn, err := grpc.Dial("192.168.31.87:7160", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := serverStreamService.NewServerStreamClient(conn)
	request := &serverStreamService.Request{Data: "It is an download request"}
	stream, err := client.DataDownload(context.Background(), request)
	if err != nil {
		fmt.Println("Donwloading stream failed from server", err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Downloading stream done succesfully")
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(response)
	}
}