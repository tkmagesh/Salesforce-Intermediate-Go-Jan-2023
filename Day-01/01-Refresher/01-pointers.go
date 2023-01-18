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

func (this Employee) WhoAmI() {
	fmt.Println("I am an employee - ", this.Name)
}

func (this Employee) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %.2f", this.Id, this.Name, this.Salary)
}

func (this *Employee) AwardBonus(bonus float32) {
	this.Salary += bonus
}

func main() {

	/* Arrays are values */
	nos1 := [5]int{3, 1, 4, 2, 5}
	nos2 := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos1 == nos2)

	/* Structs are also values */
	p1 := Product{100, "Pen", 10}
	p2 := Product{100, "Pen", 10}
	fmt.Println(p1 == p2)

	/* Use reference (aka pointers) to change the state */
	pen := Product{100, "Pen", 10}
	fmt.Println("Before Applying Discount, pen -", pen)
	ApplyDiscount(&pen, 10)
	fmt.Println("After Applying Discount, pen -", pen)

	//Pointer syntax overview
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

	//Using pointers
	fmt.Println("Before incrementing, no =", no)
	increment(&no)
	fmt.Println("After incrementing, no =", no)

	//Accessing fields of struct pointers
	pencil := &Product{200, "Pencil", 10}
	fmt.Println((*pencil).Id)
	fmt.Println(pencil.Id)

	//Pointers in method receivers
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

	// passing arrays as arguments
	fmt.Println("Arrays")
	nosArray := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nosArray)
	sortArray(&nosArray)
	fmt.Println(nosArray)

	//slice
	fmt.Println("Slices")
	nosSlice := []int{3, 1, 4, 2, 5}
	fmt.Println(nosSlice)
	sortSlice(nosSlice)
	fmt.Println(nosSlice)

}

func sortArray(list *[5]int) {
	for i := 0; i < (len(list) - 1); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	fmt.Println("[sortArray], list = ", list)
}

func sortSlice(list []int) {
	for i := 0; i < (len(list) - 1); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	fmt.Println("[sortSlice], list = ", list)
}

func increment(val *int) {
	fmt.Println("Address of val =", val)
	*val++
}

func ApplyDiscount(p *Product, discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
