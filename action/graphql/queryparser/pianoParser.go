package queryparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"

	"github.com/hcloud-classic/hcc_errors"
	"github.com/hcloud-classic/pb"
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
	uuid, _ := args["uuid"].(string)
	metric, _ := args["metric"].(string)
	subMetric, _ := args["subMetric"].(string)
	period, _ := args["period"].(string)
	aggregateFn, _ := args["aggregateFn"].(string)
	duration, _ := args["duration"].(string)
	time, _ := args["time"].(string)
	groupBy, _ := args["groupBy"].(string)
	orderBy, _ := args["orderBy"].(string)
	limit, _ := args["limit"].(string)

	resMonitoringData, err := client.RC.Telegraph(&pb.ReqMetricInfo{
		MetricInfo: &pb.MetricInfo{
			Uuid:        uuid,
			Metric:      metric,
			SubMetric:   subMetric,
			Period:      period,
			AggregateFn: aggregateFn,
			Duration:    duration,
			Time:        time,
			GroupBy:     groupBy,
			OrderBy:     orderBy,
			Limit:       limit,
		},
	})
	if err != nil {
		return model.Telegraf{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelTelegraf := pbtomodel.PbMonitoringDataToModelTelegraf(resMonitoringData.MonitoringData, &resMonitoringData.HccErrorStack)

	return *modelTelegraf, nil
}
