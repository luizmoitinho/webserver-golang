package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type Editora struct{
	Id int
	Nome string
	ISBN int
}
type Book struct {
	Id int
	Titulo string
	Autor string
	Editora Editora
}

var Books []Book = []Book{
	Book{
		Id:1,
		Titulo: "Poder do Hábito",
		Autor: "Charles Duhing",
	},
	Book{
		Id:2,
		Titulo: "Dom Casmurro",
		Autor: "Machado de Assis",
	},

}

func  healthCheckServer(w http.ResponseWriter, r * http.Request){
	fmt.Fprintf(w, "Conexão estabelecida com sucesso.")
}

func getAllBooks(w http.ResponseWriter, r * http.Request){
	encoder :=  json.NewEncoder(w)
	encoder.Encode(Books)
}

func serverConfig(){

	http.HandleFunc("/", healthCheckServer );
	http.HandleFunc("/books", getAllBooks );
	
	fmt.Println("Servidor está operando na porta: 1337");

	log.Fatal(http.ListenAndServe(":1337", nil)); // DefaultServerMux

}

func main(){
	serverConfig();


}