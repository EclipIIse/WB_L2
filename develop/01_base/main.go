package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func timeNow() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(time)
}
func main() {
	timeNow()
}
