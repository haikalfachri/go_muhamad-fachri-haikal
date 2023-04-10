# Menggunakan base image golang alpine
FROM golang:alpine

# Mengatur direktori kerja pada image Docker
WORKDIR /app

# Menyalin kode proyek ke dalam direktori /app pada image Docker
COPY . .

# Menjalankan perintah go untuk menginstal dependensi proyek
RUN go mod download
RUN go build -o entry

# Mengekspos port 8000 pada container Docker
EXPOSE 8000

# Menjalankan aplikasi entry ketika container dijalankan
ENTRYPOINT [ "./entry" ]