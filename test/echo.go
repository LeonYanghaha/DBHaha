package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, fors string
	t := time.Now()
	s = strings.Join(os.Args[1:], "-")
	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)
	fmt.Println(s)

	fort := time.Now()
	for _, val := range os.Args[1:] {
		fors += val + "-"
	}
	fortEnd := time.Since(fort)
	fmt.Println(fortEnd)
}
