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

	modelTelegraf := &model.Telegraf{
		UUID:   monitoringData.Uuid,
		Result: monitoringData.Result,
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
