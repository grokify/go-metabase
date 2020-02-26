package metabase2simplekpi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/grokify/gocharts/data/statictimeseries"
	"github.com/grokify/gocharts/data/table"
	"github.com/grokify/gotilla/time/timeutil"
	"github.com/pkg/errors"
)

func HTTPResponseToSqlResponse(resp *http.Response) (*SqlResponse, error) {
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "E_METABASE_API_READ_BODY_ERROR")
	} else if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("E_METABASE_API_ERROR STATUS[%v] MSG[%v]",
			resp.StatusCode, strings.TrimSpace(string(bytes)))
	}
	fmt.Println(string(bytes))
	sr := &SqlResponse{}
	err = json.Unmarshal(bytes, sr)
	if err != nil {
		err = errors.Wrap(err, "E_METABASE_API_RESPONSE_JSON_DECODE")
	}
	return sr, err
}

type SqlResponse struct {
	Data SqlData `json:"data,omitempty"`
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

func SqlResponseToTable(sr *SqlResponse, kpiId int64, countColIdx, dateColIdx int) (table.TableData, error) {
	tbl := table.NewTableData()
	cols := []string{"KPI ID", "Date", "Actual"}
	rows := [][]string{}
	for _, row := range sr.Data.Rows {
		count := int(row[countColIdx].(float64))
		dt, err := time.Parse(time.RFC3339, row[dateColIdx].(string))
		if err != nil {
			return tbl, err
		}
		fmt.Printf("Count [%d] TIME [%v]\n", count, dt.Format(timeutil.RFC3339FullDate))
		row := []string{
			strconv.Itoa(int(kpiId)),
			dt.Format(timeutil.RFC3339FullDate),
			strconv.Itoa(count)}
		rows = append(rows, row)
	}
	tbl.Columns = cols
	tbl.Records = rows
	return tbl, nil
}
