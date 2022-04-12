package metabaseutil

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/grokify/go-metabase/metabase"
	mo "github.com/grokify/goauth/metabase"
)

const MaxPerPage = int64(2000)

func NewApiClient(cfg mo.Config) (*metabase.APIClient, *mo.AuthResponse, error) {
	httpClient, authResponse, err := mo.NewClient(cfg)
	if err != nil {
		return nil, authResponse, err
	}
	apiClient := NewApiClientHttpClient(cfg.BaseURL, httpClient)
	return apiClient, authResponse, nil
}

func NewApiClientPasswordWithSessionId(serverURL, username, password, sessionId string, tlsSkipVerify bool) (*metabase.APIClient, *mo.AuthResponse, error) {
	httpClient, authResponse, err := mo.NewClientPasswordWithSessionID(
		serverURL,
		username,
		password,
		sessionId,
		tlsSkipVerify)
	if err != nil {
		return nil, authResponse, err
	}

	apiConfig := metabase.NewConfiguration()
	apiConfig.BasePath = serverURL
	apiConfig.HTTPClient = httpClient
	return metabase.NewAPIClient(apiConfig), authResponse, nil
}

func NewApiClientEnv(cfg *mo.ConfigEnvOpts) (*metabase.APIClient, *mo.AuthResponse, error) {
	httpClient, authResponse, _, err := mo.NewClientEnv(cfg)
	if err != nil {
		return nil, authResponse, err
	}

	apiClient := NewApiClientHttpClient(os.Getenv(cfg.EnvMetabaseBaseURL), httpClient)
	return apiClient, authResponse, nil
}

func NewApiClientHttpClient(mbBaseUrl string, httpClient *http.Client) *metabase.APIClient {
	apiConfig := metabase.NewConfiguration()
	apiConfig.BasePath = mbBaseUrl
	apiConfig.HTTPClient = httpClient
	return metabase.NewAPIClient(apiConfig)
}

func NewApiClientSessionId(serverUrl, token string, tlsSkipVerify bool) *metabase.APIClient {
	return metabase.NewAPIClient(
		&metabase.Configuration{
			BasePath:   serverUrl,
			HTTPClient: mo.NewClientSessionID(token, tlsSkipVerify)})
}

type Records struct {
	Columns []metabase.DatasetQueryResultsMetadataColumn
	Cols    []string
	Rows    [][]interface{}
	//Records []Record
}

func (recs *Records) RecordsSet() RecordsSet {
	rs := RecordsSet{Records: []Record{}}
	for _, r := range recs.Rows {
		rs.Records = append(rs.Records, recs.InflateRecord(r))
	}
	return rs
}

func (recs *Records) FilterRecordsString(key, val string) RecordsSet {
	recs2 := RecordsSet{Records: []Record{}}
	for _, r := range recs.Rows {
		r2 := recs.InflateRecord(r)
		valRec := r2.GetStringOrEmpty(key)
		if valRec == val {
			recs2.Records = append(recs2.Records, r2)
		}
	}
	return recs2
}

func (recs *Records) FilterRecordsStringOr(key string, vals []string) RecordsSet {
	recs2 := RecordsSet{Records: []Record{}}
	for _, r := range recs.Rows {
		r2 := recs.InflateRecord(r)
		valRec := r2.GetStringOrEmpty(key)
		for _, val := range vals {
			if valRec == val {
				recs2.Records = append(recs2.Records, r2)
			}
		}
	}
	return recs2
}

func NewRecordsFromJsonFile(file string) (Records, error) {
	recs := Records{}
	bytes, err := os.ReadFile(file)
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

type RecordsSet struct {
	Records []Record
}

func (recs *RecordsSet) FilterRecordsString(key, val string) RecordsSet {
	recs2 := RecordsSet{Records: []Record{}}
	for _, rec := range recs.Records {
		valTry := rec.GetStringOrEmpty(key)
		if valTry == val {
			recs2.Records = append(recs2.Records, rec)
		}
	}
	return recs2
}

func (recs *RecordsSet) Distinct(key string) RecordsSet {
	recs2 := RecordsSet{Records: []Record{}}
	seen := map[string]int{}
	for _, rec := range recs.Records {
		valTry := rec.GetStringOrEmpty(key)
		if _, ok := seen[valTry]; !ok {
			recs2.Records = append(recs2.Records, rec)
			seen[valTry] = 0
		}
		seen[valTry] += 1
	}
	return recs2
}

type Record struct {
	Columns []metabase.DatasetQueryResultsMetadataColumn
	Cols    []string
	Row     []interface{}
	Related []Record
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
