package main

import (
	"context"
	"fmt"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/proto"
)

func main(){
	conn, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	serviceCLient := pb.NewHelloServiceClient(conn)
	res, err := serviceCLient.SayHello(context.Background(), &pb.HelloRequest{Name: "Preparandoce para iniciar el juego!"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Message)

	//Crear los jugadas
	var jugadas [16]int32
	var i int32
	for i < 16 {
		jugadas[i] = i
		i++
	}
	data, err := proto.Marshal(jugadas)

	//Enviar los jugadores
	res3, err := serviceCLient.GetJugadas(context.Background(), &pb.Jugadas{ID: 1,Jugadas: data})
	if err != nil {
		panic(err)
	}
	fmt.Println(res3.Message)

	res2, err2 := serviceCLient.SayHelloAgain(context.Background(), &pb.HelloRequest{Name: "Listos para iniciar el juego!"})
	if err2 != nil {
		panic(err)
	}
	fmt.Println(res2.Message)
}