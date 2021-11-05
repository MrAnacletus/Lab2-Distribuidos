package main

import (
	"context"
	"fmt"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto/github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
)

type Jugador struct {
	Name string
	Id int
	State int
}

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

	//Crear los jugadores
	jugadores := make([]Jugador,16)
	for i := 0; i < 16; i++ {
		jugadores[i].Name = "Jugador " + string(i+1)
		jugadores[i].Id = i
		jugadores[i].State = 1
	}

	res2, err2 := serviceCLient.SayHelloAgain(context.Background(), &pb.HelloRequest{Name: "Listos para iniciar el juego!"})
	if err2 != nil {
		panic(err)
	}
	fmt.Println(res2.Message)
}