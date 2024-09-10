package main

import (
	"fmt"
	"net/http" // importamos el paquete http para poder crear un servidor web
)

func holaHandler (w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Hola \n")
}

func main() {
	http.HandleFunc("/hola", holaHandler)
	http.ListenAndServe(":8080", nil)
}