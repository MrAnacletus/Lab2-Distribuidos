package main

import (
	"context"
	"fmt"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto/github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto/github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
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
	jugadas := [16]int{}
	for i := 0; i < 16; i++ {
		jugadas[i] = i
	}

	//Enviar los jugadores
	res3, err := serviceCLient.GetJugadas(context.Background(), &pb.Jugadas{ID: 1,Jugadas: jugadas})
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