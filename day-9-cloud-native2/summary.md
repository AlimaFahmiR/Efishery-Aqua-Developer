<h2> Container </h2>

automation tools 

salah satunya yg disediakan docker -> docker compose karena banyak hal yang bisa dilakukan dalam satu waktu.

docker biasa -> hanya bisa run satu-satu
docker compose -> bisa di run beberapa container sekaligus

--------------------------
secara proses ada 3 step besar 
- definisi yaml filenya
kalo imagenya gk ada maka akan di build. kalo ada ada maka akan di run dan dibuatkan container.

<h2> Building blocks </h2>

network di dalam docker dan comp kita itu beda.
satu obj service kalo di run akan jadi 1 kontainer.

Volume -> tempat unutk menyimpan data sementara si docker.

db_data:/etc/data. db_data -> local jadi yang bisa kita akses di local comp port 8000
/etc/data -> local container, port 5000

Network -> bisa membuat container nerkomuniksi sat sama lain.

Benefits of container orchestration -> jia app mati atau container maka otomatis di reset. nambah security.

Nomad -> 

service -> blok untuk define si containernya.

docker-compose up -d 
docker network ls
docker ps
docker-compose ps
*docker log*
telnet localhost ......
docker-compose down -> mematikan semua containernya

hyper v nya ngga nyala jadi virtualizationnya harus diaktifkan.

---------------------------------
Di dalam file -> Dockerfile

```
FROM golang

RUN go build -o nama_file-> akan mengenarate binary, dimana o adlaah output

CMD ["./nama_file"]
```

yang bener:

```
FROM golang

WORKDIR /nama_file

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go./

RUN go build -o /main

CMD["/main"]
```

<h2> CI/CD </h2>

COntainer Integration -> u menjanlakn task , statistic analytic
Continuous Delivery -> saat CI selesai dia akan ngebuild aplikasinya

Kelebihan:
1. feedback lebih cepat. misal ada error atau field ketahuannya lebih cepat
2. bug ketahuan lebih cepat
3. speed up the release process

Tools: 

<h2> Github Action </h2>

masing-masing CI CD di masing-masing toolsnya mainnya beda2. 

COmponent:
1. Event 
2. RUnner -> sebuah aplikasi github ynag menjalankan logic di github action.
3. Actions -> bisa di reuce oleh siapapun, bisa publish di open source. 

<h2> Observability </h2>

suatu kegiatan untuk mengukur matriks atau leadaan suatu sistem. 

Kelebihan:

alerting-. punya app yg berjalan di suatu server 
1. 

Challenges:
1. Data silos -> data hanya tersimpan disatu tempat saja.
2. lack -> susah untuk menegetesnya, jika ada trouble troubleshoot nya bukan/ngga ada menambha bisnis value

Tools: 

Tugas: Ngedocker file tugas yang kemarin. docker compose 

Banyak2 nih yap jalanin aja one by one yang penting stay on the track.
yang kemarin tanya2 reni cara mengerjakannya yap.