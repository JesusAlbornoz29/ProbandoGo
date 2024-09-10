package main 

import (
	"fmt"
	"os"
)

/*
Una empresa que se encarga de vender productos de limpieza necesita: 
1 Implementar una funcionalidad para guardar un archivo de texto 
con la información de productos comprados, separados por punto y coma (CSV).
2 Este archivo debe tener el ID del producto, precio y la cantidad.
3 Estos valores pueden ser hardcodeados o escritos en duro en una variable.

*/

type Producto struct {
	ID int
	Precio float64
	Cantidad int
}



func main() {
	productos := []Producto{
		Producto{ID: 1, Precio: 100, Cantidad: 10},
		Producto{ID: 2, Precio: 200, Cantidad: 20},
		Producto{ID: 3, Precio: 300, Cantidad: 30},
	}
	err := guardarProductos(productos)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Productos guardados en el archivo productos.txt")
}

func guardarProductos(productos []Producto) error {
	archivo, err := os.Create("productos.txt")
	if err != nil {
		return err
	}
	defer archivo.Close()
	for _, producto := range productos {
		_, err := fmt.Fprintf(archivo, "%d;%.2f;%d\n", producto.ID, producto.Precio, producto.Cantidad)
		if err != nil {
			return err
		}
	}
	return nil
}
