package main
import (
	"fmt"
)
func faktorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
func permutasi(n, r int) int {
	if n < r {
		return 0
	}
	return faktorial(n) / faktorial(n-r)
}
func kombinasi(n, r int) int {
	if n < r {
		return 0
	}
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
func main() {
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutasi(a, c)
	c1 := kombinasi(a, c)
	p2 := permutasi(b, d)
	c2 := kombinasi(b, d)
	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
