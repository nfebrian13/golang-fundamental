package main // deklarasi package

// Importing packages
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/* Apa pun (variabel, jenis, atau fungsi) yang dimulai dengan huruf kapital akan diekspor,
   dan terlihat di luar paket.
*/

/* Apa pun yang tidak dimulai dengan huruf kapital tidak akan diekspor,
   dan hanya terlihat di dalam paket yang sama.*/

func main() {
	// Finding the Max of two numbers
	fmt.Println(math.Max(73.15, 92.46))

	// Calculate the square root of a number
	fmt.Println(math.Sqrt(225))

	// Printing the value of `𝜋`
	fmt.Println(math.Pi)

	// Epoch time in milliseconds
	epoch := time.Now().Unix()
	fmt.Println(epoch)

	// Generating a random integer between 0 to 100
	rand.Seed(epoch)
	fmt.Println(rand.Intn(100))
}
