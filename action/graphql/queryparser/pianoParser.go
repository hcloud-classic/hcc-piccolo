package queryparser

import (
	"hcc/piccolo/action/graphql/pbtomodel"
	"hcc/piccolo/action/grpc/client"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"
	"strconv"
	"strings"

	"innogrid.com/hcloud-classic/hcc_errors"
	"innogrid.com/hcloud-classic/pb"
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
	time, timeOk := args["time"].(string)
	groupBy, _ := args["groupBy"].(string)
	orderBy, _ := args["orderBy"].(string)
	limit, _ := args["limit"].(string)

	if timeOk {
		time = time + "000000"
	}

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

	modelTelegraf := pbtomodel.PbMonitoringDataToModelTelegraf(resMonitoringData.MonitoringData, resMonitoringData.HccErrorStack)

	return *modelTelegraf, nil
}

// GetBillingData : Get billing data with provided options
func GetBillingData(args map[string]interface{}, isAdmin bool, isMaster bool) (interface{}, error) {
	if !isMaster || !isAdmin {
		return model.BillingData{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLInvalidToken, "Permission denied!")}, nil
	}

	groupID, groupIDOk := args["group_id"].(string)
	billingType, _ := args["billing_type"].(string)
	dateStart, _ := args["date_start"].(int)
	dateEnd, _ := args["date_end"].(int)
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)

	if !groupIDOk {
		return model.BillingData{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError,
			"need a group_id argument")}, nil
	}

	var reqBillingData = pb.ReqBillingData{
		BillingType: billingType,
		DateStart:   int32(dateStart),
		DateEnd:     int32(dateEnd),
	}
	if rowOk {
		reqBillingData.Row = int64(row)
	}
	if pageOk {
		reqBillingData.Page = int64(page)
	}

	var groupIDs []int32
	groupIDsStr := strings.Split(groupID, ".")
	for _, groupIDStr := range groupIDsStr {
		gid, err := strconv.Atoi(groupIDStr)
		if err == nil {
			groupIDs = append(groupIDs, int32(gid))
		}
	}
	reqBillingData.GroupID = groupIDs

	resBillingData, err := client.RC.GetBillingData(&reqBillingData)
	if err != nil {
		return model.BillingData{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGrpcRequestError, err.Error())}, nil
	}

	modelBillingData := pbtomodel.PbBillingDataToModelBillingData(resBillingData)

	return modelBillingData, nil
}
