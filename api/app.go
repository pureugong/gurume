package api

import (
	"net/http"

	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

// App is
type App struct {
	Logger   *logrus.Logger
	ESClient *elastic.Client
}

// NewApp is
func NewApp(logger *logrus.Logger, esClient *elastic.Client) *App {
	return &App{
		Logger:   logger,
		ESClient: esClient,
	}
}

// Start is
func (a *App) Start(handler http.Handler) {
	http.ListenAndServe(":3000", handler)
}
