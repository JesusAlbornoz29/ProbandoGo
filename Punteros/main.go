package main

import "fmt"

var p *int // p es un puntero a un entero 
var p1 = new(int) // otra forma de crear un puntero
var v int = 50




func mostrar(){
	
// Hacemos que el puntero p, almacene (apunte a) la direccion de memoria de la variable v.

p = &v
fmt.Printf("El puntero p apunta a la direccion de memoria: %v \n", p)
fmt.Printf("Al desreferenciar el puntero p obtengo el valor de la variable a la cual apunta, es decir, el valor de v: %d \n", *p)

}

func main() {
	mostrar()
}
