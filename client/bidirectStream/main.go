package main

import (
	"gRPCstream/proto/bidirectStreamService"
	"google.golang.org/grpc"
	"context"
	"math/rand"
	"time"
	"fmt"
	"io"
)

func main() {
	conn, err := grpc.Dial("192.168.31.87:7160", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := bidirectStreamService.NewBidirectStreamClient(conn)
	stream, err:= client.DataExchange(context.Background())
	if err != nil {
		panic(err)
	}

	for i := 0; i <100000; i++ {
		rand.Seed(time.Now().Unix())
		data := fmt.Sprintf("Data from client #%v: %v",i, rand.Int())
		err := stream.Send(&bidirectStreamService.Request{Data: data})
		if err != nil {
			fmt.Println("Error sending Request data: ", err)
		}

		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(response)
	}
}