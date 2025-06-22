package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func exact_integer_sqrt(n int) (int, bool) {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt, sqrt*sqrt == n
}

func main() {
	var n int
	_, err := fmt.Scanln(&n)
    if err != nil {
        log.Fatal(err)
    }

	for i := 0; i < n; i++ {
		var year string
		_, err := fmt.Scanln(&year)
		if err != nil {
			log.Fatal(err)
		}

		year_int, err := strconv.Atoi(year)
		if err != nil {
			log.Fatal(err)
		}

		integer_sqrt, is_perfect_square := exact_integer_sqrt(year_int)

		if ! is_perfect_square {
			fmt.Println(-1);
			continue;
		}

		var a, b int
		a = integer_sqrt / 2;
		b = integer_sqrt - a;

		fmt.Println(a, b);
	}
}