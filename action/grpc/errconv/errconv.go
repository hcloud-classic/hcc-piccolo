package errconv

import (
	errh "github.com/hcloud-classic/hcc_errors"
	errg "github.com/hcloud-classic/pb"
)

// PiccoloHccError : HccErrorType for Piccolo
type PiccoloHccError struct {
	ErrCode uint64 `json:"errcode"` // decimal error code
	ErrText string `json:"errtext"` // error string
}

// GrpcToHcc : Convert gRPC error code and text to HCCError
func GrpcToHcc(eg *errg.HccError) *errh.HccError {
	return errh.NewHccError(eg.GetErrCode(), eg.GetErrText())
}

// HccToGrpc : Convert HCCError to gRPC error
func HccToGrpc(eh *errh.HccError) *errg.HccError {
	return &errg.HccError{ErrCode: eh.Code(), ErrText: eh.Text()}
}

// GrpcStackToHcc : Convert gRPC error stack to HCCError stack
func GrpcStackToHcc(esg *[]*errg.HccError) *errh.HccErrorStack {
	errStack := errh.NewHccErrorStack()

	for i, e := range *esg {
		if i == 0 {
			continue
		}

		_ = errStack.Push(errh.NewHccError(e.GetErrCode(), e.GetErrText()))
	}

	return errStack
}

// HccStackToGrpc : Convert HCCError stack to gRPC error stack
func HccStackToGrpc(esh *errh.HccErrorStack) []*errg.HccError {
	ges := []*errg.HccError{}
	for i := 0; i <= esh.Len(); i++ {
		ge := &errg.HccError{ErrCode: (*esh.Stack())[i].Code(), ErrText: (*esh.Stack())[i].Text()}
		ges = append(ges, ge)
	}
	return ges
}

// HccErrorStackToPiccoloHccErr: Convert HccErrorStack type to PiccoloHccErr type
func HccErrorToPiccoloHccErr(hccErr errh.HccErrorStack) []PiccoloHccError {
	var piccoloHccErr []PiccoloHccError

	form := hccErr.ConvertReportForm()
	if form != nil {
		stack := form.Stack()
		for _, s := range *stack {
			piccoloHccErr = append(piccoloHccErr, PiccoloHccError{
				ErrCode: s.Code(),
				ErrText: s.Text(),
			})
		}
	}

	return piccoloHccErr
}

// ReturnHccErrorPiccolo: Get error code and error string and return as HccErrorStack
func ReturnHccErrorPiccolo(errCode uint64, errText string) []PiccoloHccError {
	var piccoloHccErr []PiccoloHccError

	stack := *errh.NewHccErrorStack(errh.NewHccError(errCode, errText)).ConvertReportForm().Stack()
	for _, s := range stack {
		piccoloHccErr = append(piccoloHccErr, PiccoloHccError{
			ErrCode: s.Code(),
			ErrText: s.Text(),
		})
	}

	return piccoloHccErr
}

// ReturnHccEmptyErrorPiccolo: Return dummy error as HccErrorStack
func ReturnHccEmptyErrorPiccolo() []PiccoloHccError {
	var piccoloHccErr []PiccoloHccError

	return piccoloHccErr
}
