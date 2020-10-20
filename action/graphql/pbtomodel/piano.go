package pbtomodel

import (
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcmsgType"
	"hcc/piccolo/action/grpc/pb/rpcpiano"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
)

// PbMonitoringDataToModelTelegraf : Change monitoringData of proto type to telegraf model
func PbMonitoringDataToModelTelegraf(monitoringData *rpcpiano.MonitoringData, hccGrpcErrStack *[]*rpcmsgType.HccError) *model.Telegraf {
	var seriesArr [][]float64

	for _, monitoringDataSeries := range monitoringData.Series {
		var series []float64
		series = (*monitoringDataSeries).Values
		seriesArr = append(seriesArr, series)
	}

	modelTelegraf := &model.Telegraf{
		Metric:    monitoringData.Metric,
		SubMetric: monitoringData.SubMetric,
		UUID:      monitoringData.UUID,
		Series:    seriesArr,
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelTelegraf.Errors = *hccErrStack.ConvertReportForm()
		if len(modelTelegraf.Errors) != 0 && modelTelegraf.Errors[0].ErrCode == 0 {
			modelTelegraf.Errors = errors.ReturnHccEmptyErrorPiccolo()
		}
	} else {
		modelTelegraf.Errors = errors.ReturnHccEmptyErrorPiccolo()
	}

	return modelTelegraf
}
