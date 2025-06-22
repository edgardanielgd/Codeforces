package main
import (
	"fmt"
	"log"
)

func main() {
	var t int
	_, err := fmt.Scanln(&t)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < t; i ++ {
		var n, k int
		_, err := fmt.Scanln(&n, &k)
		if err != nil {
			log.Fatal(err)
		}

		var input string
		_, err = fmt.Scanln(&input)
		if err != nil {
			log.Fatal(err)
		}

		var ones, zeros int = 0, 0;

		for j := 0; j < n; j ++ {
			if input[j] == '1' {
				ones ++;
			} else {
				zeros ++;
			}
		}

		// We can represent this problem as follows
		// Define mx and nx as the maximum and minimum counts of same elements
		// So if a string has more ones than zeros, then mx will be the number of ones and
		// nx the number of zeros
		// 
		// so we can say that:
		//
		// nx + mx = n, where n is the length of the sequence
		//
		// now, we need to make sure that after pairing same elements in the palindrome,
		// we end up with having the same number of ones and zeros in order to match them
		// and leave the rest of the string with non-same pairs
		//
		// so for example, if we have 6 ones and 4 zeros in the string, and we are asked to
		// form 3 same element pairs (or k groups), then we can easily take 4 ones to form 2 groups
		// and 2 zeros to form the remaining group, we end with 2 pairs of ones, and 1 pair of zeros
		// the remaining elements are 2 zeros and 2 ones, which can be paired to form the remaining 2
		// pairs that must not be formed by the same elements, here, we can see that we need to end
		// with the same number of zeros and ones to guarantee that there are exactly k "good" pairs
		//
		// Let's call this remaining number of zeros and ones as b
		// Now, in the previos example we saw that there were 2 pairs of ones and 1 pairs of zero, to sum
		// up to 3 pairs, which is the value of k for this case
		// we can denote d as the number of pairs formed by the initially majoritary group (in this case, the group of ones)
		// and c as the number of pairs formed by the minoritary group (in the example, the group of zeros)
		// so we can write:
		//
		// b + 2c = nx
		// b + 2d = mx
		// c + d = k
		//
		// where b is the final number of elements that must be there after pairing same element groups, c the number of groups
		// conformed by the minoritary group, and m the number of groups conformed by the majoritary group. nx and mx were already
		// defined above. K is the number of pairs of same elements that we have to have.
		//
		// we know b,c and d are integers, as well as nx, mx and k
		//
		// we can work over this system, and see how we get to the final intuition
		// 
		// b = nx - 2c
		// nx - 2c + 2d = mx
		// d - c = ( mx - nx ) / 2
		// c = k - d
		// d - k + d = ( mx - nx ) / 2
		// 2d = ( mx - nx ) / 2 + k
		// d = ( ( mx - nx ) / 2 + k ) / 2
		//
		// there is then a straight-forward way to solve the problem,
		// just calculate d and c and check if they are integers >= 0, if they are
		// then they are solutions to this system, otherwise there is no solution



		var mx, nx int = 0, 0;

		if ones > zeros {
			mx = ones;
			nx = zeros;
		} else {
			mx = zeros;
			nx = ones;
		}

		difference := mx - nx;

		if difference % 2 == 1 {
			fmt.Println("NO");
			continue;
		}

		d_mul_2 := difference / 2 + k;

		if d_mul_2 % 2 == 1 {
			fmt.Println("NO");
			continue;
		}

		c := k - d_mul_2 / 2;

		if c < 0 {
			fmt.Println("NO");
			continue;
		}

		fmt.Println("YES");
	}
}