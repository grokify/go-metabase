package metabase2simplekpi

import (
	"fmt"
	"strings"
)

// DatasetInfo captures information for a single SQL query representing
// data for a single KPI.
type DatasetInfo struct {
	KpiName                      string        `json:"kpiName"`
	ExecMBQuery                  bool          `json:"execMBQuery"`
	ExecSKUpdate                 bool          `json:"execSKUpdate"`
	MetabaseQueryDatabaseId      int           `json:"metabaseQueryDatabaseId"`
	MetabaseQueryNativeSQL       string        `json:"metabaseQueryNativeSQL"`
	MetabaseQueryNativeSQLFormat string        `json:"metabaseQueryNativeSQLFormat"`
	MetabaseQueryNativeSQLVars   []interface{} `json:"metabaseQueryNativeSQLVars"`
	MetabaseQueryColIdxCount     int           `json:"metabaseQueryColIdxCount"`
	MetabaseQueryColIdxDate      int           `json:"metabaseQueryColIdxDate"`
	SimplekpiKpiId               int           `json:"simplekpiKpiId"`
}

func (dsi *DatasetInfo) Inflate() {
	dsi.MetabaseQueryNativeSQL = dsi.NativeSQL()
}

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
