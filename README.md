[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/-H262Fkp)
# Graded Challenge 1 - P2

Graded Challenge ini dibuat guna mengevaluasi pembelajaran pada Hacktiv8 Program Fulltime Golang khususnya pada pembelajaran pembuatan REST API 

## Assignment Objectives
Graded Challenge 1 ini dibuat guna mengevaluasi pemahaman REST API dengan Go Languages sebagai berikut:

- Mampu memahami konsep API
- Mampu membuat REST API From scratch
- Mampu membuat REST API dengan implementasi database sql


## Assignment Directions - Restful API
Kamu ditugaskan untuk membuat sebuah RESTful API menggunakan Go dan MySQL yang bertujuan untuk me-manage data karyawan pada sebuah perusahaan. Web API yang dibuat harus memiliki fitur CRUD(Create, Read/Retrieve, Update, dan Delete) data karyawan yang tersimpan di database.

### Requirements:
- Database Requirement
  - Database harus memiliki sebuah table untuk menyimpan data karyawan, dengan data yang harus disimpan adalah
    - id : unique identifier untuk karyawan.
    - name : nama karyawan.
    - email : alamat email karyawan.
    - phone : no handphone karyawan.
    - created_at : timestamp ketika data dibuat.
    - updated_at : timestamp ketika data terakhir kali diperbaharui.
  - RESTRICTION:
    - Email setiap pegawai harus unik dan tidak boleh ada yang sama antar karyawan
    - Field name, email, dan phone tidak boleh null/kosong
  - Pastikan untuk menyertakan query DDL pada folder project GC ini, buatlah file dengan nama ddl.sql pada root folder.

- Web API harus memiliki beberapa endpoint sebagai berikut:
  - <b>POST</b> /employees - Menyimpan data karyawan baru, 
    - Request dari endpoint ini harus meliputi nama, email, no handphone, dari karyawan baru tersebut
    - Response dari endpoint ini harus berupa sebuah message sukses, dan data karyawan yang berhasil disimpan, jika terdapat kesalahan pada request maka response harus terdiri dari message yang menjelaskan kesalahan pada input request nya.
  - <b>GET</b> /employees/:id - Menampilkan data karyawan berdasarkan ID
    - Response dari endpoint ini adalah data karyawan sesuai dengan ID pada parameter endpoint, jika karyawan dengan ID tidak ditemukan, maka response harus terdiri dari message yang menjelaskan bahwa data karyawan tidak ditemukan
  - <b>GET</b> /employees - Menampilkan seluruh data karyawan
    - Response dari endpoint ini adalah array of object dari karyawan. Data karyawan yang ditampilkan pada endpoint ini hanya ID, nama, dan email saja.
  - <b>PUT</b> /employees/:id - Memperbaharui data karyawan
    - Request dari endpoint ini harus meliputi nama, email, no handphone, dari karyawan baru tersebut
    - Response dari endpoint ini harus berupa sebuah message sukses, dan data karyawan yang berhasil diperbaharui, jika terdapat kesalahan pada request maka response harus terdiri dari message yang menjelaskan, jika karyawan dengan ID tidak ditemukan, maka response harus terdiri dari message yang menjelaskan bahwa data karyawan tidak ditemukan
  - <b>DELETE</b> /employees/:id - Menghapus data karyawan
    - Response dari endpoint ini adalah data karyawan sesuai dengan ID pada parameter endpoint dan message yang yang menjelaskan bahwa proses penghapusan data karyawan berhasil, jika karyawan dengan ID tidak ditemukan, maka response harus terdiri dari message yang menjelaskan bahwa data karyawan tidak ditemukan
  - Setiap endpoint diatas harus menerapkan best practice REST termasuk status code dan http method yang digunakan
  - Setiap endpoint harus terdokumentasi dengan baik (termasuk setiap request dan kemungkinan resopnse yang terjadi) dengan format OpenAPI

- Deployment Requirement
  - Buatlah database pada platform atau Railway(mySQL) dan sambungkan dengan aplikasi anda.
  - Deploy REST API yang sudah anda buat dengan menggunakan platform Heroku, dan pastikan mencantumkan url hasil deployment pada section expected result dan deployment notes.
- Pastikan untuk mengikuti best practice untuk penggunaan environment variable

## Expected Results
- Web API dapat diakses pada _________ (isi dengan url hasil deployment anda).
- Web API memiliki endpoint sebagai berikut
  - POST /employees - Creates a new customer record in the database.
    - request body -> `{ name, email, phone }`
  - GET /employees/:id - Retrieves a specific customer record by ID from the database.
  - GET /employees - Retrieves all customer records from the database.
  - PUT /employees/:id - Updates a customer record in the database.
    - request body -> `{ name, email, phone }`
  - DELETE /employees/:id - Deletes a customer record from the database.

## Assignment Submission
Push Assigment yang telah Anda buat ke akun Github Classroom Anda masing-masing.

### Assignment Rubrics Aspect : 
|Criteria|Meet Expectations|
|---|---|
|Problem Solving|5 API Endpoints are implemented and working correctly|
|Database Design |MySQL database meets the required specifications|
||Database queries are efficient and appropriately indexed|
|Readability|Code is well-documented and easy to read|
||Code includes appropriate comments and documentation|


### Assignment Notes:
- Jangan terburu-buru dalam menyelesaikan masalah atau mencoba untuk menyelesaikannya sekaligus.
- Jangan menyalin kode dari sumber eksternal tanpa memahami bagaimana kode tersebut bekerja.
- Jangan menentukan nilai secara hardcode atau mengandalkan asumsi yang mungkin tidak berlaku dalam semua kasus.
- Jangan lupa untuk menangani negative case, seperti input yang tidak valid
- Jangan ragu untuk melakukan refaktor kode Anda, buatlah struktur project anda lebih mudah dibaca dan dikembangkan kedepannya, pisahkanlah setiap bagian kode program pada folder sesuai dengan tugasnya masing-masing.
- Jangan lupa untuk mendokumentasikan endpoint yang dibuat dengan format OpenAPI.

### Additional Notes
Total Points : 100

Deadline : Diinformasikan oleh instruktur saat briefing GC. Keterlambatan pengumpulan tugas mengakibatkan skor GC 1 menjadi 0.

Informasi yang tidak dicantumkan pada file ini harap dipastikan/ditanyakan kembali kepada instruktur. Kesalahan asumsi dari peserta mungkin akan menyebabkan kesalahan pemahaman requirement dan mengakibatkan pengurangan nilai.

### Deployment Notes
- Deployed url: _________ (isi dengan url hasil deployment anda)
