package main

/*
Context With Value
- Pada saat awal membuat context, context tidak memiliki value
- Kita bisa menambah sebuah value dengan data Pair (key-vaue) ke dalam context
- Saat kita menambah vaue ke context, secara otomatis akan tercipta child context baru, artinya original context nya tidak akan berubah sama sekali
- Untuk membuat menambahkan value ke context, kita bisa menggunakan function context.WithValue(parent, key, value)
*/
