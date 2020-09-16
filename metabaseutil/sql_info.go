package metabaseutil

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/grokify/gocharts/data/statictimeseries"
)

// SQLInfo holds the information for a SQL query.
type SQLInfo struct {
	Name        string
	DatabaseID  int64
	SQL         string
	ColIdxCount int
	ColIdxDate  int
}

// Validate checks the parameter information before the
// query is performed.
func (sqli *SQLInfo) Validate(checkColUnique bool) error {
	sqli.SQL = strings.TrimSpace(sqli.SQL)
	if len(sqli.SQL) == 0 {
		return errors.New("E_SQL_QUERY_NOT_PRESENT")
	}
	if checkColUnique && sqli.ColIdxCount == sqli.ColIdxDate {
		return errors.New("E_COLIDX_COUNT_AND_DATE_SAME")
	}
	return nil
}

// QuerySTS executes a raw SQL query that is designed to provide
// counts by date.
func QuerySTS(httpClient *http.Client, baseURL string, opts SQLInfo) (statictimeseries.DataSeries, *SqlResponse, error) {
	ds := statictimeseries.NewDataSeries()
	err := opts.Validate(true)
	if err != nil {
		return ds, nil, err
	}
	resp, err := QuerySQLHttp(httpClient, baseURL, opts.DatabaseID, opts.SQL)
	if err != nil {
		return ds, nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ds, nil, err
	}
	sqlResponse, err := NewSqlResponse(bytes)
	if err != nil {
		return ds, sqlResponse, err
	}
	sts, err := SqlResponseToSTS(opts.Name, sqlResponse, opts.ColIdxCount, opts.ColIdxDate)
	if err != nil {
		return ds, sqlResponse, err
	}
	return sts, sqlResponse, nil
}
