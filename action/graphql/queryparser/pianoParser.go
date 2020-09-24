package queryparser

import (
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/action/grpc/pb/rpcpiano"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
)

func pbMonitoringDataToModelTelegraf(monitoringData *rpcpiano.MonitoringData) *model.Telegraf {
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

func checkTelegrafArgsAll(args map[string]interface{}) bool {
	_, metricOk := args["metric"].(string)
	_, subMetricOk := args["subMetric"].(string)
	_, periodOk := args["period"].(string)
	_, aggregateTypeOk := args["aggregateType"].(string)
	_, durationOk := args["duration"].(string)
	_, uuidOk := args["uuid"].(string)

	return metricOk && subMetricOk && periodOk && aggregateTypeOk && durationOk && uuidOk
}

// Telegraf : Set telegraf with provided options
func Telegraf(args map[string]interface{}) (interface{}, error) {
	metric, _ := args["metric"].(string)
	subMetric, _ := args["subMetric"].(string)
	period, _ := args["period"].(string)
	aggregateType, _ := args["aggregateType"].(string)
	duration, _ := args["duration"].(string)
	uuid, _ := args["uuid"].(string)

	if !checkTelegrafArgsAll(args) {
		return model.Telegraf{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGraphQLArgumentError, "check needed arguments (metric, subMetric, period, aggregateType, duration, uuid)")}, nil
	}

	resMonitoringData, err := client.RC.Telegraph(&rpcpiano.ReqMetricInfo{
		MetricInfo: &rpcpiano.MetricInfo{
			Metric:        metric,
			SubMetric:     subMetric,
			Period:        period,
			AggregateType: aggregateType,
			Duration:      duration,
			Uuid:          uuid,
		},
	})
	if err != nil {
		return model.Telegraf{Errors: errors.ReturnHccErrorPiccolo(errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelTelegraf := pbMonitoringDataToModelTelegraf(resMonitoringData.MonitoringData)

	hccErrStack := errconv.GrpcStackToHcc(&resMonitoringData.HccErrorStack)
	modelTelegraf.Errors = *hccErrStack.ConvertReportForm()

	return *modelTelegraf, nil
}
