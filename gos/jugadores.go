package main

import (
	"context"
	"log"
	// "math/rand"
	// "net"
	// "sync"
	"time"
	pb "github.com/MrAnacletus/Lab2-Distribuidos/tree/main/protos"

	"google.golang.org/grpc"
)

type Jugada struct {
	Id int32
	Jugador int32
	Carta int32
}


func EnviarJugada(id int32, jugador int32, carta int32) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Hubo un error al conectarse %v", err)
	}
	defer conn.Close()
	c := pb.NewLiderClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	jugada := &pb.Jugada{
		Id: id,
		Jugador: jugador,
		Carta: carta}
	r, err := c.EnviarJugada(ctx, jugada)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
