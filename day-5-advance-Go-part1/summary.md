<h1> Advance Golang Part. 1 </h1>

<h2> Pointer </h2>

------------
Pointer (*reference atau alamat memori*) adalah sebuah variabel yang digunakan untuk menyimpan *memory address of another variable*. Pointer di Golang juga disebut sebagai variabel khusus. Variabel digunakan untuk menyimpan beberapa data pada alamat memori tertentu dalam sistem.

Variabel yang memiliki *reference* atau alamat pointer yang sama, saling berhubungan satu sama lain dan nilainya pasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain (yang referensi-nya sama) yaitu nilainya ikut berubah. Variabel bertipe pointer ditandai dengan adanya tanda asterisk (*) tepat sebelum penulisan tipe data ketika deklarasi.

Contoh:
```
var numA int = 4

//deklarasi pointer dan refernsi variabel numA
var numB *int = &numA
```

<h2> Method </h2>

-----------------
Method Go mirip dengan fungsi Go dengan satu perbedaan, yaitu method berisi argumen penerima di dalamnya. Dengan bantuan argumen penerima, method ini dapat mengakses properti penerima. Method menempel pada 'type' (bisa 'struct' atau tipe data lainnya).

Contoh implementasi Method:
``` Go
type student struct {
    name string
}

func (s student) sayHello() {
    fmt.Println("halo", s.name)
}
```
Memanggil Method:
```Go
func main() {
    var s1 = student{"John"}
    s1.sayHello()
}
```
Method dengan pointer:
```Go
func (s *student) rubahNama(name string){
    s.name = name
}
```

<h2> Public and Private Property Method </h2>

----------------------
Di Go sebenarnya tidak ada istilah public modifier dan private modifier. Yang ada adalah exported yang kalau di bahasa lain ekuivalen dengan public modifier, dan unexported untuk private modifier.
Ketika namanya diawali dengan huruf kapital menandakan kalau exported (atau public) seperti `SayHello()`. Dan sebaliknya, jika diawali huruf kecil, berarti unexported (atau private) seperti `sayHello()`.

<h2> Interface </h2>

--------------------
Interface adalah custom type berupa kumpulan dari 1 atau lebih method signatures. Interface adalah abstract, tidak dapat membuat instance dari interface.
Penerapan Interface:
```
package main

import "fmt"
import "math"

type hitung interface {
    luas() float64
    keliling() float64
}
```
Pada kode di atas, interface hitung memiliki 2 definisi method, luas() dan keliling(). Interface ini nantinya digunakan sebagai tipe data pada variabel, di mana variabel tersebut akan menampung objek bangun datar hasil dari struct yang akan kita buat.
Dengan memanfaatkan interface hitung, perhitungan luas dan keliling bangun datar bisa dilakukan, tanpa perlu tahu jenis bangun datarnya sendiri itu apa.

<h2> Go Routine </h2>

---------------------
Goroutine merupakan salah satu bagian terpenting dalam concurrent programming di Go. Salah satu hal yang membuat goroutine istimewa adalah eksekusinya dijalankan pada prosesor multi-core. Kita bisa menentukan berapa core yang aktif, semakin banyak semakin cepat.

Eksekusi goroutine bersifat asynchronous, menjadikannya tidak saling tunggu dengan goroutine lain. Namun, karena kecepatannya dan berjalan secara concurrent, maka hasil output yang dihasilkan akan teracak atau kadang tidak terlihat.

Untuk menerapkan goroutine, proses yang akan dieksekusi sebagai goroutine harus dibungkus ke dalam sebuah fungsi. Pada saat pemanggilan fungsi tersebut, ditambahkan keyword go di depannya, dengan itu goroutine baru akan dibuat dengan tugas adalah menjalankan proses yang ada dalam fungsi tersebut.
