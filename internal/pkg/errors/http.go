package errors


type HttpError struct {
	status int
	message string
	err error
}

func (err *HttpError) Error() string {
	return ``
}