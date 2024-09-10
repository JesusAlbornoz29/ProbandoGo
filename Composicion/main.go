package main 
// Composicion en Go : Metodo para escribir segmentos de codigo reutilizables. 
// Estructuras se componen de otras estructuras para usar sus funcionalidades.

import "fmt"


type Autor struct {
	Nombre string
	Apellido string
}

func (a Autor) nombrecompleto() string {
	return fmt.Sprintf("%s %s", a.Nombre, a.Apellido)
}

// Embeber? : Se puede embeber una estructura en otra estructura.
type Libro struct {
	Titulo string
	Descripcion string
	Autor Autor
}

func (l Libro) informacion() {
	fmt.Println("Titulo: ", l.Titulo)
	fmt.Println("Descripcion: ", l.Descripcion)
	fmt.Println("Autor: ", l.Autor.nombrecompleto())
}

func main() {
	// instanciamos un autor 
	autor := Autor{
		Nombre: "Juan",
		Apellido: "Lopez",
	}

	libro := Libro{
		Titulo: "Go para principiantes",
		Descripcion: "Un libro para aprender Go",
		Autor: autor,
	}

	libro.informacion()

}