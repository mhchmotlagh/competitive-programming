package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {

		var n int
		fmt.Scan(&n)

		a := make([]uint64, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}

		if n == 1 {
			fmt.Println(0)
			continue
		}

		prefix := make([]uint64, n)
		suffix := make([]uint64, n)

		prefix[0] = a[0]
		for i := 1; i < n; i++ {
			prefix[i] = prefix[i-1] | a[i]
		}

		suffix[n-1] = a[n-1]
		for i := n - 2; i >= 0; i-- {
			suffix[i] = suffix[i+1] | a[i]
		}

		var ans uint64 = 0

		for i := 0; i < n; i++ {

			var cur uint64

			if i == 0 {
				cur = suffix[1]
			} else if i == n-1 {
				cur = prefix[n-2]
			} else {
				cur = prefix[i-1] | suffix[i+1]
			}

			if cur > ans {
				ans = cur
			}
		}

		fmt.Println(ans)
	}
}
