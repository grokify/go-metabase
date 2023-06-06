package metabaseutil

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/grokify/go-metabase/metabase"
	"github.com/grokify/gocharts/v2/data/timeseries"
	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/grokify/mogo/net/urlutil"
)

const (
	ApiUrlDataset   = "/api/dataset"
	QueryTypeNative = "native"
)

func QuerySQL(apiClient *metabase.APIClient, databaseID int64, sql string) (*metabase.DatasetQueryResults, *http.Response, error) {
	request := apiClient.DatasetApi.QueryDatabase(context.Background())
	qtn := QueryTypeNative

	request.DatasetQueryJsonQuery(metabase.DatasetQueryJsonQuery{
		Database: &databaseID,
		Type:     &qtn,
		Native:   &metabase.DatasetQueryNative{Query: &sql},
		// Constraints: metabase.DatasetQueryConstraints{MaxResults: limit},
	})

	return apiClient.DatasetApi.QueryDatabaseExecute(request)
}

func QuerySQLHttpMore(httpClient *http.Client, baseURL string, databaseID int64, sql string) (*SqlResponse, *http.Response, error) {
	resp, err := QuerySQLHttp(httpClient, baseURL, databaseID, sql)
	if err != nil {
		return nil, resp, err
	}
	sqlr, err := NewSqlResponseHttp(resp)
	return sqlr, resp, err
}

func QuerySQLHttp(httpClient *http.Client, baseURL string, databaseID int64, sql string) (*http.Response, error) {
	apiUrl := urlutil.JoinAbsolute(baseURL, ApiUrlDataset)
	qtn := "native"

	qry := metabase.DatasetQueryJsonQuery{
		Database: &databaseID,
		Type:     &qtn,
		Native:   &metabase.DatasetQueryNative{Query: &sql},
		// Constraints: metabase.DatasetQueryConstraints{MaxResults: limit},
	}

	qryJSON, err := json.Marshal(qry)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Post(
		apiUrl,
		httputilmore.ContentTypeAppJSONUtf8,
		bytes.NewBuffer(qryJSON))
	return resp, err
}

type SqlData struct {
	Rows [][]interface{} `json:"rows,omitempty"`
}

type SqlResponse struct {
	Data SqlData `json:"data,omitempty"`
}

func NewSqlResponseHttp(resp *http.Response) (*SqlResponse, error) {
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return NewSqlResponse(bytes)
}

func NewSqlResponse(bytes []byte) (*SqlResponse, error) {
	sr := &SqlResponse{}
	err := json.Unmarshal(bytes, sr)
	return sr, err
}

func SqlResponseToTimeSeries(seriesName string, sr *SqlResponse, countColIdx, timeColIdx int) (timeseries.TimeSeries, error) {
	ts := timeseries.NewTimeSeries(seriesName)
	for _, row := range sr.Data.Rows {
		count := row[countColIdx].(float64)
		dtStr := row[timeColIdx].(string)
		dt, err := time.Parse(time.RFC3339, dtStr)
		if err != nil {
			return ts, err
		}
		ts.AddItems(timeseries.TimeItem{
			SeriesName: seriesName,
			Value:      int64(count),
			Time:       dt})
	}
	return ts, nil
}
