package main

import (
	"fmt"
	"time"
)

/*
	func add(x int, y int) (int, int) {
		return x, y
	}

	func swap(x, y string) (string, string) {
		return y, x
	}

	func main() {
		//fmt.Println("hello world")
		//fmt.Println(time.Now())
		//fmt.Println(rand.Intn(10))
		//fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
		//fmt.Println(split(17))
		//a, b := swap("hello", "world")
		//fmt.Print(a + " " + b)
		//var i int
		//var j bool
		//fmt.Print(i, j)
		//var h = 2
		//fmt.Print(h)
		//var i int
		//var f float32
		//var b bool
		//var s string
		//fmt.Printf("%v %v %v %q\n", i, , b, s)
		var x, y int = 3, 4
		var f float64 = math.Sqrt(float64(x*x + y*y))
		var z uint = uint(f)
		fmt.Println(x, y, z)
	}

	func split(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		return
	}
*/
func main() {
	today := time.Now().Weekday()
	fmt.Print(today + 1)

}
