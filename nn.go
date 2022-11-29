package main

import "fmt"

var x int = 2

type product struct {
	price       float64
	productName string
}

// declaring a function that takes in a struct value and modifies it
func changeProduct(p *product) {
	fmt.Println("viod")
	p.price = 600
	p.productName = "Bicycle"
	// the changes are not seen to the outside world
}

func main() {
	// p := &x

	// fmt.Printf("the address of a pointer is %p,%p", &x, p)
	p := product{
		price:       300,
		productName: "car",
	}
	fmt.Println(p)
	// p.changeProduct()
	changeProduct(&p)
	fmt.Println(p)
	fmt.Println(p.price)

}
