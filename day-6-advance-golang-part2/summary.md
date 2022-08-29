<h1> REST API </h1>

API (Application Programming Interface) merupakan sekumpulan (method, fungsi atau URL endpoint) yang digunakan untuk mengembangkan aplikasi lebih dari satu platform yang berbeda. Dengan menggunakan API maka kita dapat menggunakan sumber daya dari aplikasi lain tanpa perlu mengetahui bagaimana aplikasi itu dibuat.

REST (REpresentational State Transfer) merupakan standar arsitektur komunikasi berbasis web yang sering diterapkan dalam pengembangan layanan berbasis web. Umumnya menggunakan HTTP (Hypertext Transfer Protocol) sebagai protocol untuk komunikasi data.

REST API merupakan implementasi dari API dengan cara mengakses layanan web dengan cara yang sederhana dan fleksibel tanpa pemrosesan apa pun. 

<h2> HTTP/HTTPS </h2>

* HTTP atau Hypertext Transfer Protocol adalah protokol tingkat aplikasi/website untuk sistem informasi terdistribusi, kolaboratif, dan hypermedia.
* HTTPS atau Hypertext Transfer Protocol Secure adalah protokol HTTP yang dienkripsi untuk mengamankan komunikasi. 

> HTTPS -> versi secure HTTP -> proses komunikasi ada proses encrypt.

Bagian HTTP/HTTPS:
* HTTP messages/body/payload merupakan bagaimana data dipertukarkan antara server dan client. Terdapat dua jenis pesan, yaitu **Request** yang dikirim oleh client untuk memicu tindakan di server dan **Response** yang merupakan jawaban dari server.
* HTTP Header memungkinkan client dan server menyampaikan informasi tambahan dengan permintaan atau respon HTTP. Header HTTP terdiri dari nama case-insensitive diikuti oleh titik dua (:), kemudian dengan nilainya. Spasi sebelum nilai diabaikan.

> body/payload -> berhubungan dengan request and response.
> HTTP header bagian berisi informasi tambahan yang nantinya dibutuhkan dari sisi server. hasil autentifiksi biasanya disimpan di header.

<h2> Methods REST API </h2>

1. GET -> digunakan untuk membaca atau mendapatkan data dari server.

2. POST -> digunakan untuk membuat atau menambahkan data baru dengan menyisipkan data dalam bosy saat request dilakukan.
> untuk create atau seperti proses login. WHY? akan mengcreate session or token untnuk memasang end point dari server.

3. PUT -> digunakan untuk update data, prosesnya dia bakal ngerequest semua existing datanya.

4. Patch -< digunakan untuk update dispesikif data.

5. DELETE -> digunakan untuk mendelete spesifik data dengan mengirimkan path param. 

<h2> Status Code </h2>

merupakan kode standart dalam menginformasikan hasil request ke client.

1. **2xx : success** -> Menunjukkan bahwa request klien berhasil diterima.
    * **200 - OK** -> Menunjukkan bahwa request telah berhasil
    * **201 - created** -> Menunjukkan bahwa request telah berhasil dan resource baru telah dibuat sebagai hasilnya.
    * **202 - accepted** -> Menunjukkan bahwa request telah diterima tetapi belum diselesaikan. Ini biasanya digunakan dalam request menjalankan log dan pemrosesan batch.

2. **4xx : client error** -> Request yang dikirimkan client tidak valid.
    * **400 - Bad Request** -> Request tidak dapat dipahami oleh server karena sintaks yang salah. Client TIDAK HARUS mengulangi permintaan tanpa modifikasi.
    * **401 - Unauthorized** -> Menunjukkan bahwa request memerlukan informasi otentikasi pengguna. Client harus mengulangi request authorisasi yang sesuai.
    * **403 - Forbidden** -> Request tidak sah. Client tidak memiliki hak akses ke konten. Tidak seperti 401, identitas Client diketahui oleh server.
    * **404 - Not Found** -> Server tidak dapat menemukan resource yang diminta.
    * **405 - Method Not Allowed** -> Metode HTTP request diketahui oleh server tetapi telah dinonaktifkan dan tidak dapat digunakan untuk resource tersebut.

3. **5xx : Server Error** -> Request mengalami kesalahan pada sisi server.
    * **500 - Interna Server Error** -> Server mengalami kondisi tak terduga yang mencegahnya memenuhi request.
    * **502 - Bad Gateway** -> Server mendapat respons yang tidak valid saat bekerja sebagai gateway untuk mendapatkan respons yang diperlukan untuk menangani request.
    * **504 - Gateway Timeout** -> Server bertindak sebagai gateway dan tidak bisa mendapatkan respons tepat waktu untuk request.

<h2> JSON (JavaScript Object Notation) </h2>

JSON adalah format pertukaran data yang ringan. Memudahkan manusia untuk membaca dan menulis. Sangat mudah bagi mesin untuk menguraikan dan menghasilkan.

<h2> Authorization </h2>

JSON Web Token (JWT) - adalah standar terbuka (RFC 7519) yang mendefinisikan cara yang ringkas dan mandiri untuk mentransmisikan informasi antar pihak secara aman sebagai objek JSON. Informasi ini dapat diverifikasi dan dipercaya karena ditandatangani secara digital. JWT dapat ditandatangani menggunakan rahasia (dengan algoritma HMAC) atau pasangan kunci publik/pribadi menggunakan RSA atau ECDSA.

Format JWT:
```
<header>.<payload>.<signature>
```
* Header -> Berisi Algoritma & jenis token
* Payload -> Yang berisi klaim. Klaim adalah pernyataan tentang suatu entitas (biasanya, pengguna) dan data tambahan.
* Signature -> Untuk membuat sign yang ditujukan untuk encoded header, payload, serta algoritma yang tersimpan di header.

<h2> Tools pemanggil API </h2>

Disini untuk tools API menggunakan Postman. Postman sendiri merupakan sebuah tools khusus yang dapat melakukan pengujian API tanpa memerlukan sebuah tampilan Front End.

<h2> Echo Framework </h2>

Echo adalah framework bahasa golang untuk pengembangan aplikasi web. Framework ini cukup terkenal di komunitas. Echo merupakan framework besar, di dalamnya terdapat banyak sekali dependensi. Echo digunakan karena *High performance, extensible, minimalist Go web framework*.

<h2> Unit Testing </h2>

Unit testing adalah proses pengembangan perangkat lunak di mana bagian aplikasi terkecil yang dapat diuji, yang disebut unit, diperiksa secara individual dan independen untuk operasi yang tepat. Unutk melakukan unit testing terdapat beberapa hal yang harus diperhatikan:
* File pengujian unit harus diakhiri dengan nama _ testing.go dalam paket yang sama.
* Tes unit di Go ditulis dalam bentuk fungsi, yang memiliki parameter tipe *testing.T, dengan nama fungsi dimulai dengan kata Test.

<h2> Logging </h2>

Dalam komputasi, file log adalah file yang merekam peristiwa baik yang terjadi dalam sistem operasi atau perangkat lunak lain yang berjalan, atau pesan antara pengguna yang berbeda dari perangkat lunak komunikasi. Logging adalah tindakan menyimpan log. Dalam kasus yang paling sederhana, pesan ditulis ke satu file log.

Logging berfungsi untuk menemukan bug di aplikasi, menemukan masalah kinerja, dan melakukan analisis post-mortem dari *outages and security incidents.*