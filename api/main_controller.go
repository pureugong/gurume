package api

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/olivere/elastic"
	"github.com/pureugong/gurume/model"
	"github.com/sirupsen/logrus"
)

// MainController is
type MainController struct {
	Logger   *logrus.Logger
	ESClient *elastic.Client
}

// NewMainController is
func NewMainController(app *App) *MainController {
	return &MainController{
		Logger:   app.Logger,
		ESClient: app.ESClient,
	}
}

// Router is
func (m *MainController) Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/gurume", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Request Param
		terms := r.URL.Query().Get("search")

		// ES search

		q := elastic.NewMultiMatchQuery(
			terms,
			"category.name",
			"town",
			"station.name",
			"name",
			"note",
		).Type("cross_fields")
		// .Operator("and")

		// q := elastic.NewBoolQuery()
		// q.Must(elastic.NewTermQuery("station.name", terms))

		searchResult, err := m.ESClient.Search().
			Index("gurume_index"). // search in index
			Query(q).              // specify the query
			From(0).Size(100).     // take documents 0-9
			Do(ctx)                // execute
		if err != nil {
			m.Logger.WithError(err).Error("es client search fail")
			panic(err)
		}
		m.Logger.Debugf("Query took %d milliseconds", searchResult.TookInMillis)

		// API Response
		gurumeList := make([]*model.Gurume, 0)
		var gurume model.Gurume
		for _, item := range searchResult.Each(reflect.TypeOf(gurume)) {
			if g, ok := item.(model.Gurume); ok {
				gurumeList = append(gurumeList, &g)
			}
		}

		m.Logger.Debugf("Found a total of %d gurume", searchResult.TotalHits())

		w.Header().Set("Content-Type", "application/json")
		response := New200GurumeResponse(searchResult.TotalHits(), gurumeList)
		result, _ := json.Marshal(response)
		w.Write(result)
	})

	return r
}

// GurumeResponse is
type GurumeResponse struct {
	Status string  `json:"status"`
	Result *Result `json:"result"`
}

// Result is
type Result struct {
	Found   int64           `json:"found"`
	Gurumes []*model.Gurume `json:"gurumes"`
}

// New200GurumeResponse is
func New200GurumeResponse(found int64, gurumes []*model.Gurume) *GurumeResponse {
	return &GurumeResponse{
		Status: "ok",
		Result: &Result{found, gurumes},
	}
}
