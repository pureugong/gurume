package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/pureugong/gurume/config"
	"github.com/pureugong/gurume/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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

	// bind es information ENV
	viper.BindEnv("ES_CLUSTER_HOST")
	viper.BindEnv("ES_CLUSTER_PORT")
	viper.BindEnv("ES_CLUSTER_USER_ID")
	viper.BindEnv("ES_CLUSTER_USER_PW")

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
	var gurumeList []*model.Gurume
	for {
		// read
		line, _, err := reader.ReadLine()
		if string(line) != "" {
			gurume := &model.Gurume{}
			_ = json.Unmarshal(line, &gurume)
			logger.Info(gurume)
			gurumeList = append(gurumeList, gurume)
		}

		// EOF(end of file)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

	} // file read done

	// ES client
	client := config.NewESClient()
	defer client.Stop()

	// ES client - create index
	ctx := context.Background()
	err = createIndex(ctx, client, "gurume_index")
	if err != nil {
		panic(err)
	}

	// ES client - bulk request
	bulkRequest := client.Bulk()
	for _, gurume := range gurumeList {
		indexReq := elastic.NewBulkIndexRequest().Index("gurume_index").Type("gurume").Doc(gurume)
		bulkRequest = bulkRequest.Add(indexReq)
	}

	// ES client - bulk response
	bulkResponse, err := bulkRequest.Do(ctx)
	if err != nil {
		panic(err)
	}

	// ES client - bulk response check
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

// TODO: add "user_dictionary": "userdict_ko.txt"
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
						"decompound_mode": "mixed"
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
					"properties": {
						"name": {
							"type":"text",
							"analyzer": "nori_korean",
							"fields": {
								"keyword": {
									"type": "keyword",
									"ignore_above": 256
								}
							}
						}
					}
				},
				"station":{
					"properties": {
						"name": {
							"type":"text",
							"analyzer": "nori_korean",
							"fields": {
								"keyword": {
									"type": "keyword",
									"ignore_above": 256
								}
							}
						}
					}
				},
				"town":{
					"type":"text",
					"analyzer": "nori_korean",
					"fields": {
						"keyword": {
							"type": "keyword",
							"ignore_above": 256
						}
					}
				},
				"name":{
					"type":"text",
					"analyzer": "nori_korean",
					"fields": {
						"keyword": {
							"type": "keyword",
							"ignore_above": 256
						}
					}
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
