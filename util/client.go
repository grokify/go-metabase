package util

import (
	"context"
	"fmt"
	"os"

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

type Record struct {
	Columns []metabase.DatasetQueryResultsMetadataColumn
	Cols    []string
	Row     [][]interface{}
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
