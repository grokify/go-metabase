package metabase2simplekpi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/grokify/gocharts/data/statictimeseries"
)

func HTTPResponseToSqlResponse(resp *http.Response) (*SqlResponse, error) {
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	sr := &SqlResponse{}
	err = json.Unmarshal(bytes, sr)
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
