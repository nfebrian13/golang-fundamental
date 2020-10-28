package branchinglooping

import (
	"fmt"
)

func branchinglooping() {

	planCapacities := []float64{30, 30, 30, 60, 60, 100}

	activePlants := []int{0, 1}
	gridLoad := 75. // arti titik (.) mendandakan bahwa inisialisasi bilangan decimal untuk type inference

	fmt.Println("1. Generate power plan report")
	fmt.Println("2. Generate power plan report")

	var option string
	fmt.Scanln(&option)
	fmt.Println(option)

	switch option {
	case "1":
		for i, e := range planCapacities {
			fmt.Println("Plan is: ", i, " Capacity is:", e)
		}
	case "2":
		var capacity float64 = 0
		for _, plantId := range activePlants {
			capacity += planCapacities[plantId]
		}

		fmt.Println("Capacity", gridLoad)
		fmt.Println("Load", capacity)
		fmt.Println("Utilization", gridLoad/capacity*100)
	default:
		fmt.Println("Uknown option, No option taker")
	}
}
