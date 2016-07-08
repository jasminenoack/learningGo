package main

import (
	"bufio"
	"fmt"
	"os"
)

// // dup 1
// func main() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		if input.Text() == "" {
// 			break
// 		}
// 		counts[input.Text()]++
// 	}
// 	// ignoring input.Err()
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// // dup 2
// func main() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	if len(files) == 0 {
// 		countlines(os.Stdin, counts, true)
// 	} else {
// 		for _, arg := range files {
// 			f, err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 				continue
// 			}
// 			countlines(f, counts, false)
// 			f.Close()
// 		}
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }
//
// func countlines(f *os.File, counts map[string]int, isStdin bool) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		if isStdin && input.Text() == "" {
// 			break
// 		}
// 		counts[input.Text()]++
// 	}
// 	// ignoring input.Err
// }

// func main() {
// 	counts := make(map[string]int)
// 	for _, filename := range os.Args[1:] {
// 		data, err := ioutil.ReadFile(filename)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
// 			continue
// 		}
// 		for _, line := range strings.Split(string(data), "\n") {
// 			counts[line]++
// 		}
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

type linedata struct {
	files []string
	count int
}

// // dup 2 with file names
func main() {
	data := make(map[string]linedata)
	files := os.Args[1:]
	if len(files) == 0 {
		countlines(os.Stdin, data, true, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countlines(f, data, false, arg)
			f.Close()
		}
	}
	for line, n := range data {
		if n.count > 1 {
			fmt.Printf("%d\t%s\n%v\n\n", n.count, line, n.files)
		}
	}
}
func countlines(f *os.File, data map[string]linedata, isStdin bool, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if isStdin && input.Text() == "" {
			break
		}
		oldStruct := data[input.Text()]
		files := oldStruct.files
		if !exists(files, filename) {
			files = append(files, filename)
		}
		newStruct := linedata{
			count: oldStruct.count + 1,
			files: files,
		}
		data[input.Text()] = newStruct
	}
	// ignoring input.Err
}

func exists(list []string, value string) bool {
	for _, key := range list {
		if key == value {
			return true
		}
	}
	return false
}
