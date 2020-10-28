package looping

import (
	"fmt"
)

func looping() {

	arr := []string{"foo", "bar", "buzz"}
	for i, element := range arr {
		fmt.Println(i, element)
	}

	/* example map */
	m := make(map[string]string)
	m["first"] = "foo"
	m["second"] = "bar"
	m["third"] = "buzz"

	for i, e := range m {
		fmt.Println(i, e)
	}

}
