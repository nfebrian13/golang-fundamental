package branches

import (
	"fmt"
)

func branches() {

	if foo := 1; foo == 1 {
		fmt.Println("Foo")
	} else {
		fmt.Println("Buz")
	}

	switch f := 3; f {
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	}
}
