package metabaseutil

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/grokify/go-metabase/metabase"
	"github.com/grokify/gocharts/data/statictimeseries"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/gotilla/net/urlutil"
)

const (
	ApiUrlDataset = "/api/dataset"
)

func QuerySQL(apiClient *metabase.APIClient, databaseId int64, sql string) (metabase.DatasetQueryResults, *http.Response, error) {
	opts := metabase.DatasetQueryJsonQuery{
		Database: databaseId,
		Type:     "native",
		Native: metabase.DatasetQueryNative{
			Query: sql},
		// Constraints: metabase.DatasetQueryConstraints{MaxResults: limit},
	}

	info, resp, err := apiClient.DatasetApi.QueryDatabase(
		context.Background(), opts)
	return info, resp, err
}

func QuerySQLHttp(httpClient *http.Client, baseUrl string, databaseId int64, sql string) (*http.Response, error) {
	apiUrl := urlutil.JoinAbsolute(baseUrl, ApiUrlDataset)

	qry := metabase.DatasetQueryJsonQuery{
		Database: databaseId,
		Type:     "native",
		Native: metabase.DatasetQueryNative{
			Query: sql},
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

type SqlResponse struct {
	Data SqlData `json:"data,omitempty"`
}

func NewSqlResponseHttp(resp *http.Response) (*SqlResponse, error) {
	bytes, err := ioutil.ReadAll(resp.Body)
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

type SqlData struct {
	Rows [][]interface{} `json:"rows,omitempty"`
}

func SqlResponseToSTS(seriesName string, sr *SqlResponse, countColIdx, dateColIdx int) (statictimeseries.DataSeries, error) {
	sts := statictimeseries.NewDataSeries()
	sts.SeriesName = seriesName
	for _, row := range sr.Data.Rows {
		count := row[countColIdx].(float64)
		dtStr := row[dateColIdx].(string)
		dt, err := time.Parse(time.RFC3339, dtStr)
		if err != nil {
			return sts, err
		}
		item := statictimeseries.DataItem{
			SeriesName: seriesName,
			Value:      int64(count),
			Time:       dt}
		sts.AddItem(item)
	}
	return sts, nil
}
