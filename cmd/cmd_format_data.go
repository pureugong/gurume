package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const categoryPrefix = "="
const dataDir = "/tmp/data"

var resultFile = fmt.Sprintf("%s/%s", dataDir, "gurume.processed.1.txt")
var fileName string

var formatDataCmd = &cobra.Command{
	Use:   "formatData",
	Short: "format data..",
	Long:  "format data...",
	Run:   formatDataExecute,
}

func init() {

	// fileName
	formatDataCmd.Flags().StringVarP(&fileName, "file", "f", "", "Source file to read")
	formatDataCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(formatDataCmd)
}

func formatDataExecute(cmd *cobra.Command, args []string) {
	var fp *os.File
	var err error

	// read file
	fullFilePath := fmt.Sprintf("%s/%s", dataDir, fileName)
	logger.Infof("read file: %s", fullFilePath)
	fp, err = os.Open(fullFilePath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)

	// process start
	var category string
	var gurumeList []string
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

					// logger.Infof("%s, %s - N/A - N/A - %s", category, subCategory, name)
					gurume := fmt.Sprintf("%s, %s - N/A - N/A - %s", strings.Replace(category, "hotel, ", "", 1), subCategory, name)
					gurumeList = append(gurumeList, gurume)

				} else if strings.Contains(category, "노포 식당") {
					// town - name - since, town - name - since - note
					nopo := strings.Split(str, "-")
					town := strings.TrimSpace(nopo[0])
					name := strings.TrimSpace(nopo[1])
					since := strings.TrimSpace(nopo[2])

					if len(nopo) == 3 {
						// logger.Infof("%s - %s - N/A - %s (since %s)", category, town, name, since)
						gurume := fmt.Sprintf("%s - %s - N/A - %s (since %s)", category, town, name, since)
						gurumeList = append(gurumeList, gurume)
					} else if len(nopo) == 4 {
						note := strings.TrimSpace(nopo[3])
						// logger.Infof("%s - %s - N/A - %s (since %s) - %s", category, town, name, since, note)
						gurume := fmt.Sprintf("%s - %s - N/A - %s (since %s) - %s", category, town, name, since, note)
						gurumeList = append(gurumeList, gurume)
					} else {
						// logger.Infof("exception: %s - %s", category, str)
						gurume := fmt.Sprintf("exception: %s - %s", category, str)
						gurumeList = append(gurumeList, gurume)
					}

				} else {
					cnt := strings.Count(str, "-")

					if cnt == 0 {
						// restraunt only
						// logger.Infof("%s - N/A - N/A - %s", category, str)
						gurume := fmt.Sprintf("%s - N/A - N/A - %s", category, str)
						gurumeList = append(gurumeList, gurume)
					} else if cnt == 2 || cnt == 3 {
						// town - station - restraunt, town - station - restraunt - note case
						// logger.Infof("%s - %s", category, str)
						gurume := fmt.Sprintf("%s - %s", category, str)
						gurumeList = append(gurumeList, gurume)
					} else {
						// exception
						// logger.Infof("exception: %s - %s", category, str)
						gurume := fmt.Sprintf("exception: %s - %s", category, str)
						gurumeList = append(gurumeList, gurume)
					}

					// output:
					// category - town - station - restraunt - note
					// category - town - station - restraunt
				}
			}
		}

		// EOF(end of file)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	} // file read done

	writeGurumeList(gurumeList)
}

// write gurume file
func writeGurumeList(gurumeList []string) {
	f, err := os.Create(resultFile)
	if err != nil {
		logger.WithError(err).Error("error")
		f.Close()
		return
	}

	for _, gurume := range gurumeList {
		fmt.Fprintln(f, gurume)
		if err != nil {
			logger.WithError(err).Error("error")
			return
		}
	}

	err = f.Close()
	if err != nil {
		logger.WithError(err).Error("error")
		return
	}
	logger.Infof("file written successfully: %s", resultFile)
}
