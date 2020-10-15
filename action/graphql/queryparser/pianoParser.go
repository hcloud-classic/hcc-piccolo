package queryparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/pb/rpcpiano"
	"hcc/piccolo/lib/errors"
	"hcc/piccolo/model"
)

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

	modelTelegraf := pbtomodel.PbMonitoringDataToModelTelegraf(resMonitoringData.MonitoringData, &resMonitoringData.HccErrorStack)

	return *modelTelegraf, nil
}
