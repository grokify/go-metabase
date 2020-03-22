package metabase2simplekpi

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	mo "github.com/grokify/oauth2more/metabase"
)

type Config struct {
	MetabaseConfig     mo.Config
	MetabaseHttpClient *http.Client
	SimplekpiConfig    simplekpiutil.Config
	SimplekpiUserID    int64
	SimplekpiApiClient *simplekpi.APIClient
	LoadSimpleKpi      bool
	Datasets           []DatasetInfo
}

// DatasetInfo captures information for a single SQL query representing
// data for a single KPI.
type DatasetInfo struct {
	KpiName                      string        `json:"kpiName"`
	MetabaseQueryDatabaseId      int           `json:"metabaseQueryDatabaseId"`
	MetabaseQueryNativeSQL       string        `json:"metabaseQuerySQLNative"`
	MetabaseQueryNativeSQLFormat string        `json:"metabaseQuerySQLNativeFormat"`
	MetabaseQueryNativeSQLVars   []interface{} `json:"metabaseQuerySQLNativeVars"`
	MetabaseQueryColIdxCount     int           `json:"metabaseQueryColIdxCount"`
	MetabaseQueryColIdxDate      int           `json:"metabaseQueryColIdxDate"`
	SimplekpiKpiId               int           `json:"simplekpiKpiId"`
}

func (cfg *Config) InitClients() error {
	err := cfg.InitMetabaseClient()
	if err != nil {
		return err
	}
	return cfg.InitSimplekpiClient()
}

func (cfg *Config) InitMetabaseClient() error {
	if cfg.MetabaseHttpClient != nil {
		return nil
	}
	httpClient, _, err := mo.NewClient(cfg.MetabaseConfig)
	if err != nil {
		return err
	}
	cfg.MetabaseHttpClient = httpClient
	return nil
}

func (cfg *Config) InitSimplekpiClient() error {
	if cfg.SimplekpiApiClient != nil {
		return nil
	}
	client, err := simplekpiutil.NewApiClient(
		cfg.SimplekpiConfig.Site,
		cfg.SimplekpiConfig.Username,
		cfg.SimplekpiConfig.Token)
	if err != nil {
		return err
	}
	cfg.SimplekpiApiClient = client
	return nil
}

// NativeSQL returns a formatted or raw SQL statement.
func (dsi *DatasetInfo) NativeSQL() string {
	dsi.MetabaseQueryNativeSQL = strings.TrimSpace(dsi.MetabaseQueryNativeSQL)
	dsi.MetabaseQueryNativeSQLFormat = strings.TrimSpace(dsi.MetabaseQueryNativeSQLFormat)
	if len(dsi.MetabaseQueryNativeSQLFormat) > 0 {
		if len(dsi.MetabaseQueryNativeSQLVars) > 0 {
			return fmt.Sprintf(
				dsi.MetabaseQueryNativeSQLFormat,
				dsi.MetabaseQueryNativeSQLVars...)
		}
		return dsi.MetabaseQueryNativeSQLFormat
	}
	return dsi.MetabaseQueryNativeSQL
}
