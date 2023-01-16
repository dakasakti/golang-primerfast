# PrimerFast

## Soal1

### Endpoint Questions

- METHOD [POST](http://localhost:3000/questions/first)
  => Untuk menghitung uang

  ```
  http://localhost:3000/questions/first
  ```

- RequestBody (required => RAW JSON)

  ```
  {
        "nominal" : "195.000"
  }
  ```

## Soal2

### Endpoint Questions

- METHOD [POST](http://localhost:3000/questions/second)
  => Untuk menghitung data proses

  ```
  http://localhost:3000/questions/second
  ```

- RequestBody (required => RAW JSON)
  ```
  {
      "except" : "telkom",
      "input" : "telecom"
  }
  ```

## Soal3

Paling dari saya pada syntax `FROM golang` harus ditentukan versi image seperti FROM golang:1.19-alpine. tujuannya untuk menginstal aplikasi Go yang akan dimasukkan ke dalam container.

## Soal4

Tujuan utama dari penggunaan arsitektur microservices adalah untuk meningkatkan skalabilitas dan fleksibilitas.

Skalabilitas: untuk mengelola dan mengukur kinerja setiap layanan secara terpisah, sehingga Anda dapat mengoptimalkan kinerja aplikasi Anda dengan menambah atau mengurangi instansi dari layanan yang diperlukan.

Fleksibilitas: untuk membuat perubahan dalam aplikasi Anda tanpa harus mengubah keseluruhan sistem. Anda dapat mengembangkan, mengubah, atau menghapus layanan individual tanpa mempengaruhi layanan lain dalam aplikasi.

Lebih mudah untuk di maintain : karena setiap microservice dapat di maintain secara terpisah dan tidak mempengaruhi microservice lainnya.

Lebih mudah dalam mengintegrasikan teknologi baru : karena setiap microservice dapat dikembangkan menggunakan teknologi yang berbeda-beda.

Lebih mudah dalam mengelola lingkup pengujian dan deployment : karena setiap microservice dapat diuji dan di-deploy secara terpisah.

## Soal5

Index pada sebuah database digunakan untuk meningkatkan performa dalam mencari data. Index bekerja dengan cara menyimpan salinan dari data yang diindeks dalam struktur data yang disebut indeks, yang disusun dalam urutan tertentu. Setiap indeks mengacu pada lokasi data asli dalam tabel.

Pada saat mencari data, database akan menggunakan indeks untuk menemukan data yang dicari dengan lebih cepat daripada harus mengecek setiap baris dalam tabel dan sangat efisien dalam waktu.

## Soal6

### Endpoint Products

- METHOD [POST](http://localhost:3000/products)
  => Untuk menambahkan produk
  ```
  http://localhost:3000/products
  ```
  RequestBody (required => RAW JSON)
  ```
  {
      "kodeProduk" : "DM-001",
      "namaProduk" : "Nama Produk",
      "kuantitas" : 1
  }
  ```
- METHOD [GET](http://localhost:3000/products)
  => Untuk melihat daftar produk
  ```
  http://localhost:3000/products
  ```

### Endpoint Carts

- Headers with user_id (required)
  ```
  "Authorization" : 1
  ```
- METHOD [POST](http://localhost:3000/carts)
  => Untuk menambahkan produk ke dalam cart

  ```
  http://localhost:3000/carts
  ```

  RequestBody (required => RAW JSON)

  ```
  {
      "kodeProduk" : "DM-001",
      "kuantitas" : 1
  }
  ```

- METHOD [GET](http://localhost:3000/carts?namaProduk=jualan&kuantitas=1)
  => Untuk melihat daftar produk di dalam cart

  ```
  http://localhost:3000/carts?namaProduk=jualan&kuantitas=1
  ```

  QueryParam (opsional)

  ```
  "namaProduk" : "jualan"
  "kuantitas : 1
  ```

- METHOD [DELETE](http://localhost:3000/carts/DM-001)
  => Untuk melihat daftar produk di dalam cart
  ```
  http://localhost:3000/carts/DM-001
  ```
  Param with product_id (required)
  ```
  DM-001
  ```
