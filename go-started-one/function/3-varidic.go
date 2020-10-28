package variadic

import (
	"fmt"
)

func variadic() {
	programLanguage("go", "java", "ruby", "python")
}

func programLanguage(messages ...string) {
	for _, e := range messages {
		fmt.Println(e)
	}
}
