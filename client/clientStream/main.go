package main

import (
	"gRPCstream/proto/clientStreamService"
	"google.golang.org/grpc"
	"context"
	"math/rand"
	"time"
	"fmt"
)

func main() {
	conn, err := grpc.Dial("192.168.31.87:7160", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := clientStreamService.NewClientStreamClient(conn)

	stream, err := client.DataUpload(context.Background())	

	for i := 0; i <100000; i++ {
		rand.Seed(time.Now().Unix())
		data := fmt.Sprintf("%v: %v",i, rand.Int())
		err := stream.Send(&clientStreamService.Request{Data: data})
		if err != nil {
			fmt.Println("Error sending Request data: ", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("Error recieving response: ", err)
	}
	fmt.Println(response.GetData())
}