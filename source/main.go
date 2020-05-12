package main

import (
	"fmt"
	"math"
)

func main() {
	var x_first float64
	var z_first float64
	var angle_first float64
	var x_second float64
	var z_second float64
	var angle_second float64

	x_first = 231
	z_first = 89
	angle_first = -54

	x_second = 437
	z_second = 27
	angle_second = -45

	var k_first float64

	k_first = math.Tan((angle_first + 90) * (math.Pi / 180))
	fmt.Println(k_first)

	var b_first float64

	b_first = z_first - k_first*x_first
	fmt.Println(b_first)

	var k_second float64

	k_second = math.Tan((angle_second + 90) * (math.Pi / 180))
	fmt.Println(k_second)

	var b_second float64

	b_second = z_second - k_second*x_second
	fmt.Println(b_second)

	var x float64
	var z float64

	x = (b_second - b_first) / (k_first - k_second)
	z = k_first*x + b_first

	fmt.Println("")
	fmt.Println("x: ", x, "z: ", z)
}
