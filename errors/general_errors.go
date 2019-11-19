package errors

var (
	ErrParseBody = UserlandError{
		Code:    REQUEST_BODY_UNDECODABLE,
		Message: REQUEST_BODY_UNDECODABLE_MESSAGE,
	}
)
