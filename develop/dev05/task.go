package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func PrintLines(startPos int, endPos int, n bool, lines []string) {
	if startPos < 0 {
		startPos = 0
	}
	if endPos >= len(lines) {
		endPos = len(lines) - 1
	}
	for j := startPos; j <= endPos; j++ {
		if n {
			fmt.Printf("%d. ", j+1)
		}
		fmt.Println(lines[j])
	}
}

func main() {
	AFlag := flag.Int("AFlag", 0, "count after equals")
	BFlag := flag.Int("BFlag", 0, "count before equals")
	CFlag := flag.Int("CFlag", 0, "count between equals")
	cFlag := flag.Bool("cFlag", false, "count")
	ignCaseFlag := flag.Bool("i", false, "ignore case")
	nFlag := flag.Bool("nFlag", false, "display line num")
	FFlag := flag.Bool("FFlag", false, "fixed")
	vFlag := flag.Bool("vFlag", false, "invert")
	flag.Parse()
	str := flag.Arg(0)
	f := flag.Arg(1)
	if *ignCaseFlag {
		str = strings.ToLower(str)
	}
	if *FFlag {
		str = regexp.QuoteMeta(str)
	}
	r, err := os.Open(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer r.Close()
	count := 0
	strScan := bufio.NewScanner(r)
	var lines []string
	var indexes []int
	for i := 0; strScan.Scan(); i++ {
		temp := strScan.Text()
		line := temp
		if *ignCaseFlag {
			line = strings.ToLower(line)
		}
		ok, err := regexp.MatchString(str, line)
		if err != nil {
			log.Fatal(err.Error())
		}
		lines = append(lines, temp)
		if ok {
			count++
			indexes = append(indexes, i)
		}
	}
	if *cFlag {
		fmt.Println(count)
		return
	}
	if *CFlag != 0 {
		AFlag = CFlag
		BFlag = CFlag
	}
	if len(indexes) == 0 {
		return
	}
	if *vFlag {
		index := 0
		for i, lineVal := range lines {
			if index < len(indexes) && indexes[index] == i {
				index++
				continue
			}
			if *nFlag {
				fmt.Printf("%d. ", i+1)
			}
			fmt.Println(lineVal)
		}
		return
	}
	for _, i := range indexes[:len(indexes)-1] {
		PrintLines(i-*BFlag, i+*AFlag, *nFlag, lines)
		fmt.Println()
	}
	PrintLines(indexes[len(indexes)-1]-*BFlag, indexes[len(indexes)-1]+*AFlag, *nFlag, lines)
}
