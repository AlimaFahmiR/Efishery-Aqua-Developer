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

	
	fmt.Println("---------Struct----------")
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

	fmt.Println("---------Map----------")
	fmt.Println("nilai: ", nilai)
	fmt.Println("nilai spesific: ", nilai["caca"])

	//Array
	var daftar_barang = [...]string{"minyak", "beras", "kopi", "garam"}

	daftar_barang[3] = "telur"

	fmt.Println("---------Array----------")
	for i, barang := range daftar_barang {
		fmt.Printf("barang %d : %s\n", i, barang)
	}
	//Slice
	fmt.Println("---------Slice----------")

	var barang_baru = daftar_barang[0:2]

	fmt.Println(barang_baru)
}



