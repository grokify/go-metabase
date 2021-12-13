package metabaseutil

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/grokify/go-metabase/metabase"
	"github.com/grokify/gocharts/data/timeseries"
	"github.com/grokify/mogo/net/httputilmore"
	"github.com/grokify/mogo/net/urlutil"
)

const (
	ApiUrlDataset   = "/api/dataset"
	QueryTypeNative = "native"
)

func QuerySQL(apiClient *metabase.APIClient, databaseId int64, sql string) (metabase.DatasetQueryResults, *http.Response, error) {
	return apiClient.DatasetApi.QueryDatabase(
		context.Background(),
		metabase.DatasetQueryJsonQuery{
			Database: databaseId,
			Type:     QueryTypeNative,
			Native:   metabase.DatasetQueryNative{Query: sql},
			// Constraints: metabase.DatasetQueryConstraints{MaxResults: limit},
		})
}

func QuerySQLHttpMore(httpClient *http.Client, baseUrl string, databaseId int64, sql string) (*SqlResponse, *http.Response, error) {
	resp, err := QuerySQLHttp(httpClient, baseUrl, databaseId, sql)
	if err != nil {
		return nil, resp, err
	}
	sqlr, err := NewSqlResponseHttp(resp)
	return sqlr, resp, err
}

func QuerySQLHttp(httpClient *http.Client, baseUrl string, databaseId int64, sql string) (*http.Response, error) {
	apiUrl := urlutil.JoinAbsolute(baseUrl, ApiUrlDataset)

	qry := metabase.DatasetQueryJsonQuery{
		Database: databaseId,
		Type:     "native",
		Native:   metabase.DatasetQueryNative{Query: sql},
		// Constraints: metabase.DatasetQueryConstraints{MaxResults: limit},
	}

	qryJSON, err := json.Marshal(qry)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Post(
		apiUrl,
		httputilmore.ContentTypeAppJsonUtf8,
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
