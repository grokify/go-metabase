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
	/*
		MetabaseQueryDatabaseId      int                  `json:"metabaseQueryDatabaseId"`
		MetabaseQueryNativeSQL       string               `json:"metabaseQueryNativeSQL"`
		MetabaseQueryNativeSQLFormat string               `json:"metabaseQueryNativeSQLFormat"`
		MetabaseQueryNativeSQLVars   []interface{}        `json:"metabaseQueryNativeSQLVars"`
		MetabaseQueryColIdxCount     int                  `json:"metabaseQueryColIdxCount"`
		MetabaseQueryColIdxDate      int                  `json:"metabaseQueryColIdxDate"`
	*/
}

func (dsi *DatasetInfo) Inflate() {
	dsi.KpiName = strings.TrimSpace(dsi.KpiName)
	dsi.MetabaseQuery.Inflate()
	dsi.MetabaseQuery.Name = strings.TrimSpace(dsi.MetabaseQuery.Name)
	if len(dsi.KpiName) > 0 && len(dsi.MetabaseQuery.Name) == 0 {
		dsi.MetabaseQuery.Name = dsi.KpiName
	}
	//	dsi.MetabaseQueryNativeSQL = dsi.NativeSQL()
}

/*
// NativeSQL returns a formatted or raw SQL statement.
func (dsi *DatasetInfo) NativeSQL() string {
	dsi.MetabaseQueryNativeSQL = strings.TrimSpace(dsi.MetabaseQueryNativeSQL)
	dsi.MetabaseQueryNativeSQLFormat = strings.TrimSpace(dsi.MetabaseQueryNativeSQLFormat)
	if len(dsi.MetabaseQueryNativeSQLFormat) > 0 {
		if len(dsi.MetabaseQueryNativeSQLVars) > 0 {
			nativeSQL := fmt.Sprintf(
				dsi.MetabaseQueryNativeSQLFormat,
				dsi.MetabaseQueryNativeSQLVars...)
			dsi.MetabaseQueryNativeSQL = nativeSQL
			return nativeSQL
		}
		return dsi.MetabaseQueryNativeSQLFormat
	}
	return dsi.MetabaseQueryNativeSQL
}
*/
