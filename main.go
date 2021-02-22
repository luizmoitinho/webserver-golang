package main

import (
	"fmt"
	"net/http"
)

func main(){
	
	http.HandleFunc("/", func (w http.ResponseWriter, r * http.Request){
		fmt.Fprintf(w, "Bem-vindo")
	});

	err := http.ListenAndServe(":1337", nil); // DefaultServerMux
	if err != nil{
		fmt.Println(err);
	}


}