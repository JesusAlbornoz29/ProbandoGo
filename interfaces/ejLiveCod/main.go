package main

import "fmt"

type Compania struct {
	Nombre string
	Puesto string
}

type Empleado struct {
	Nombre   string
	Apellido string
	Compania Compania
}

func (e Empleado) Informacion() {
	fmt.Printf("Empleado: %s %s \n Compania: %s \n", e.Nombre, e.Apellido, e.Compania.Puesto, e.Compania.Nombre)
}

func (c *Compania) CambiarPuesto(nuevoPueto string) {
	c.Puesto = nuevoPueto
}

func main() {
	empleado := Empleado {
		Nombre: "Juan",
		Apellido: "Perez",
		Compania: Compania {
			Nombre: "Google",
			Puesto: "Desarrollador",
		},
	}


	empleado.Informacion()
	empleado.Compania.CambiarPuesto("Gerente")
	empleado.Informacion()

}
