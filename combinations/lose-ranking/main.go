package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Participant struct {
	name  string
	score float64
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	baseScores := make(map[string]float64)

	for i := 0; i < n; i++ {
		var S, P int
		fmt.Fscan(reader, &S, &P)

		line, _ := reader.ReadString('\n')
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)

		names := strings.Split(line, ",")

		for r := 0; r < P; r++ {
			score := computeScore(S, P, r+1)
			baseScores[names[r]] += score
		}
	}

	var missingS, missingP int
	fmt.Fscan(reader, &missingS, &missingP)

	line, _ := reader.ReadString('\n')
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	missingNames := strings.Split(line, ",")

	var totalP int
	fmt.Fscan(reader, &totalP)

	line, _ = reader.ReadString('\n')
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	finalRanking := strings.Split(line, ",")

	used := make([]bool, missingP)
	current := make([]string, 0, missingP)

	var answer []string

	var backtrack func()
	backtrack = func() {
		if len(current) == missingP {
			if checkPermutation(current, baseScores, missingS, missingP, finalRanking) {
				answer = append([]string{}, current...)
			}
			return
		}

		for i := 0; i < missingP; i++ {
			if !used[i] {
				used[i] = true
				current = append(current, missingNames[i])
				backtrack()
				current = current[:len(current)-1]
				used[i] = false
			}
		}
	}

	backtrack()

	for _, name := range answer {
		fmt.Println(name)
	}
}

func computeScore(S, P, R int) float64 {
	if P == 1 {
		return float64(S)
	}
	return float64(S) * float64(P-R) / float64(P-1)
}

func checkPermutation(order []string, base map[string]float64, S, P int, final []string) bool {

	scores := make(map[string]float64)
	for k, v := range base {
		scores[k] = v
	}

	for i := 0; i < P; i++ {
		score := computeScore(S, P, i+1)
		scores[order[i]] += score
	}

	list := make([]Participant, 0)
	for name, score := range scores {
		list = append(list, Participant{name, score})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].score == list[j].score {
			return list[i].name < list[j].name
		}
		return list[i].score > list[j].score
	})

	if len(list) != len(final) {
		return false
	}

	for i := 0; i < len(list); i++ {
		if list[i].name != final[i] {
			return false
		}
	}

	return true
}
