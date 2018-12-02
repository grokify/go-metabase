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
	Cols []string
	Rows [][]string
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
		if len(records.Cols) == 0 {
			records.Cols = info.Data.Columns
		}
		records.Rows = append(records.Rows, info.Data.Rows...)
		opts.Query.Page.Page += int64(1)
		pageCount = int64(len(info.Data.Rows))
	}
	return records, nil
}
