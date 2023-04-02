# Section 22 - Middleware
- Middleware adalah fungsi yang digunakan dalam pengembangan perangkat lunak untuk memodifikasi atau menangani permintaan HTTP sebelum permintaan tersebut diteruskan ke fungsi pengontrol atau handler utama. Dalam konteks web development, middleware biasanya digunakan untuk melakukan tugas-tugas seperti otentikasi pengguna, otorisasi, logging, penanganan kesalahan, kompresi respons, dan banyak lagi.
- Berikut adalah beberapa contoh third-party middleware yang tersedia untuk digunakan pada aplikasi web yang menggunakan bahasa pemrograman Go:
    1. Negroni
    2. Gorilla/mux
    3. Alice
    4. Gin
    5. Echo
- Middleware bekerja sebagai lapisan antara permintaan dan respons pada aplikasi web. Setiap middleware dapat memodifikasi permintaan dan/atau respons, serta melakukan tugas tambahan seperti autentikasi, logging, atau caching. Setiap middleware dipanggil secara berurutan dan jika salah satu middleware menghentikan proses, respons akan dikirim ke klien. Urutan dan pengaturan middleware sangat penting dalam pembuatan aplikasi web yang aman, cepat, dan andal.