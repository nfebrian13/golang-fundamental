package field

import "fmt"

func field() {
	/* inisialisasi nilai struct cara pertama
	foo := myStruct{}
	foo.myfield = "test"  */

	/* inisialisasi nilai struct cara kedua
	foo := myStruct{myfield: "hallo"} */

	foo := new(myStruct)
	foo.myfield = "hallo nana"
	fmt.Println(foo.myfield)
}

type myStruct struct {
	myfield string
}
