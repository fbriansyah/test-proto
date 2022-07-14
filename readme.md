Aplikasi sederhana yang mengimplementasikan gRPC dan protobuff 3. Aplikasi ini menggunakan GORM sebagai driver untuk sqlite.

## Setup

Sebelum bisa menjalankan aplikasi ini, kita harus menginstall beberapa aplikasi, seperti:

1. [Protobuff](https://developers.google.com/protocol-buffers/docs/gotutorial)
2. [gRPC](https://grpc.io/docs/languages/go/quickstart/)

## Project Files

Di dalam projek ini ada beberapa file dan folder penting, seperti:

1. **Makefile**: File ini digunakan untuk menyederhanakan command menggunakan make.
2. **users**: Folder ini berisi file .proto yang digunakan untuk menggenerate file **users_grpc.pb.go** dan **users.pb.go**.
3. **server**: Folder yang bersisi script untuk server.
4. **client**: Folder yang bersisi script untuk client.

## Run Application

Sebelum menjalankan aplikasi, kita bisa menjalankan command `make proto` untuk menggenerate file **users_grpc.pb.go** dan **users.pb.go**. Kedua file tersebut digunakan oleh client dan server.

Untuk menjalankan server, kita bisa menggunakan command `make run-server`.

## Create User

Kita bisa menambah user menggunakan command berikut.

```bash
$ go run client/main.go --mode=create-user --data='{"name": "Dummy1", "email": "dummy1@mail.com"}'
```

## Get All User

Untuk mengambil seluruh data user kita bisa menggunakan command berikut
```bash
$ go run client/main.go --mode=get-all-user
```