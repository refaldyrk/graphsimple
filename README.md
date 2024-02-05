# Graphsimple

Repo ini adalah implementasi server GraphQL dengan menggunakan Gqlgen dan integrasi dengan MongoDB untuk manajemen data terkait transportasi umum.

## Instalasi

1. Pastikan Anda telah menginstal Go di komputer Anda. Jika belum, dapat diunduh dan diinstal dari [https://golang.org/dl/](https://golang.org/dl/).

2. Dapatkan proyek ke dalam direktori lokal:

    ```bash
    git clone https://github.com/refaldyrk/graphsimple.git
    cd graphsimple
    ```

3. Install dependensi proyek:

    ```bash
    go mod tidy
    ```

4. Pastikan MongoDB sudah terinstal. Jika belum, dapat diunduh dan diinstal dari [https://www.mongodb.com/try/download/community](https://www.mongodb.com/try/download/community).

## Menjalankan Proyek

1. Jalankan server GraphQL:

    ```bash
    go run server.go
    ```

   Server akan berjalan di [http://localhost:8080](http://localhost:8080) atau port yang telah Anda konfigurasi.

2. Buka GraphQL Playground di browser Anda [http://localhost:8080/playground](http://localhost:8080/playground).

3. Anda dapat mulai menjalankan query dan mutation di dalam GraphQL Playground untuk mengakses dan mengelola data terkait transportasi umum.

## Struktur Direktori

- `graph/`: Berkas-berkas terkait GraphQL dan resolvers.
- `model/`: Definisi struktur data model GraphQL.
- `server.go`: Berkas utama untuk menjalankan server GraphQL.