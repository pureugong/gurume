package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

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
	var gurumeList []*Gurume
	for {
		// read
		line, _, err := reader.ReadLine()
		str := string(line)
		str = strings.TrimSpace(str)

		// skip empty line
		if str != "" {
			// TODO: encapsulate
			rawGurume := strings.Split(str, "-")

			for index := 0; index < len(rawGurume); index++ {
				rawGurume[index] = strings.TrimSpace(rawGurume[index])
				if rawGurume[index] == "N/A" {
					rawGurume[index] = ""
				}
			}

			if len(rawGurume) == 4 {
				gurumeList = append(gurumeList, NewGurume(
					rawGurume[0],
					rawGurume[1],
					rawGurume[2],
					rawGurume[3],
					"",
				))
			} else if len(rawGurume) == 5 {
				gurumeList = append(gurumeList, NewGurume(
					rawGurume[0],
					rawGurume[1],
					rawGurume[2],
					rawGurume[3],
					rawGurume[4],
				))
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

// Gurume data type
type Gurume struct {
	Category string `json:"category"`
	Station  string `json:"station,omitempty"`
	Town     string `json:"town,omitempty"`
	Name     string `json:"name"`
	Note     string `json:"note,omitempty"`
}

// NewGurume is to init gurume
func NewGurume(category, town, station, name, note string) *Gurume {
	return &Gurume{
		Category: category,
		Station:  station,
		Town:     town,
		Name:     name,
		Note:     note,
	}
}

var resultJSONFile = fmt.Sprintf("%s/%s", dataDir, "gurume.processed.1.json")

// write gurume json file
func writeGurumeJSON(gurumeList []*Gurume) {
	f, err := os.Create(resultJSONFile)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	for _, gurume := range gurumeList {
		bytes, _ := json.Marshal(gurume)
		fmt.Fprintln(f, string(bytes))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("file written successfully: %s\n", resultJSONFile)
}
