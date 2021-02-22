package main

import (
	"fmt"
	"net/http"
)

func  healthCheckServer(w http.ResponseWriter, r * http.Request){
	fmt.Fprintf(w, "Conexão estabelecida com sucesso.")
}


func serverConfig(){
	http.HandleFunc("/", healthCheckServer );

	err := http.ListenAndServe(":1337", nil); // DefaultServerMux
	if err != nil{
		fmt.Println(err);
	}
}

func main(){
	serverConfig();


}