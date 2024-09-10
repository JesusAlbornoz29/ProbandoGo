package main

import "math";
import "fmt"

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}


func details (c circle){
	fmt.Println(c)
	fmt.Println("Area: ", c.area())
	fmt.Println("Perimeter: ", c.perim())
}

// ejemplo 2

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func details2(g geometry){
	fmt.Println(g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimeter: ", g.perim())
}

// Ejemplo 3 
func newCircle(values float64) circle {
	return circle{radius: values}
}

// ejemplo 4
const(
	rectType = "RECT"
	circleType = "CIRCLE"
)

func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case rectType:
		return rect{width: values[0], height: values[1]}
	case circleType:
		return circle{radius: values[0]}
	}
	return nil
}


// Ejemplo de interfaz vacia
type ListaHeterogenea struct {
	data []interface{}
}


func main() {
	c := circle {radius: 5}
	details(c)

	// ejemplo 2
	j := circle{radius: 2}
	fmt.Println(j.area())
	fmt.Println(j.perim())

	// ejemplo 3
	f := newCircle(2)
	fmt.Println(f.area())
	fmt.Println(f.perim())

	// ejemplo 4
	u := newGeometry(rectType, 2, 3)
	fmt.Println(u.area())
	fmt.Println(u.perim())
	i := newGeometry(circleType, 2)
	fmt.Println(i.area())
	fmt.Println(i.perim())


}

