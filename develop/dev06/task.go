package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	dFlag := flag.String("dFlag", " ", "separator")
	fFlag := flag.Int("fFlag", 1, "fields")
	sFlag := flag.Bool("sFlag", false, "separated")
	flag.Parse()
	strScan := bufio.NewScanner(os.Stdin)
	for strScan.Scan() {
		text := strScan.Text()
		columns := slices.DeleteFunc[[]string, string](strings.Split(text, *dFlag), func(s string) bool {
			return s == ""
		})
		if len(columns) < *fFlag || (*sFlag && !strings.Contains(text, *dFlag)) {
			continue
		}
		fmt.Println(columns[*fFlag-1])
	}
}
