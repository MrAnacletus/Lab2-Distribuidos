package main

import (
	"context"
	"fmt"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto/github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto/github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
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
	players := make([]Jugador,16)
	for i := 0; i < 16; i++ {
		players[i].Name = "Jugador " + string(i+1)
		players[i].Id = i
		players[i].State = 1
	}

	//Enviar los jugadores
	res3, err := serviceCLient.GetJugadores(context.Background(), &pb.Jugadores{jugadores: players})
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