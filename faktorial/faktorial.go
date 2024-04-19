package faktorial

func Faktorial(n int, c chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	c <- result 

}
