package main

import (
	"fmt" //https://github.com/Chandands1/Go_100_Projects.git
)

func hello() {
	fmt.Println("Hello world")
}
func input() {
	var a string
	fmt.Println("enter your name")
	fmt.Scan(&a)
	fmt.Println("Hi ", a)
}
func add(a, b int) {
	fmt.Println(a + b)
}
func isEven(a int) {
	if a%2 == 0 {
		fmt.Println("even")
	}
	fmt.Println("odd")
}
func pos(a int) {
	if a == 0 {
		fmt.Println("zero")
	} else if a < 0 {
		fmt.Println("negitive")
	} else {
		fmt.Println("positive")
	}
}
func max(a, b int) {
	if a > b {
		fmt.Println(a, "is greater")
	}
	fmt.Println(b, "greatest")
}
func swap(a, b int) {
	var t int = a
	a = b
	b = t
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)

}
func leap(a int) {
	if a%4 == 0 {
		if a%100 == 0 {
			fmt.Println("not leap")
		} else if a%400 == 0 {
			fmt.Println("leap year")
		} else {
			fmt.Println("not leap")
		}
		fmt.Println("not leap")
	}
}
func print(a int) {
	for i := 1; i < a; i++ {
		fmt.Println(i)
	}
}
func sum(a int) {
	n := 0
	for i := 0; i < a; i++ {
		n += i
	}
	fmt.Println("natural sum", n)
}
func table(n int) {
	for i := 0; i < n; i++ {

		fmt.Println(n, " X ", i+1, " = ", n*(i+1))
	}
}
func count(a int) {
	coun := 0
	for i := 0; i <= a; i++ {
		coun += 1
	}
	fmt.Println(coun)
}

func fact(n int) {
	fact := 0
	//num := 1
	for n != 0 {
		digit := n % 10
		fact = digit

		for i := fact; i > 0; i-- {

		}
	}
}
func a(n int) {
	for i := 0; i < n; i++ {
		num := 1
		for j := 0; j < i; j++ {
			fmt.Print(num)
			fmt.Print(" ")
			num = num * (i - j) / (j + 1)
			fmt.Print(num)
		}
		fmt.Println("")
	}
}

func recuFact(n int) int {

	if n == 0 || n == 1 {
		return 1
	}
	return n * recuFact(n-1)
}

func recRev(n, rev int) int {

	if n == 0 {
		return rev
	}
	digit := n % 10
	rev = rev*10 + digit

	return recRev(n/10, rev)
}
func revSum(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + revSum(n/10)
}
func readarr() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	sum := 0
	for _, values := range arr {
		sum += values
		//fmt.Println(values)
	}
	sum1 := 0
	var sum2 int
	var even int
	var odd int
	for _, values := range arr {
		if sum1 < values {
			sum1 = values
		}
		if sum2 > values {
			sum2 = values
		}
		if values%2 == 0 {
			even++
		} else {
			odd++
		}
		//fmt.
		fmt.Println(even, odd)
	}
}

func linearSearch() {
	arr := [7]int{10, 20, 30, 40, 54, 45, 56}
	found := false
	for _, values := range arr {
		if values == 20 {
			found = true
		}
	}
	fmt.Println(found)
}

func main() {
	//hello()
	//input()
	//add(3, 4)
	//isEven(29)
	//pos(-2)
	//max(2, 67)
	//swap(2, 3)
	//leap(2000)
	//print(6)
	//sum(8)
	//table(10)
	//count(6578)
	//LCMHCF(12,18)
	//fmt.Println(recuFact(5))
	//fmt.Print(recRev(1234, 0))
	//fmt.Println(revSum(647))
	//readarr()
	linearSearch()
}
