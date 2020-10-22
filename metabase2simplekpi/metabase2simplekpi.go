package metabase2simplekpi

import (
	"github.com/grokify/go-metabase/metabaseutil"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	"github.com/grokify/gocharts/data/statictimeseries"
)

func QueryMetabase(cfg Config, dsInfo DatasetInfo) (*SqlResponse, error) {
	resp, err := metabaseutil.QuerySQLHttp(
		cfg.MetabaseHttpClient,
		cfg.MetabaseConfig.BaseURL,
		dsInfo.MetabaseQuery.DatabaseID,
		dsInfo.MetabaseQuery.NativeSQL())
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
	return SqlResponseToSTS(
		dsInfo.KpiName,
		sqlResp,
		dsInfo.MetabaseQuery.ColIdxCount,
		dsInfo.MetabaseQuery.ColIdxTime)
}

func UpdateSimpleKpi(cfg Config, dsInfo DatasetInfo) []error {
	sts, err := QueryMetabaseSTS(cfg, dsInfo)
	if err != nil {
		return []error{err}
	}
	return UpdateSimpleKpiSTS(cfg, dsInfo, sts)
}

func UpdateSimpleKpiSqlResponse(cfg Config, dsInfo DatasetInfo, sqlResp *SqlResponse) []error {
	sts, err := SqlResponseToSTS(dsInfo.KpiName, sqlResp,
		dsInfo.MetabaseQuery.ColIdxCount,
		dsInfo.MetabaseQuery.ColIdxTime)
	if err != nil {
		return []error{err}
	}
	return UpdateSimpleKpiSTS(cfg, dsInfo, sts)
}

func UpdateSimpleKpiSTS(cfg Config, dsInfo DatasetInfo, sts statictimeseries.DataSeries) []error {
	_, resps, err := simplekpiutil.UpsertKpiEntriesStaticTimeSeries(
		cfg.SimplekpiApiClient,
		int64(cfg.SimplekpiUserID),
		int64(dsInfo.SimplekpiKpiId),
		sts)
	if err != nil {
		return []error{err}
	}
	return simplekpiutil.KpiEntryResponseErrors(resps)
}
