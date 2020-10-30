package test

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func test() {

	path := flag.String("path", "myapp.log", "The path the log shoul be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
