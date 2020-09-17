package errors

// ReturnHccErrorPiccolo: Get error code and error string and return as HccErrorStack
func ReturnHccErrorPiccolo(errCode uint64, errText string) HccErrorStack {
	return *NewHccErrorStack(NewHccError(errCode, errText)).ConvertReportForm()
}

// ReturnHccEmptyError: Return dummy error as HccErrorStack
func ReturnHccEmptyError() HccErrorStack {
	return *NewHccErrorStack().ConvertReportForm()
}
