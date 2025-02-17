package main

/*
Context With Cancel
- Selain menambahkan value ke context, kita juga bisa menambahkan sinyal cancel ke context
- Kapan sinyal cancel diperlukan dalam context?
- Biasanya ketika kita butuh menjalankan proses lain, dan kita ingin bisa memberi sinyal cancel ke proses tersebut
- Biasanya proses ini berupa goroutine yang berbeda, sehingga dengan mudah jika kita ingin membatalkan eksekusi goroutine, kita bisa mengirim sinyal cancel ke context nya
- Namun ingat, goroutine yang menggunakan context, tetap harus melakukan pengecekan terhadap context nya, jika tidak, tidak ada gunanya
- Untuk membuat context dengan cancel signal, kita bisa menggunakan function context.WithCancel(parent)
*/
