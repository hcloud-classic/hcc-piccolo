package queryparser

import (
	"errors"
	"hcc/piccolo/data"
	"hcc/piccolo/http"
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
		return nil, errors.New("check needed arguments (metric, subMetric, period, aggregateType, duration, uuid)")
	}
	var telegrafData data.TelegrafData
	query := "query { telegraf(metric:\"" + metric + "\", subMetric:\"" + subMetric + "\", period:\"" + period + "\", " +
		"aggregateType:\"" + aggregateType + "\", duration:\"" + duration + "\", uuid:\"" + uuid + "\"){ metric subMetric id data { " +
		"x y } } }"

	return http.DoHTTPRequest("piano", true, "TelegrafData", telegrafData, query)
}
