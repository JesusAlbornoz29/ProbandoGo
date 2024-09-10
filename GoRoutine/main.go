package main 

/*
Calcular precio
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento. Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda, y para optimizar la velocidad, requieren que el cálculo de la sumatoria se realice en paralelo mediante tres goroutines.
Se requieren tres estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.
Y se requieren tres funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio por cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio por media hora trabajada. Si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.
Las tres se deben ejecutar concurrentemente y, al final, se debe mostrar por pantalla el monto final (sumando el total de las tres).
*/

import (
	"fmt"
	"sync"
)

type Producto struct {
	nombre string
	precio float64
	cantidad int
}

type Servicio struct {
	nombre string
	precio float64
	minutosTrabajados int
}

type Mantenimiento struct {
	nombre string
	precio float64
}


func main() {
	productos := []Producto{
		Producto{nombre: "Producto 1", precio: 100, cantidad: 10},
		Producto{nombre: "Producto 2", precio: 200, cantidad: 20},
		Producto{nombre: "Producto 3", precio: 300, cantidad: 30},
	}
	servicios := []Servicio{
		Servicio{nombre: "Servicio 1", precio: 100, minutosTrabajados: 10},
		Servicio{nombre: "Servicio 2", precio: 200, minutosTrabajados: 20},
		Servicio{nombre: "Servicio 3", precio: 300, minutosTrabajados: 30},
	}
	mantenimientos := []Mantenimiento{
		Mantenimiento{nombre: "Mantenimiento 1", precio: 100},
		Mantenimiento{nombre: "Mantenimiento 2", precio: 200},
		Mantenimiento{nombre: "Mantenimiento 3", precio: 300},
	}

	var wg sync.WaitGroup
	wg.Add(3)
	var total float64
	go func() {
		total += sumarProductos(productos)
		wg.Done()
	}()
	go func() {
		total += sumarServicios(servicios)
		wg.Done()
	}()
	go func() {
		total += sumarMantenimientos(mantenimientos)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("El monto total es:", total)
}

func sumarProductos(productos []Producto) float64 {
	var total float64
	for _, producto := range productos {
		total += producto.precio * float64(producto.cantidad)
	}
	return total
}

func sumarServicios(servicios []Servicio) float64 {
	var total float64
	for _, servicio := range servicios {
		total += servicio.precio * float64(servicio.minutosTrabajados) / 30
	}
	return total
}

func sumarMantenimientos(mantenimientos []Mantenimiento) float64 {
	var total float64
	for _, mantenimiento := range mantenimientos {
		total += mantenimiento.precio
	}
	return total
}

// Output

// El monto total es: 2100
// Program exited.
