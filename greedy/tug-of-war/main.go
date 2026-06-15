package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var target string
	fmt.Fscan(in, &target)

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	var rominaSum, aliSum int64
	bribeCandidates := make([]int, 0)

	i := 0

	if i < n {
		rominaSum += int64(a[i])
		if target == "ali" {
			bribeCandidates = append(bribeCandidates, a[i])
		}
		i++
	}

	rominaTurn := false

	for i < n {
		take := 2
		for k := 0; k < take && i < n; k++ {
			if rominaTurn {
				rominaSum += int64(a[i])
				if target == "ali" {
					bribeCandidates = append(bribeCandidates, a[i])
				}
			} else {
				aliSum += int64(a[i])
				if target == "romina" {
					bribeCandidates = append(bribeCandidates, a[i])
				}
			}
			i++
		}
		rominaTurn = !rominaTurn
	}

	var mySum, oppSum int64
	if target == "romina" {
		mySum = rominaSum
		oppSum = aliSum
	} else {
		mySum = aliSum
		oppSum = rominaSum
	}
	if mySum > oppSum {
		fmt.Fprintln(out, 0)
		return
	}

	var bribed []int
	for _, w := range bribeCandidates {
		oppSum -= int64(w)
		bribed = append(bribed, w)
		if mySum > oppSum {
			break
		}
	}

	fmt.Fprintln(out, len(bribed))
	if len(bribed) > 0 {
		for i, w := range bribed {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, w)
		}
		fmt.Fprintln(out)
	}
}
