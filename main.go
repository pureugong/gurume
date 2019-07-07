package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const categoryPrefix = "="

func main() {
	var fp *os.File
	var err error

	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		fmt.Printf(">> read file: %s\n", os.Args[1])
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	reader := bufio.NewReaderSize(fp, 4096)
	var category string
	for {
		// read
		line, _, err := reader.ReadLine()
		str := string(line)
		str = strings.TrimSpace(str)

		// skip empty line
		if str != "" {
			// 1. cateogry (also includes info, review, hotel)
			if strings.HasPrefix(str, categoryPrefix) {
				// 1.a. remove =
				category = strings.Split(str, categoryPrefix)[1]
				// 1.b trim category string
				category = strings.TrimSpace(category)
			} else {

				cnt := strings.Count(str, "-")
				if cnt == 2 || cnt == 3 {
					fmt.Printf("%s - %s\n", category, str)
				} else {
					fmt.Printf("exception: %s - %s\n", category, str)
				}

			}
			// fmt.Println(str)
		}

		// end of file
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
}
