package pbtomodel

import (
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"

	"github.com/hcloud-classic/pb"
)

// PbMonitoringDataToModelTelegraf : Change monitoringData of proto type to telegraf model
func PbMonitoringDataToModelTelegraf(monitoringData *pb.MonitoringData, hccGrpcErrStack *[]*pb.HccError) *model.Telegraf {

	modelTelegraf := &model.Telegraf{
		UUID:   monitoringData.Uuid,
		Result: string(monitoringData.Result),
	}

	if hccGrpcErrStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(hccGrpcErrStack)
		modelTelegraf.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
		if len(modelTelegraf.Errors) != 0 && modelTelegraf.Errors[0].ErrCode == 0 {
			modelTelegraf.Errors = errconv.ReturnHccEmptyErrorPiccolo()
		}
	} else {
		modelTelegraf.Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return modelTelegraf
}
