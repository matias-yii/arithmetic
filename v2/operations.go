package arithmetic

//Factorial factorial of n
func Factorial(n int) int {
	var f int = 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

//Add sums n int numbers
func Add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
