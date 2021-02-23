package main

import (
	"net"
	"fmt"
	"io"
	"gRPCstream/proto/clientStreamService"
	"google.golang.org/grpc"
)

type Server struct {}

func(s Server) DataUpload(stream clientStreamService.ClientStream_DataUploadServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&clientStreamService.Response{Data: "Data loaded succesffully"})
		}
		if err != nil {
			return err
		}
		fmt.Println(request.GetData())
	}
	return nil	
}

const(
	port = "7160"
)

func main() {

	svr := grpc.NewServer()

	clientStreamService.RegisterClientStreamServer(svr, &Server{})

	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server listening at %s port...\n", port)

	svr.Serve(conn)
}