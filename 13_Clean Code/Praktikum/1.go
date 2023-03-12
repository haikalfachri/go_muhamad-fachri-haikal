package main

/* 
- Penamaan struct umumnnya menggunakan kapital agar bisa diakses package lain dan 
membedakan antara variable dan struct,
contoh koreksinya: user -> User
-Tipe data yang digunakan untuk username dan password seharusnya string bukan int,
karena umumnya data tersebut berupa string.
*/
type user struct {
	id       int
	username int
	password int
}
/* 
- Penamaan struct agak sulit dibaca karena terdiri dari 2 kata berbeda yang seharusnya 
untuk setiap kata yang berbeda harus diawali huruf kapital, lalu karena struct umumnya ditulis kapital maka,
contoh koreksinya: userservice -> UserService
- Penamaan variable t yang kurang dapat dipahami, seharusnya variable yang digunakan jelas 
dan mudah dimengerti, 
contoh koreksinya: t -> listUser
*/
type userservice struct {
	t []user
}
/* 
- Penamaan method agak sulit dibaca karena terdiri dari 2 atau lebih kata berbeda yang seharusnya 
untuk setiap kata yang berbeda harus diawali huruf kapital, 
contoh koreksinya: getallusers -> getAllUsers
- Penamaan receiver mungkin lebih deskriptif jika menggunakan kata dibanding huruf
contoh koreksinya: u -> user
*/
func (u userservice) getallusers() []user {
	return u.t
}
/* 
- Penamaan method agak sulit dibaca karena terdiri dari 2 atau lebih kata berbeda yang seharusnya 
untuk setiap kata yang berbeda harus diawali huruf kapital, 
contoh koreksinya: getuserbyid -> getUserById
- Penamaan variable r pada method getuserbyid kurang dapat dipahami,
contoh korekisnya r -> value; r -> account
- Penamaan receiver mungkin lebih deskriptif jika menggunakan kata dibanding huruf
contoh koreksinya: u -> user
*/
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
	return user{}
}

// Tidak ada main function

/* 
- Berapa banyak kekurangan dalam penulisan kode tersebut?
	Sejauh ini saya menemukan 6 kekurangan.
- Bagian mana saja terjadi kekurangan tersebut?
	1. Tidak ada main function
	2. Tipe data untuk username dan password kurang sesuai
	3. Penamaan variable t pada userservice kurang dapat dipahami, 
	penamaan variable r pada method getuserbyid kurang dapat dipahami
	4. Penamaan method dan type struct yang agak sulit dibaca
	5. Penamaan struct umumnya kapital agar bisa diakses package lain dan membedakan antara struct dan variable
	6. Penamaan receiver mungkin lebih deskriptif jika menggunakan kata dibanding huruf
- Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar 
pada kode yang disediakan berikut.
*/