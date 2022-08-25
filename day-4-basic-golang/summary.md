<h1> Introduction to Golang </h1>

Golang (*Go Language*) merupakan bahasa pemrograman yang dikelola oleh google yang bekerja sama dengan 3 orang handal pada tahun 2009. Hampir semua *tech company* di Indonesia sudah menggunakan Go, di Efishery sendiri hampir 90% telah menggunakan Go sebagai bahasa pemrogramannya.

<h2> Data Type </h2>

* Numeric Non-Decimal

    | Data Type       | Range       |
    | ---------- | ---------- |
    | uint8   | 0 - 255  |
    | uint16 | 0 - 65535   |
    | uint32   | 0 - 4294967295  |
    | uint64   | 0 - 18446744073709551615  |
    | uint   | sama dengan uint32 atau uint64 (tergantung nilai)  |
    | byte   | sama dengan uint8  |
    | int8   | -128 - 127  |
    | int16   | -32768 - 32767  |
    | int32   | -2147483648 - 2147483648  |
    | int64   | -9223372036854775808 - 9223372036854775807  |
    | int   | sama dengan int32 atau int64 (tergantung nilai)  |
    | rune   | sama dengn int32  |

* Numeric Decimal

    Contoh:
    ```sql
    var decimalNumber = 2.62

    fmt.Println("bilangan desimal: %f\n", decimalNumber)
    ```
* Boolean
    
    Digunakan untuk memberikan 2 nilai yaitu TRUE dan FALSE.
    Contoh:
    ```sql
    var exist bool = true
    fmt.Println("exist? %t \n", exist)
    ```
* String

    Digunakan untuk data yang berup teks.
    Contoh:
    ```sql
    var message = 'Nama saya "John". Salam kenal. Mari belajar "Golang".'
    fmt.Println(message)
    ```  
<h2> Variable </h2>

* Variables Declaration
    
    Contoh:
    ```sql
    package main

    import "fmt"

    func main() {
        var firstName string = "John"
        var lastName string
        lastName = "Wick"

        fmt.Println("halo %s %s!\n", firstName, lastName)
    }
    ```  
* Declaration Variables without data type

    Contoh:
    ```sql
    var firstName string = "John"
    lastName := "Wick"

    fmt.Println("halo %s %s!\n", firstName, lastName)
    ```
* Declartion Multi Variables

    Contoh:
    ```sql
    var first, second, third string
    first, second, third = "satu", "dua", "tiga"

    seventh, eight, ninth := "tujuh", "delapan", "sembilan"
    ```
* Declartion Underscore Variables
    
    Untuk menyimpan variabel yang tidak digunakan.
    Contoh:
    ```sql
    _ = "belajar Golang"
    _ = "Golang itu mudah"
    name, _ := "John", "Wick"
    ```
* Constants

    Pada dasarnya constant adalah seperti variable, namu nilainya tidak bisa diubah setelah proses inisialisasi. Syntax dasar constant:

    ```sql
    const nama_constant type_data = value_assignment
    ```
<h2> Condition </h2>

* if..else..then
    
    Contoh:
    ```sql
    	var point = 8

	if point = 10 {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}
    ```
* Switch..case
* Comparison Operator -> && (Logical AND), || (Logicl OR) dan ! (Logical NOT)

<h2> Looping </h2>

* for..range

    Contoh:
    ```sql
    var buah = [3]string{"apel", "nanas", "melon"}

	for i, value := range buah {
		fmt.Println("index: ", i)
		fmt.Println("value: ", value)
		fmt.Println("---------------")
	}
    ```
* for..loop
* for..loop break continue

    Contoh:
    ```sql
    for i := 0; i<10; i++ {
		if i == 2{
			continue
		}

		if i ==8{
			break
		}
		fmt.Println("number: ", i)
		fmt.Println("---------------")
	}
    ```
<h2> Function </h2>

Format syntax function:
```sql
func nama_function (parameter) {
    perintah
    }
```
Function dengan return value:
```sql
func nama_function (parameter) return-type-data{
    perintah
    return value
}
```
Function dengn multiple return values:
```sql
func nama_function (parameter) (tipe_data_return1, tipe_data_return2){
    perintah
    return value1, value2
}
```
Function Variadic merupakan suatu bentuk fungsi yang bisa memiliki input lebih dari satu, meskipun hanya memiliki satu parameter, asalkan masih dalam tipe data yang sama.

<h2> Struct </h2>

Struct adalah kumpulan definisi variabel dan atau fungsi, yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi.
``` sql
type student struct {
    name string
    grade int
}

func main() {
    vr s1 student
    s1.name = "John Wick"
    s1.grade = 2

    fmt.Println("name : ", s1.name)
    fmt.Println("grade : ", s1.grade)
}
```
* Embedded Struct

    Memasukkan atribut struct satu ke struct yang lain.
    ```sql
    type person struct {
        name string
        age int
    }

    type student struct {
        grade int
        person
    }
    ```
* Combine Struct with Slice

<h2> Data Structure </h2>

<h3> Map </h3>

Map biasanya digunakan untuk memprsing data JSON.
```sql
	var student map[string]int
	student = map[string]int{}

	student["nurhuda"] = 90
	student["bikhoir"] = 100

	fmt.Println("student: ", student)
	fmt.Println("student spesific: ", student["nurhuda"])
```
<h3> Array </h3>

Merupakan kumpulan data yang bertipe sam dalam suatu variabel. Di Go jumlah kapasitas array haru di definisikan. Agar lebih fleksibel bisa menggunakan *Variadic*.

<h2> Slice </h2>

Merupakan reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama.
```sql
var fruits = []string{"apple", "grape", "banana", "melon"}
fmt.Println(fruits[0]) // "apple"
```
<h2> Defer </h2>

Defer digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai. Meskipun letaknya di awal defer akan dieksekusi paling akhir.