package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
)

type Jugada struct {
	ID int32
	jugada int32
}

func EnviarPeticionJugar(){
	//Se establece la conexión con el servidor
	conn, err := grpc.Dial("10.6.40.218:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor Lider: %v", err)
	}
	defer conn.Close()

	//Se crea un cliente para la comunicación con el servidor
	serviceCLient := pb.NewLiderServiceClient(conn)
	//Se envia la petición de jugar al servidor
	res, err := serviceCLient.SayHello(context.Background(), &pb.HelloRequest{Name: "Preparandoce para iniciar el juego!"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Message)
}

func EnviarJugada(J Jugada)(bool){
	//Se establece la conexión con el servidor
	conn, err := grpc.Dial("10.6.40.218:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor Lider: %v", err)
	}
	defer conn.Close()

	//Se crea un cliente para la comunicación con el servidor
	serviceCLient := pb.NewLiderServiceClient(conn)
	//Se envia la Jugada al servidor
	res, err := serviceCLient.SendJugada(context.Background(), &pb.Jugada{ID: J.ID, Jugada: J.jugada})
	if err != nil {
		log.Fatalf("Error al enviar la Jugada: %v", err)
	}
	return res.GetEstado() == 1
}

func PedirPozo(){
	//Se establece la conexión con el servidor
	conn, err := grpc.Dial("10.6.40.218:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor Lider: %v", err)
	}
	defer conn.Close()

	//Se crea un cliente para la comunicación con el servidor
	serviceCLient := pb.NewLiderServiceClient(conn)
	//Se envia la petición de pozo al servidor
	res, err := serviceCLient.RequestPozo(context.Background(), &pb.RequestPozoActual{Pozo: 0})
	if err != nil {
		panic(err)
	}
	fmt.Println("El pozo actual es: " + fmt.Sprint(res.GetPozo()))

}

type Bots struct{
	ID int32
	Estado int32
}

var ListaJugadores [16]Bots



func juego1()(bool){
	var suma int32 = 0
	for i := 0; i < 4; i++ {
		if ListaJugadores[0].Estado == 1 {
			fmt.Println("Es tu turno")
			fmt.Println("Selecciona una jugada")
			fmt.Println("Elija un numero entre 1 y 10, recuerde que sus numeros deben sumar 21. Actualmente tienes" + fmt.Sprint(suma))
			var numero1 int
			fmt.Scanln(&numero1)
			suma += int32(numero1)
			var j = Jugada{ID: 1, jugada: int32(numero1)}
			if EnviarJugada(j){
				fmt.Println("El jugador 1 sigue vivo")
			}else{
				fmt.Println("El jugador 1 ha muerto")
				ListaJugadores[0].Estado = 0
			}
		}

		//Juego de los bots
		for i:= 1; i < 16; i++ {
			if ListaJugadores[i].Estado == 1 {
				numero2 := rand.Intn(10) + 1
				var j = Jugada{ID: ListaJugadores[i].ID, jugada: int32(numero2)}
				if EnviarJugada(j){
					fmt.Println("El jugador " + fmt.Sprint(ListaJugadores[i].ID) + " sigue vivo")
				}else{
					fmt.Println("El jugador " + fmt.Sprint(ListaJugadores[i].ID) + " ha muerto")
					ListaJugadores[i].Estado = 0
				}
			}
		}
	}
	// Pedir los muertos
	return true
}

func main(){
	for i := 0; i < 16; i++ {
		ListaJugadores[i].ID = int32(i + 1)
		ListaJugadores[i].Estado = 1
	}
	fmt.Println("Iniciando el cliente")
	fmt.Println("Bienvenido al Juego del Calamar")
	fmt.Println("Para comenzar a jugar, presiona enter")
	fmt.Scanln()
	EnviarPeticionJugar()
	fmt.Println("El juego ha comenzado")
	fmt.Println("Selecciona que vas a hacer:")
	fmt.Println("1. Jugar")
	fmt.Println("2. Ver pozo")
	var opcion int
	fmt.Scanln(&opcion)
	for{
		if opcion == 1 {
			if !juego1(){
				fmt.Println("El juego ha terminado")
				break
			}
		}
		if opcion == 2 {
			PedirPozo()
		}
		fmt.Println("Selecciona que vas a hacer:")
		fmt.Println("1. Jugar")
		fmt.Println("2. Ver pozo")
		fmt.Scanln(&opcion)
	}
	
}