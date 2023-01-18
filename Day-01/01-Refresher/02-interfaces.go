package main

import (
	"fmt"
	"math"
)

type Circle struct /* implements AreaFinder, PerimeterFinder, ShapeStatsFinder  */ {
	Name   string
	Radius float32
}

func (this Circle) Area() float32 {
	return math.Pi * this.Radius * this.Radius
}

func (this Circle) Perimeter() float32 {
	return 2 * math.Pi * this.Radius
}

func (this *Circle) ChangeName(newName string) {
	this.Name = newName
}

/* fmt.Stringer interface implementation */
func (this Circle) String() string {
	return fmt.Sprintf("Circle -> Radius:%v", this.Radius)
}

type Rectangle struct /* implements AreaFinder */ {
	Length float32
	Width  float32
}

func (this Rectangle) Area() float32 {
	return this.Length * this.Width
}

func (this Rectangle) Perimeter() float32 {
	return 2 * (this.Length + this.Width)
}

//utility functions
type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder /* anything that implements the Area() method */) {
	fmt.Println("Area :", x.Area())
}

type PerimeterFinder interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

/* new sprint */
type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShape(x ShapeStatsFinder) {
	PrintArea(x)      // x should implement AreaFinder interface
	PrintPerimeter(x) // x should implement PerimeterFinder interface
}

/*  */

type NameChanger interface {
	ChangeName(string)
}

/*
func UpdateName(x ?, newName string) {
	x.ChangeName(newName)
}
*/
func main() {
	c := Circle{Radius: 12, Name: "Small Circle"}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{Length: 10, Width: 12}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)

	// fmt.Println(c) // Circle -> Radius:12

	//changing the name using the 'Circle' type
	fmt.Println("changing the name using the 'Circle' type")
	c.ChangeName("Large Circle")
	fmt.Println(c.Name)

	// changing the name using the NameChanger interface
	fmt.Println("changing the name using the NameChanger interface")
	var nc NameChanger
	nc = &c
	nc.ChangeName("Large Circle")
	fmt.Println(c.Name)
}
