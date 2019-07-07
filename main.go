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
				// trim category string
				category = strings.Split(str, categoryPrefix)[1]
				category = strings.TrimSpace(category)
			} else {

				// TODO: review, info, hotel handling
				if strings.Contains(category, "review") || strings.Contains(category, "info") {
					//
				} else if strings.Contains(category, "hotel") {
					// subCategory - name
					hotel := strings.Split(str, "-")
					subCategory := strings.TrimSpace(hotel[0])
					if strings.Contains(subCategory, "델리") {
						subCategory = fmt.Sprintf("%s, %s", subCategory, "베이커리")
					}
					name := strings.TrimSpace(hotel[1])

					fmt.Printf("%s, %s - N/A - N/A - %s\n", category, subCategory, name)
				} else if strings.Contains(category, "노포 식당") {
					// town - name - since
					nopo := strings.Split(str, "-")
					town := strings.TrimSpace(nopo[0])
					name := strings.TrimSpace(nopo[1])
					since := strings.TrimSpace(nopo[2])

					fmt.Printf("%s - %s - N/A - %s (since %s)\n", category, town, name, since)
				} else {
					cnt := strings.Count(str, "-")

					if cnt == 0 {
						// restraunt only
						fmt.Printf("%s - N/A - N/A - %s\n", category, str)
					} else if cnt == 2 || cnt == 3 {
						// town - station - restraunt, town - station - restraunt - note case
						fmt.Printf("%s - %s\n", category, str)
					} else {
						// exception
						fmt.Printf("exception: %s - %s\n", category, str)
					}

					// output:
					// category - town - station - restraunt - note
					// category - town - station - restraunt
				}
			}
		}

		// end of file
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
}
