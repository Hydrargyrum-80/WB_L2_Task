package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var (
	OutOfRangeError = errors.New("out of range")
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	fileScan := bufio.NewScanner(file)
	var resultText []string
	for fileScan.Scan() {
		resultText = append(resultText, fileScan.Text())
	}
	return resultText
}

func getValue(str string, col int) (string, error) {
	splitStr := strings.Split(str, " ")
	if col >= len(splitStr) || col < 0 {
		return "", OutOfRangeError
	}
	return splitStr[col], nil
}

func main() {
	kFlag := flag.Int("k", 1, "")
	nFlag := flag.Bool("n", false, "")
	rFlag := flag.Bool("r", false, "")
	uFlag := flag.Bool("u", false, "")
	flag.Parse()
	if *nFlag {
		rows := ReadFile(flag.Arg(0))
		sort.Slice(rows,
			func(i, j int) bool {
				str1, err := getValue(rows[i], *kFlag-1)
				if err != nil {
					log.Fatal(err.Error())
				}
				str2, err := getValue(rows[j], *kFlag-1)
				if err != nil {
					log.Fatal(err.Error())
				}
				result := str1 < str2
				if err != nil {
					log.Fatal(err.Error())
				}
				return result == !*rFlag
			})
		if *uFlag {
			set := make(map[string]struct{})
			var result []string
			for _, i := range rows {
				if _, ok := set[i]; !ok {
					result = append(result, i)
					set[i] = struct{}{}
				}
			}
			rows = result
		}
		for _, i := range rows {
			fmt.Println(i)
		}
	}
}
