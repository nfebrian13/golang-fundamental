package main

import "fmt"

/*
	Fungsi change() memiliki 2 parameter, yaitu original yang tipenya adalah pointer int,
	dan value yang bertipe int.
	Di dalam fungsi tersebut nilai asli parameter pointer original diubah.

	Fungsi change() kemudian diimplementasikan di main.
	Variabel number yang nilai awalnya adalah 4 diambil referensi-nya lalu digunakan sebagai parameter
	pada pemanggilan fungsi change().

	Nilai variabel number berubah menjadi 10 karena perubahan yang terjadi di dalam fungsi change
	adalah pada variabel pointer
*/

func pointerparam() {
	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 10)
	fmt.Println("after  :", number) // 10
}

func change(original *int, value int) {
	*original = value
}
