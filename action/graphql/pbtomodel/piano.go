package pbtomodel

import (
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/model"

	"innogrid.com/hcloud-classic/pb"
)

// PbMonitoringDataToModelTelegraf : Change monitoringData of proto type to telegraf model
func PbMonitoringDataToModelTelegraf(monitoringData *pb.MonitoringData, hccGrpcErrStack *pb.HccErrorStack) *model.Telegraf {

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

// PbBillingDataToModelBillingData : Change billingData of proto type to billingData model
func PbBillingDataToModelBillingData(billingData *pb.ResBillingData) *model.BillingData {

	modelBillingData := &model.BillingData{
		BillingType: billingData.BillingType,
		GroupID:     billingData.GroupID,
		Result:      string(billingData.Result),
	}

	if billingData.HccErrorStack != nil {
		hccErrStack := errconv.GrpcStackToHcc(billingData.HccErrorStack)
		modelBillingData.Errors = errconv.HccErrorToPiccoloHccErr(*hccErrStack)
		if len(modelBillingData.Errors) != 0 && modelBillingData.Errors[0].ErrCode == 0 {
			modelBillingData.Errors = errconv.ReturnHccEmptyErrorPiccolo()
		}
	} else {
		modelBillingData.Errors = errconv.ReturnHccEmptyErrorPiccolo()
	}

	return modelBillingData
}
