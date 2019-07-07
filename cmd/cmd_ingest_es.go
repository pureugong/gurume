package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	elastic "github.com/olivere/elastic"
)

var ingestElasticSearch = &cobra.Command{
	Use:   "ingestES",
	Short: "ingestES json..",
	Long:  "ingestES json...",
	Run:   ingestElasticSearchExecute,
}

func init() {
	rootCmd.AddCommand(ingestElasticSearch)
}

func ingestElasticSearchExecute(cmd *cobra.Command, args []string) {
	var fp *os.File
	var err error

	// read file
	fp, err = os.Open(resultJSONFile)
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
		if string(line) != "" {
			gurume := &Gurume{}
			_ = json.Unmarshal(line, &gurume)
			fmt.Println(gurume)
			gurumeList = append(gurumeList, gurume)
		}

		// EOF(end of file)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

	} // file read done

	// ES
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	defer client.Stop()

	ctx := context.Background()
	err = createIndex(ctx, client, "gurume_index")
	if err != nil {
		panic(err)
	}

	bulkRequest := client.Bulk()
	for _, gurume := range gurumeList {
		indexReq := elastic.NewBulkIndexRequest().Index("gurume_index").Type("gurume").Doc(gurume)
		bulkRequest = bulkRequest.Add(indexReq)
	}

	bulkResponse, err := bulkRequest.Do(ctx)
	if err != nil {
		panic(err)
	}

	indexed := bulkResponse.Indexed()
	if len(indexed) != len(gurumeList) {
		panic("there are missing gurume on es")
	}
}

func createIndex(ctx context.Context, client *elastic.Client, indexName string) error {

	// 1. check index exist and delete if there i
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
		// Handle error
	}

	if exists {

		deleteIndex, err := client.DeleteIndex(indexName).Do(ctx)
		if err != nil {
			// Handle error
			return err
		}
		if !deleteIndex.Acknowledged {
			// Not acknowledged
		}
	}

	// TODO: use alias
	createIndex, err := client.CreateIndex(indexName).BodyString(gurumeMapping).Do(ctx)
	if err != nil {
		// Handle error
		return err
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
	}

	return nil
}

const gurumeMapping = `
{
	"settings":{
		"number_of_shards": 5,
		"number_of_replicas": 1,
		"index": {
			"analysis": {
				"tokenizer": {
					"nori_user_dict": {
						"type": "nori_tokenizer",
						"decompound_mode": "mixed",
						"user_dictionary": "userdict_ko.txt"
					}
				},
				"analyzer": {
					"nori_korean":{
						"type": "custom",
						"tokenizer": "nori_user_dict"
					}
				}
			}
		}
	},
	"mappings":{
		"gurume":{
			"properties":{
				"category":{
					"type":"text",
					"analyzer": "nori_korean"
				},
				"station":{
					"type":"text"
				},
				"town":{
					"type":"text"
				},
				"name":{
					"type":"text",
					"analyzer": "nori_korean"
				},
				"note":{
					"type":"text",
					"analyzer": "nori_korean"
				}
			}
		}
	}
}
`
