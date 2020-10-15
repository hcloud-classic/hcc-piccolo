package pbtomodel

import (
	"hcc/piccolo/action/grpc/pb/rpcpiano"
	"hcc/piccolo/model"
)

// PbMonitoringDataToModelTelegraf : Change monitoringData of proto type to telegraf model
func PbMonitoringDataToModelTelegraf(monitoringData *rpcpiano.MonitoringData) *model.Telegraf {
	var seriesArr []model.Series

	for _, monitoringDataSeries := range monitoringData.Series {
		var series model.Series

		series.Time = int(monitoringDataSeries.Time)
		series.Value = int(monitoringDataSeries.Value)

		seriesArr = append(seriesArr, series)
	}

	modelTelegraf := &model.Telegraf{
		Metric:    monitoringData.Metric,
		SubMetric: monitoringData.SubMetric,
		UUID:      monitoringData.UUID,
		Series:    seriesArr,
	}

	return modelTelegraf
}
