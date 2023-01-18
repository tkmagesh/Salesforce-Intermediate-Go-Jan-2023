package main

import (
	"fmt"
)

type Product struct {
	Id   int
	Name string
	Cost float32
}

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func (emp Employee) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %.2f", emp.Id, emp.Name, emp.Salary)
}

func (emp *Employee) AwardBonus(bonus float32) {
	emp.Salary += bonus
}

func main() {
	nos1 := [5]int{3, 1, 4, 2, 5}
	nos2 := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos1 == nos2)

	p1 := Product{100, "Pen", 10}
	p2 := Product{100, "Pen", 10}
	fmt.Println(p1 == p2)

	pen := Product{100, "Pen", 10}
	fmt.Println("Before Applying Discount, pen -", pen)
	ApplyDiscount(&pen, 10)
	fmt.Println("After Applying Discount, pen -", pen)

	var no int
	no = 100

	var noPtr *int
	noPtr = &no // address(pointer) from the value
	fmt.Println(no, noPtr)

	//dereferencing
	//value from the address (pointer)
	x := *noPtr
	fmt.Println(x)

	fmt.Println(no == *(&no))

	fmt.Println("Before incrementing, no =", no)
	increment(&no)
	fmt.Println("After incrementing, no =", no)

	pencil := &Product{200, "Pencil", 10}
	fmt.Println((*pencil).Id)
	fmt.Println(pencil.Id)

	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	// fmt.Println(Format(emp))
	fmt.Println(emp.Format())
	fmt.Println("After awarding bonus(2000)")
	// (&emp).AwardBonus(2000) // DONOT have to do this
	emp.AwardBonus(2000)
	fmt.Println(emp.Format())
}

func increment(val *int) {
	fmt.Println("Address of val =", val)
	*val++
}

func ApplyDiscount(p *Product, discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
