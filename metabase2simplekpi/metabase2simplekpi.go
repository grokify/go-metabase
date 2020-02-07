package metabase2simplekpi

// $ go run sql.go -d 3 -q abc -s 'https://data.corp.ringcentral.com/pla-prod' -t 'ab33281e-d9c1-46e5-8c3b-63ea9e4a77fb'

import (
	"github.com/grokify/go-metabase/metabaseutil"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	"github.com/grokify/gocharts/data/statictimeseries"
)

func QueryMetabaseSTS(cfg Config, dsInfo DatasetInfo) (statictimeseries.DataSeries, error) {
	sts := statictimeseries.DataSeries{}
	resp, err := metabaseutil.QuerySQLHttp(
		cfg.MetabaseHttpClient,
		cfg.MetabaseConfig.BaseUrl,
		int64(dsInfo.MetabaseQueryDatabaseId),
		dsInfo.MetabaseQueryNativeSQL)
	if err != nil {
		return sts, err
	}
	sqlResp, err := HTTPResponseToSqlResponse(resp)
	if err != nil {
		return sts, err
	}
	sts, err = SqlResponseToSTS(dsInfo.KpiName, sqlResp, 0, 1)
	return sts, err
}

func UpdateSimpleKpiSTS(cfg Config, dsInfo DatasetInfo) []error {
	sts, err := QueryMetabaseSTS(cfg, dsInfo)
	if err != nil {
		return []error{err}
	}
	_, resps, err := simplekpiutil.UpsertKpiEntriesStaticTimeSeries(
		cfg.SimplekpiApiClient,
		int64(cfg.SimplekpiUserID),
		int64(dsInfo.SimplekpiKpiId),
		sts)
	errs := simplekpiutil.KpiEntryResponseErrors(resps)
	return errs
}
