package errconv

import (
	errh "github.com/hcloud-classic/hcc_errors"
)

// PiccoloHccError : HccErrorType for Piccolo
type PiccoloHccError struct {
	ErrCode uint64 `json:"errcode"` // decimal error code
	ErrText string `json:"errtext"` // error string
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
