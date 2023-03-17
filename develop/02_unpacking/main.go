package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

func StringUnpacking(s string) string {
	var scan scanner.Scanner
	var res string
	var prev string
	scan.Init(strings.NewReader(s))
	scan.Mode = scanner.ScanChars | scanner.ScanInts
	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		switch tok {
		case scanner.Int:
			num, err := strconv.Atoi(scan.TokenText())
			if err != nil {
				fmt.Printf("Something went wrong %s", err)
			}
			res += strings.Repeat(prev, num-1)
		default:
			prev = scan.TokenText()
			res += scan.TokenText()
		}
	}
	return res
}

func main() {
	fmt.Println(StringUnpacking("a4bc2d5e"))
}
