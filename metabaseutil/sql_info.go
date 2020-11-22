package metabaseutil

import (
	"errors"
	"fmt"
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
	SQLFormat   string
	SQLVars     []interface{}
	ColIdxCount int
	ColIdxTime  int
}

// Inflate expands data.
func (sqli *SQLInfo) Inflate() {
	sqli.SQL = sqli.NativeSQL()
}

// NativeSQL returns a formatted or raw SQL statement.
func (sqli *SQLInfo) NativeSQL() string {
	sqli.SQL = strings.TrimSpace(sqli.SQL)
	sqli.SQLFormat = strings.TrimSpace(sqli.SQLFormat)
	if len(sqli.SQLFormat) > 0 {
		if len(sqli.SQLVars) > 0 {
			return fmt.Sprintf(sqli.SQLFormat, sqli.SQLVars...)
		} else if len(sqli.SQL) == 0 {
			return sqli.SQLFormat
		}
	}
	return sqli.SQL
}

// Validate checks the parameter information before the
// query is performed.
func (sqli *SQLInfo) Validate(checkColUnique bool) error {
	sqli.SQL = strings.TrimSpace(sqli.SQL)
	if len(sqli.SQL) == 0 {
		return errors.New("E_SQL_QUERY_NOT_PRESENT")
	}
	if checkColUnique && sqli.ColIdxCount == sqli.ColIdxTime {
		return errors.New("E_COLIDX_COUNT_AND_DATE_SAME")
	}
	return nil
}

// QueryDataSeries executes a raw SQL query that is designed to provide
// counts by date.
func QueryDataSeries(httpClient *http.Client, baseURL string, opts SQLInfo) (*statictimeseries.DataSeries, *SqlResponse, error) {
	err := opts.Validate(true)
	if err != nil {
		return nil, nil, err
	}
	resp, err := QuerySQLHttp(httpClient, baseURL, opts.DatabaseID, opts.SQL)
	if err != nil {
		return nil, nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	sqlResponse, err := NewSqlResponse(bytes)
	if err != nil {
		return nil, sqlResponse, err
	}
	sts, err := SqlResponseToSTS(opts.Name, sqlResponse, opts.ColIdxCount, opts.ColIdxTime)
	if err != nil {
		return nil, sqlResponse, err
	}
	return &sts, sqlResponse, nil
}
