package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pureugong/gurume/model"
	"github.com/spf13/cobra"
)

var formatJSONCmd = &cobra.Command{
	Use:   "formatJSON",
	Short: "format json..",
	Long:  "format json...",
	Run:   formatJSONCmdExecute,
}

func init() {
	rootCmd.AddCommand(formatJSONCmd)
}

func formatJSONCmdExecute(cmd *cobra.Command, args []string) {
	var fp *os.File
	var err error

	// read file
	fp, err = os.Open(resultFile)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)

	// process start
	var gurumeList []*model.Gurume
	for {
		// read
		line, _, err := reader.ReadLine()
		str := string(line)
		str = strings.TrimSpace(str)

		// skip empty line
		if str != "" {
			// TODO: encapsulate
			rawGurume := strings.Split(str, "-")

			if len(rawGurume) == 4 {

				gurumeList = append(
					gurumeList,
					model.NewGurume().
						SetCategory(rawGurume[0]).
						SetTown(rawGurume[1]).
						SetStation(rawGurume[2]).
						SetName(rawGurume[3]),
				)

			} else if len(rawGurume) == 5 {
				gurumeList = append(
					gurumeList,
					model.NewGurume().
						SetCategory(rawGurume[0]).
						SetTown(rawGurume[1]).
						SetStation(rawGurume[2]).
						SetName(rawGurume[3]).
						SetNote(rawGurume[4]),
				)
			}
		}

		// EOF(end of file)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	} // file read done

	// writeGurumeJSON
	writeGurumeJSON(gurumeList)
}

var resultJSONFile = fmt.Sprintf("%s/%s", dataDir, "gurume.processed.1.json")

// write gurume json file
func writeGurumeJSON(gurumeList []*model.Gurume) {
	f, err := os.Create(resultJSONFile)
	if err != nil {
		logger.WithError(err).Error("error")
		f.Close()
		return
	}

	for _, gurume := range gurumeList {
		bytes, _ := json.Marshal(gurume)
		fmt.Fprintln(f, string(bytes))
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
	logger.Infof("file written successfully: %s", resultJSONFile)
}
