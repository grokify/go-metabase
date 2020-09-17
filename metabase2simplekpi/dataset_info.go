package metabase2simplekpi

import (
	"strings"

	"github.com/grokify/go-metabase/metabaseutil"
)

// DatasetInfo captures information for a single SQL query representing
// data for a single KPI.
type DatasetInfo struct {
	KpiName             string               `json:"kpiName"`
	MetricName          string               `json:"metricName"`
	MetabaseQueryExec   bool                 `json:"metabaseExec"`
	MetabaseQuery       metabaseutil.SQLInfo `json:"metabaseQuery"`
	SimplekpiUpdateExec bool                 `json:"simplekpiExec"`
	SimplekpiKpiId      int                  `json:"simplekpiKpiId"`
}

func (dsi *DatasetInfo) Inflate() {
	dsi.KpiName = strings.TrimSpace(dsi.KpiName)
	dsi.MetabaseQuery.Inflate()
	dsi.MetabaseQuery.Name = strings.TrimSpace(dsi.MetabaseQuery.Name)
	if len(dsi.KpiName) > 0 && len(dsi.MetabaseQuery.Name) == 0 {
		dsi.MetabaseQuery.Name = dsi.KpiName
	}
}
