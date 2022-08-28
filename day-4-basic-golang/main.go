package main

import "fmt"

//struct
type barang struct {
	name string
	stock int
}

func main() {
	var toko barang
	toko.name = "minyak"
	toko.stock = 10

	fmt.Println("nama : ", toko.name)
	fmt.Println("stock : ", toko.stock)

	toko.stock = 15
	fmt.Println("nama : ", toko.name)
	fmt.Println("stock : ", toko.stock)

	//map
	var nilai map[string]int
	nilai = map[string]int{}

	nilai["caca"] = 85
	nilai["nata"] = 95

	fmt.Println("nilai: ", nilai)
	fmt.Println("nilai spesific: ", nilai["caca"])

	//Array
	

	//Slice

}



