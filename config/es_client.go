package config

import (
	"fmt"
	"net/http"

	"github.com/olivere/elastic"
	"github.com/spf13/viper"
)

func init() {
	viper.BindEnv("ES_CLUSTER_HOST")
	viper.BindEnv("ES_CLUSTER_PORT")
	viper.BindEnv("ES_CLUSTER_USER_ID")
	viper.BindEnv("ES_CLUSTER_USER_PW")
}

// NewESClient is
func NewESClient() *elastic.Client {
	// ES client
	httpClient := &http.Client{
		Transport: &BasicAuthTransport{
			username: viper.GetString("ES_CLUSTER_USER_ID"),
			password: viper.GetString("ES_CLUSTER_USER_PW"),
		},
	}
	elasticURL := fmt.Sprintf("%s:%s", viper.GetString("ES_CLUSTER_HOST"), viper.GetString("ES_CLUSTER_PORT"))
	client, err := elastic.NewClient(
		elastic.SetURL(elasticURL),
		elastic.SetHttpClient(httpClient),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	return client
}

// BasicAuthTransport is to store username, password for bearer header
type BasicAuthTransport struct {
	username string
	password string
}

// RoundTrip is to add bearer header
func (tr *BasicAuthTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(tr.username, tr.password)
	return http.DefaultTransport.RoundTrip(r)
}
