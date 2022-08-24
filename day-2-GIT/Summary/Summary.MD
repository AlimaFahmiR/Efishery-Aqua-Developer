<h1> Git </h1>

Git adalah salah satu Version Control System pada proyek perangkat lunak yang bertugas mencatat setiap perubahan pada file yang dikerjakan baik individu maupun secara tim.

<h2> Git Command </h2>
Berikut merupakan git command yang sering digunakan:

1. Git init : untuk membuat repository pada file lokal yang nantinya ada folder .git
2. Git add : menambahkan file baru pada repository yang dipilih
3. Git remote : menghubungkan repository lokal ke remote server.
4. Git commit : untuk menyimpan perubahan yang dilakukan, tetapi tidak ada perubahan pada remote repository.
5. Git push : untuk mengirimkan perubahan file setelah di commit ke remote repository.
6. Git fetch : mengambil data dari repository tanpa mendownload.
7. Git pull : menggabungkan semua perubahan yang ada di remote repository ke direktori lokal.
8. Git status : untuk mengetahui status dari repository lokal.
9. Git branch : melihat seluruh branch yang ada pada repository
10. Git checkout : menukar branch yang aktif dengan branch yang dipilih
11. Git merge : untuk menggabungkan branch yang aktif dan branch yang dipilih

<h2> Commit Convention </h2>
Commit Convention adalah aturan penulisan pesan saat melakukan commit untuk menyelaraskan hasil komit antar kontributor.
Disini menggunakan style "Git Karma Convention".

Format:

```java
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

Contoh:

```java
fix(middleware): ensure Range headers adhere more closely to RFC 2616

Add one new dependency, use `range-parser` (Express dependency) to compute
range. It is more well-tested in the wild.

Fixes #2310
```

<h2> Semantic Versioning </h2>
Semantic Versioning adalah aturan baku yang dapat digunakan untuk membantu supaya pemberian versi dalam pengembangan software lebih jelas.

Penulisannya sebagai berikut:

```java
v.<major>.<minor>.<patch>
```

Contoh: v1.2.3

Major: berhubungan dengan ditambahkannya fitur baru termasuk ada beberapa API/fitur yang tidak sesuai dengan versi sebelumnya.

Minor: terdapat fitur baru tanpa menghilangkan fitur sebelumnya.

Patch: terdapat perbaikan bug tanpa menghilangkan fitur sebelumnya.

<h2> Git Management </h2>

Di Efishery sendiri sekarang telah menggunakan git management dengan metode Trunk Based Development. Dimana, proses manajemen git dilakukan oleh kontributor hanya pada satu branch yang sama.