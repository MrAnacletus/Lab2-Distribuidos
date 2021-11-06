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

func EnviarJugada2(J Jugada)(Jugada){
	//Se establece la conexión con el servidor
	conn, err := grpc.Dial("10.6.40.218:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor Lider: %v", err)
	}
	defer conn.Close()

	//Se crea un cliente para la comunicación con el servidor
	serviceCLient := pb.NewLiderServiceClient(conn)
	//Se envia la Jugada al servidor
	res, err := serviceCLient.SendJugada2(context.Background(), &pb.Jugada{ID: J.ID, Jugada: J.jugada})
	if err != nil {
		log.Fatalf("Error al enviar la Jugada: %v", err)
	}

	return Jugada{ID: res.GetID(), jugada: res.GetEstado()}
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
	Team int32
}

var ListaJugadores [16]Bots

var Vivos int32



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
				Vivos -= 1
			}
			if Vivos == 1 {
				fmt.Println("Tenemos un ganador la conchadesumadre")
				return false
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
					Vivos -= 1
				}
				if Vivos == 1 {
					fmt.Println("Tenemos un ganador la conchadesumadre")
					return false
				}
			}
		}
	}
	// Pedir los muertos
	return true
}

func juego2()(bool){
	var suma1 int32 = 0
	var suma2 int32 = 0
	//Dividir los equipos
	//Contar los vivos
	if Vivos % 2 != 0 {
		Elegido := rand.Intn(int(Vivos))
		for i := 0; i < 16; i++ {
			if ListaJugadores[i].Estado == 1 {
				if Elegido == 0 {
					ListaJugadores[i].Estado = 0
				}
				Elegido -= 1
			}
		}
		Vivos -= 1
	}
	for i := 0 ;i < 16; i++ {
		if ListaJugadores[i].Estado == 1 {
			team := rand.Intn(1) + 1
			if team == 1 {
				if suma1 == Vivos/2 {
					team = 2
				}else{
					suma1 += 1
				}
			}else{
				if suma2 == Vivos/2 {
					team = 1
				}else{
					suma2 += 1
				}
			}
			ListaJugadores[i].Team = int32(team)
		}
	}
	suma1 = 0
	suma2 = 0
	//Contar las sumas
	for i := 0; i < 16; i++ {
		if ListaJugadores[i].Estado == 1 {
			if ListaJugadores[i].Team == 1 {
				if ListaJugadores[i].ID == 1 {
					fmt.Println("Es tu turno, usted pertenece al equipo 1")
					fmt.Println("Selecciona una jugada")
					fmt.Println("Elija un numero entre 1 y 4, la suma de su equipo debera tener la misma paridad que el numero elegido por el lider")
					var numero1 int
					fmt.Scanln(&numero1)
					suma1 += int32(numero1)

				}else{
					numero2 := rand.Intn(4) + 1
					suma1 += int32(numero2)
				}
			}else{
				if ListaJugadores[i].ID == 1 {
					fmt.Println("Es tu turno, usted pertenece al equipo 2")
					fmt.Println("Selecciona una jugada")
					fmt.Println("Elija un numero entre 1 y 4, la suma de su equipo debera tener la misma paridad que el numero elegido por el lider")
					var numero1 int
					fmt.Scanln(&numero1)
					suma2 += int32(numero1)

				}else{
					numero2 := rand.Intn(4) + 1
					suma2 += int32(numero2)
				}
			}
		}
	}

	//Enviar sumas al Lider
	var resultado = EnviarJugada2(Jugada{ID: suma1, jugada: suma2})
	//Recorrer los jugadores
	for i := 0; i < 16; i++{
		if ListaJugadores[i].Estado == 1{
			if ListaJugadores[i].Team == 1 {
				if resultado.ID == 1 {
					fmt.Println("El jugador " + fmt.Sprint(ListaJugadores[i].ID) + " sigue vivo")
				}else{
					fmt.Println("El jugador " + fmt.Sprint(ListaJugadores[i].ID) + " ha muerto")
					ListaJugadores[i].Estado = 0
					Vivos -= 1
				}
			}else{
				if resultado.jugada == 1 {
					fmt.Println("El jugador " + fmt.Sprint(ListaJugadores[i].ID) + " sigue vivo")
				}else{
					fmt.Println("El jugador " + fmt.Sprint(ListaJugadores[i].ID) + " ha muerto")
					ListaJugadores[i].Estado = 0
					Vivos -= 1
				}
			}
		}
	}
	if Vivos == 1 {
		fmt.Println("Tenemos un ganador la conchadesumadre")
		return false
	}
	return true
}

func main(){
	for i := 0; i < 16; i++ {
		ListaJugadores[i].ID = int32(i + 1)
		ListaJugadores[i].Estado = 1
	}

	Vivos = 16

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
				fmt.Println("El juego ha terminado, y tenemos un ganador!!!!!!!!!")
				for i := 0; i < 16; i++ {
					if ListaJugadores[i].Estado == 1 {
						fmt.Println("El ganador es el jugador " + fmt.Sprint(ListaJugadores[i].ID))
					}
				}
				break
			}
			if !juego2(){
				fmt.Println("El juego ha terminado, y tenemos un ganador!!!!!!!!!")
				for i := 0; i < 16; i++ {
					if ListaJugadores[i].Estado == 1 {
						fmt.Println("El ganador es el jugador " + fmt.Sprint(ListaJugadores[i].ID))
					}
				}
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