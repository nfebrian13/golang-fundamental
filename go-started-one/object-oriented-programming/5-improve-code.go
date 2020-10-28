package main

import (
	"fmt"
	"strings"
)

func main() {

	// planCapacities := []float64{30, 30, 30, 60, 60, 100}
	plants := []PowerPlant{
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 30, active},
		PowerPlant{wind, 25, inactive},
		PowerPlant{wind, 35, active},
		PowerPlant{solar, 45, unavailable},
		PowerPlant{solar, 40, inactive},
	}

	// gridLoad := 75. // arti titik (.) mendandakan bahwa inisialisasi bilangan decimal untuk type inference
	grid := PowerGrid{300, plants}
	// activePlants := []int{0, 1}

	fmt.Println("1. Generate power plan report")
	fmt.Println("2. Generate power grid report")

	var option string
	fmt.Scanln(&option)

	// if option, err := requestOption(); err == nil {
	// 	fmt.Println()

	switch option {
	case "1":
		grid.generatePlantReport()
	case "2":
		grid.generateGridReport()
	default:
		fmt.Println("Uknown option, nothing to do")
	}

	// } else {
	// 	fmt.Println(err.Error())
	// }

	// if option != "1" && option != "2" {
	// 	err = errors.New("invalid option selected")
	// }

	// return
}

func generatePlanCapacityReport(planCapacities []float64) {
	for i, e := range planCapacities {
		fmt.Println("Plan is: ", i, " Capacity is:", e)
	}
}

func generatePowerGridReport(activePlants []int, planCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantId := range activePlants {
		capacity += planCapacities[plantId]
	}
	fmt.Println("Capacity", gridLoad)
	fmt.Println("Load", capacity)
	fmt.Println("Utilization", gridLoad/capacity*100)
}

type PlantType string

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

type PlantStatus string

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
	for idx, p := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", idx)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%s\n", "Type:", p.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity:", p.capacity)
		fmt.Printf("%-20s%s\n", "Status:", p.status)
		fmt.Println()
	}
}

func (pg *PowerGrid) generateGridReport() {
	capacity := 0.
	for _, p := range pg.plants {
		if p.status == active {
			capacity += p.capacity
		}
	}

	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Load:", pg.load)
	fmt.Printf("%-20s%.2f%%\n", "Utilization:", pg.load/capacity*100)
}
