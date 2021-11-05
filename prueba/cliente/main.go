package main

import (
	"context"
	"fmt"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
)

func main(){
	conn, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	serviceCLient := pb.NewHelloServiceClient(conn)
	res, err := serviceCLient.SayHello(context.Background(), &pb.HelloRequest{Name: "Anacletus"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Message)
}