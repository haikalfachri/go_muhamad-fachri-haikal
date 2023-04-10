# Section 24 - Clean and Hexagonal Architecture
- Clean Architecture adalah sebuah konsep arsitektur perangkat lunak yang fokus pada pemisahan antara business logic atau logika bisnis dengan teknologi atau infrastruktur yang digunakan untuk mengimplementasikannya.
- Dalam Clean Architecture, aplikasi dibagi menjadi beberapa lapisan atau layer, yaitu:
    1. Entity: Lapisan yang berisi objek dan model bisnis (business model) yang merepresentasikan data dan perilaku bisnis aplikasi.
    2. Use Case: Lapisan yang berisi logic bisnis (business logic) atau aturan bisnis (business rules) yang memproses dan mengubah data dari lapisan entitas.
    3. Interface: Lapisan yang menghubungkan aplikasi dengan dunia luar, seperti antarmuka pengguna (user interface), API, database, dan sebagainya.
    4. Infrastructure: Lapisan yang berisi detail teknis dan implementasi, seperti koneksi database, framework, library, dan sebagainya.
- Tujuan dari Clean Architecture adalah untuk menciptakan aplikasi perangkat lunak yang mudah di-maintain (dapat dipelihara), scalable (dapat diperluas), dan mudah diuji (testable). Konsep Clean Architecture bertujuan untuk menghindari adanya code smell, yaitu kode yang buruk dan sulit dipelihara.