package metabase2simplekpi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/grokify/gocharts/data/table"
	"github.com/grokify/gocharts/data/timeseries"
	"github.com/grokify/simplego/time/timeutil"
	"github.com/pkg/errors"
)

// HTTPResponseToSqlResponse parses a Metabase SQL response.
func HTTPResponseToSqlResponse(resp *http.Response) (*SqlResponse, error) {
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "E_METABASE_API_READ_BODY_ERROR")
	} else if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("E_METABASE_API_ERROR STATUS[%v] MSG[%v]",
			resp.StatusCode, strings.TrimSpace(string(bytes)))
	}

	sr := &SqlResponse{}
	err = json.Unmarshal(bytes, sr)
	if err != nil {
		err = errors.Wrap(err, "E_METABASE_API_RESPONSE_JSON_DECODE")
	} else if len(strings.TrimSpace(sr.Error)) > 0 {
		sr.Error = strings.TrimSpace(sr.Error)
		sr.Status = strings.TrimSpace(sr.Status)
		if len(sr.Status) > 0 {
			err = fmt.Errorf("E_METABASE_API_STATUS [%v] API_ERROR [%v]", sr.Status, sr.Error)
		} else {
			err = fmt.Errorf("E_METABASE_API_ERROR [%v]", sr.Error)
		}
	}
	return sr, err
}

// SqlResponse represents a Metabase API SQL response body.
type SqlResponse struct {
	Context     string    `json:"context,omitempty"`
	Data        SqlData   `json:"data,omitempty"`
	Error       string    `json:"error,omitempty"`
	JSONQuery   JSONQuery `json:"json_query"`
	RowCount    int       `json:"row_count"`
	RunningTime int       `json:"running_time"`
	StartedAt   time.Time `json:"started_at"`
	Status      string    `json:"status,omitempty"`
}

type SqlData struct {
	Rows [][]interface{} `json:"rows,omitempty"`
}

type JSONQuery struct {
	Database    int         `json:"database"`
	Type        string      `json:"type"`
	Constraints Constraints `json:"constraints"`
}

type Constraints struct {
	MaxResults         int `json:"max-results"`
	MaxResultsBareRows int `json:"max-results-bare-rows"`
}

func SqlResponseToTS(seriesName string, sr *SqlResponse, countColIdx, dateColIdx int) (timeseries.TimeSeries, error) {
	ts := timeseries.NewTimeSeries(seriesName)
	for _, row := range sr.Data.Rows {
		count := row[countColIdx].(float64)
		dtStr := row[dateColIdx].(string)
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

func SqlResponseToTable(sr *SqlResponse, kpiId int64, countColIdx, dateColIdx int) (table.Table, error) {
	tbl := table.NewTable("")
	cols := []string{"KPI ID", "Date", "Actual"}
	rows := [][]string{}
	for _, row := range sr.Data.Rows {
		count := int(row[countColIdx].(float64))
		dt, err := time.Parse(time.RFC3339, row[dateColIdx].(string))
		if err != nil {
			return tbl, err
		}
		row := []string{
			strconv.Itoa(int(kpiId)),
			dt.Format(timeutil.RFC3339FullDate),
			strconv.Itoa(count)}
		rows = append(rows, row)
	}
	tbl.Columns = cols
	tbl.Rows = rows
	return tbl, nil
}
