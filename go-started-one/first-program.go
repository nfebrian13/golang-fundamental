package constant

import (
	"fmt"
)

/* 1. constant adalah deklarasi variable,
      constant ditandai dengan keyword const.
   2. constant dapat bertipe apapun bisa character, boolean, string atau numeric.
   3. constant tidak dapat dideklarasikan seperti berikut :=
*/
const (
	message = "hello world"
)

func constant() {
	fmt.Println(message)
}
