<h1> Relational Database </h1>

Database adalah kumpulan data yang dikelola sedemikian rupa dengan ketentuan tertentu yang saling berhubungan untuk mempermudah pengelolaannya. Secara umum terdapat 5 jenis database, sebagai berikut:

1. **Operational Database** = Berguna untuk mengelola data yang dinamis secara *real-time*.
2. **Database Warehouse** = Digunakan untuk pelaporan dan analisis data.
3. **Distributed Database** = Basis data yang perangkat penyimpanannya tidak terpasang pada perangkat komputer yang sama.
4. **Relational Database** = Basis data yang mengorganisir berdasarkan model hubungan data.
5. **End-User Database**

Disini kita akan mempelajari tentang **Relational Database** atau biasa disingkat RDB yang merupakan kumpulan item data yang terdiri dari tabel, baris, dan kolom. Data pada RDB dapat diakses dengan berbagai cara tanpa menyusun ulang tabel data itu sendiri. Banyak perangkat lunak yang menggunakan sistem ini untuk mengatur dan memelihara hubungan setiap datanya, salah satunya adalah **PostgreSQL**.

<h2> DDL (Data Definition Language) </h2>

DDL digunakan untuk mendefinisikan struktur seperti skema, database, tabel, *constraint*, dll. Contoh perintah DDL adalah **create, alter** dan **drop**.

| Command | Contoh | Fungsi |
| ----------- | :---------: | ----------: |
| Create | `create database database_name;` | Membuat tabel atau database baru |
| Alter | `alter table table_name rename to table_name_change;` | Mengganti nama tabel, menambah, mengubah, menghapus kolom |
| Drop | `drop table table_name;` | Menghapus database atau tabel |

<h2> DML (Data Definition Language) </h2>

DML digunakan untuk melakukan manipulasi atau pengolahan data dalam tabel. Contoh perintah DML adalah **select, insert, update** dan **delete**.

| Command | Contoh | Fungsi |
| ----------- | :---------: | ----------: |
| Select | `select * from users;` | Menampilkan data dari tabel |
| Insert | `insert into public.table_name (name, age) values ('zizi', 25);` | Memasukkan data ke dalam tabel |
| Update | `update public.table_name set age = 25 where id = 1;` | Mengubh data pada tabel |
| Delete | `delete from public.table_name where id = 1;` | Menghapus data pada tabel |

<h2> DCL (Data Control Language) </h2>

DCL digunakan untuk merubah hak akses, memberikan roles, dan isu lain yang berhubungan dengan keamanan database. Contoh perintah DML adalah **grant** dan **revoke**.

<h2> JOIN </h2>

Join digunakan untuk menggabungkan data atau baris dari satu (self-join) atau lebih tabel berdasarkan bidang yang sama melalui kolom atau key tertentu yang memiliki nilai terkait untuk mendapatkan satu set data dengan informasi lengkap.

Berikut merupakan penggambaran dari masing-masing tipe join:

![]( ../day-3-Relational_Database/screenshot/join%20type.png)

<h2> Aggregation </h2>

Agregat menghasilkan satu hasil untuk seluruh grup atau tabel. Agregat digunakan untuk menghasilkan hasil yang ringkas. Fungsi agregat biasanya digunakan bersamaan dengan GROUP BY, yaitu untuk mengelompokkan data berdasarkan kriteria yang diinginkan.

Fungsi agregat yang sering digunakan adalah `count`, `sum`, `max`, `min`, dan `avg`.

<h2> Subquery </h2>
Subquery adalah query di dalam query SQL lain yang lebih besar dan tertanam dalam kalusa WHERE.

Contoh:
```sql
update products set stock = subquery.stock - 2 from (select id, stock from products where id = 1) as subquery where products.id = 1
```

<h2> Function </h2>

Function atau juga dikenal sebagai *Stored Procedures*, memungkinkan untuk melakukan operasi yang biasanya digunakan untuk mengambil beberapa query dan melakukan beberapa hal dalam satu fungsi database.

Contoh:
``` sql
crete function kurangi_stock(int, int) returns products as
'update products set stock = subquery.stock -$2 from (select id, stock from products where id = $1) as subquery where products.id = $1;
select * from products where id = $1'
language 'sql';