package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/grokify/go-metabase/metabase"
	mo "github.com/grokify/oauth2more/metabase"
)

const MaxPerPage = int64(2000)

func NewApiClientEnv(cfg mo.InitConfig) (*metabase.APIClient, *mo.AuthResponse, error) {
	httpClient, authResponse, err := mo.NewClientEnv(cfg)
	if err != nil {
		return nil, authResponse, err
	}
	apiConfig := metabase.NewConfiguration()
	apiConfig.BasePath = os.Getenv(cfg.EnvMetabaseBaseUrl)
	apiConfig.HTTPClient = httpClient
	apiClient := metabase.NewAPIClient(apiConfig)

	return apiClient, authResponse, nil
}

type Records struct {
	Columns []metabase.DatasetQueryResultsMetadataColumn
	Cols    []string
	Rows    [][]interface{}
}

func NewRecordsFromJsonFile(file string) (Records, error) {
	recs := Records{}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return recs, err
	}
	err = json.Unmarshal(bytes, &recs)
	return recs, err
}

func (recs *Records) InflateRecord(rec []interface{}) Record {
	return Record{
		Columns: recs.Columns,
		Cols:    recs.Cols,
		Row:     rec}
}

type Record struct {
	Columns []metabase.DatasetQueryResultsMetadataColumn
	Cols    []string
	Row     []interface{}
}

func (rec *Record) GetStringOrEmpty(key string) string {
	str, err := rec.GetString(key)
	if err != nil {
		return ""
	}
	return str
}

func (rec *Record) GetString(key string) (string, error) {
	index := -1
	for i, try := range rec.Cols {
		if key == try {
			index = i
			break
		}
	}
	if index >= 0 && index < len(rec.Row) {
		rawVal := rec.Row[index]
		rawStr, ok := rawVal.(string)
		if !ok {
			return rawStr, fmt.Errorf("Cannot convert to string [%v]", rawVal)
		}
		return rawStr, nil
	} else {
		return "", fmt.Errorf("Cannot find column for key [%v]", key)
	}
}

func (rec *Record) GetTime(key, format string) (time.Time, error) {
	index := -1
	for i, try := range rec.Cols {
		if key == try {
			index = i
			break
		}
	}
	if index >= 0 && index < len(rec.Row) {
		rawVal := rec.Row[index]
		rawStr, ok := rawVal.(string)
		if !ok {
			return time.Now(), fmt.Errorf("Cannot convert to string [%v]", rawVal)
		}
		return time.Parse(format, rawStr)
	} else {
		return time.Now(), fmt.Errorf("Cannot find column for key [%v]", key)
	}
}

func SimpleQuery(databaseId, tableId int64) metabase.DatasetQueryJsonQuery {
	return metabase.DatasetQueryJsonQuery{
		Database: databaseId,
		Type:     "query",
		Query: metabase.DatasetQueryDsl{
			SourceTable: tableId,
			Page:        metabase.DatasetQueryDslPage{Page: int64(1), Items: MaxPerPage}}}
}

func GetAllRecordsSimple(apiClient *metabase.APIClient, databaseId, tableId int64) (Records, error) {
	return GetAllRecords(apiClient, SimpleQuery(databaseId, tableId))
}

func GetAllRecords(apiClient *metabase.APIClient, opts metabase.DatasetQueryJsonQuery) (Records, error) {
	opts.Query.Page = metabase.DatasetQueryDslPage{Page: int64(1), Items: MaxPerPage}
	records := Records{}
	pageCount := MaxPerPage

	for pageCount >= MaxPerPage {
		info, resp, err := apiClient.DatasetApi.QueryDatabase(
			context.Background(), opts)
		if err != nil {
			return records, err
		} else if resp.StatusCode >= 300 {
			return records, fmt.Errorf("Metabase API Response Status [%v]", resp.StatusCode)
		}
		if len(records.Columns) == 0 {
			records.Columns = info.Data.ResultsMetadata.Columns
		}
		if len(records.Cols) == 0 {
			records.Cols = info.Data.Columns
		}
		records.Rows = append(records.Rows, info.Data.Rows...)
		opts.Query.Page.Page += int64(1)
		pageCount = int64(len(info.Data.Rows))
	}
	return records, nil
}
