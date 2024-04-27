---

# Chat App

Chat App adalah aplikasi pesan instan sederhana yang memungkinkan pengguna untuk melakukan obrolan satu lawan satu atau dalam grup. Aplikasi ini dibangun dengan menggunakan bahasa pemrograman Go dan database PostgreSQL.

## Fitur

- Obrolan satu lawan satu
- Obrolan grup dengan peran admin dan anggota
- Status pesan (dibaca/dikirim/tunda)
- Berbagi gambar/dokumen

## Daftar Isi

1. [Instalasi](#instalasi)
2. [Menjalankan Aplikasi](#menjalankan-aplikasi)
3. [API Endpoints](#api-endpoints)
4. [Kontribusi](#kontribusi)

## Instalasi

Untuk menjalankan aplikasi ini di lingkungan lokal Anda, pastikan Anda telah menginstal Go dan PostgreSQL di komputer Anda.

1. **Clone Repository**:
   ```bash
   git clone https://github.com/JDS1305/chat_app.git
   ```

2. **Masuk ke Direktori Proyek**:
   ```bash
   cd chat_app
   ```

3. **Konfigurasi Database**:
   - Buat database PostgreSQL dengan nama `chat_app`.
   - Ubah konfigurasi database di `config/config.json` sesuai dengan pengaturan Anda.

4. **Install Dependensi**:
   ```bash
   go mod tidy
   ```

## Menjalankan Aplikasi

Setelah mengatur konfigurasi database, Anda dapat menjalankan aplikasi dengan perintah:

```bash
go run main.go
```

Aplikasi akan berjalan di `http://localhost:8080` secara default.

## API Endpoints

### Pengguna (Users)

- **POST** `/users`: Membuat pengguna baru.
- **GET** `/users/{id}`: Mendapatkan informasi pengguna berdasarkan ID.

### Grup (Groups)

- **POST** `/groups`: Membuat grup baru.
- **GET** `/groups/{id}`: Mendapatkan informasi grup berdasarkan ID.

### Pesan (Messages)

- **POST** `/messages`: Mengirim pesan baru.
- **GET** `/messages/{id}`: Mendapatkan informasi pesan berdasarkan ID.

### Berkas (Files)

- **POST** `/files`: Mengunggah file baru.
- **GET** `/files/{id}`: Mendapatkan informasi file berdasarkan ID.

## Kontribusi

Kontribusi dalam bentuk laporan bug, saran, atau fitur baru sangat diterima.

---