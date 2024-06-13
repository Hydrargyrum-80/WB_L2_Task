package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindAnagram(wordsRU []string) map[string][]string {
	anagramMap := map[string]string{}
	findMap := map[string]map[string]struct{}{}
	resultMap := map[string][]string{}
	for _, i := range wordsRU {
		i = strings.ToLower(i)
		runes := []rune(i)
		sort.SliceStable(runes, func(i, j int) bool {
			return runes[i] > runes[j]
		})
		wordAnagram := string(runes)
		first, ok := anagramMap[wordAnagram]
		if !ok {
			anagramMap[wordAnagram] = i
			first = i
		}
		_, ok = findMap[first]
		if !ok {
			findMap[first] = make(map[string]struct{})
		}
		findMap[first][i] = struct{}{}
	}
	for key, value := range findMap {
		if len(value) == 1 {
			continue
		}
		var resultStrings []string
		for k := range value {
			resultStrings = append(resultStrings, k)
		}
		sort.Strings(resultStrings)
		resultMap[key] = resultStrings
	}
	return resultMap
}

func main() {
	var words = []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	fmt.Println(FindAnagram(words))
}
