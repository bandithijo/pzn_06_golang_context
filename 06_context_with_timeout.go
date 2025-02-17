package main

/*
Context With Timeout
- Selain menambahkan value ke context, dan juga sinyal cancel, kita juga bisa menambahkan sinyal cancel ke context secara otomatis dengan menggunakan pengaturan timeout
- Dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel secara manual, cancel akan otomatis di eksekusi jika weaktu timeout sudah terlewati
- Pengunaan context dengan timeout sangat cocok ketika misal kita melakukan query ke database atau http api, namun ingin menentukan batas maksimal timeout nya
- Untuk membuat context dengan cancl signal secara otomatis menggunakan timeout, kita bisa menggunakan function context.WithTimeout(parent, duration)
*/
