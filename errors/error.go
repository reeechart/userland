package errors

type UserlandError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err UserlandError) Error() string {
	return err.Message
}
