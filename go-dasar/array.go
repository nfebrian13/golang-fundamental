package main

import "fmt"

func array() {

	/* array */

	var days [7]string
	days[0] = "sunday"
	days[1] = "monday"
	days[2] = "tuesday"
	days[3] = "wednesday"
	days[4] = "thursday"
	days[5] = "friday"
	days[6] = "saturday"

	fmt.Println(days[6])
	fmt.Println(days[0])

}
