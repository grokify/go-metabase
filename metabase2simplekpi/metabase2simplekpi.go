package metabase2simplekpi

import (
	"github.com/grokify/go-metabase/metabaseutil"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	"github.com/grokify/gocharts/data/statictimeseries"
)

func QueryMetabase(cfg Config, dsInfo DatasetInfo) (*SqlResponse, error) {
	resp, err := metabaseutil.QuerySQLHttp(
		cfg.MetabaseHttpClient,
		cfg.MetabaseConfig.BaseUrl,
		int64(dsInfo.MetabaseQueryDatabaseId),
		dsInfo.MetabaseQueryNativeSQL)
	if err != nil {
		return &SqlResponse{}, err
	}
	return HTTPResponseToSqlResponse(resp)
}

func QueryMetabaseSTS(cfg Config, dsInfo DatasetInfo) (statictimeseries.DataSeries, error) {
	sts := statictimeseries.DataSeries{}
	sqlResp, err := QueryMetabase(cfg, dsInfo)
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

func UpdateSimpleKpiSqlResponse(cfg Config, dsInfo DatasetInfo, sqlResp *SqlResponse) []error {
	sts, err := SqlResponseToSTS(dsInfo.KpiName, sqlResp, 0, 1)
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
