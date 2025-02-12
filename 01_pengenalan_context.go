package main

import "time"

/*
Pengenalan Context
- Context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout, dan sinyal deadline
- Context biasanya dibuat per request (misal setiap ada request masuk ke server web melalui http request)
- Context digunakan untuk mempermudah kita meneruskan value dan sinyal antar proses

Kenapa Context Perlu Dipelajari?
- Context di Go biasa digunakan untuk mengirim data request atau sinyal ke proses lain
- Dengan menggunakan Context, ketika kita ingin membatalkan semua proses, kita cukup mengirim sinyal ke Context, maka secara otomatis semua proses akan dibatalkan
- Hampir semua bagian di Go memanfaatkan Context, seperti database, http server, http client, dan lain-lain
- Bahkan di Google sendiri, ketika menggunakan Go, Context wajib digunakan dan selalu dikirim ke setiap function yang dikirim

Package Context
- Context direpresentasikan di dalam sebuah	interface Context
- interface Context terdapat dalam package context
- <https://pkg.go.dev/context>
*/

type Context interface {
	Deadlie() (deadline time.Time, ok bool)

	Done() <- chan struct{}

	Err() error

	Value(key interface{}) interface{}
}
