package main

import (
	"fmt"
	"math/rand"
	"time"
)

//打印9*9乘法表函数；
func multiplication(r int) {
	for i := 1; i <= r; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

}

//随机生成20个非0正整数，100以内的，打印出来，并且求奇数的和，偶数的积；
func randomNumber(r int, t int) {
	rand.Seed(time.Now().UnixNano())
	var (
		product int = 1
		sum     int
	)
	for i := 1; i <= r; i++ {

		a := rand.Intn(t) + 1
		if i%2 == 0 {

			product *= a

		} else {
			sum += a
		}
		fmt.Printf("第%d个随机数是%d\n", i, a)

	}
	fmt.Printf("20个随机数数单数和为%d,双数积为%d\n", sum, product)
}

//打印100以内的斐波那契数；
func fibonacci(n int) {
	a, b := 1, 1
	if n < 2 {
		fmt.Println(n)
	} else {
		fmt.Println(a)
		fmt.Println(b)
		for {
			b, a = a+b, b
			if b >= n {
				break
			}
			fmt.Println(b)
		}
	}

}

func main() {
	fmt.Println("9*9乘法表打印")
	multiplication(9)
	fmt.Println("100内的20个随机数单数求和，双数求积")
	randomNumber(20, 100)
	fmt.Println("斐波那契数列100内的打印")
	fibonacci(100)
//gittset
}
