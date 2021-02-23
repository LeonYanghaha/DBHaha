package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//func main() {
//	counts := make(map[string]int)
//	input := bufio.NewScanner(os.Stdin)
//	count := 0
//	for input.Scan() {
//		counts[input.Text()]++
//		count++
//		if count >= 4 {
//			break
//		}
//	}
//	// NOTE: ignoring potential errors from input.Err()
//	for line, n := range counts {
//		fmt.Printf("%d\t%s\n", n, line)
//	}
//}
