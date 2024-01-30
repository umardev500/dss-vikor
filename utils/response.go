package utils

type ResponseUtil struct {
	Message string
}

func (r ResponseUtil) Error() string {
	return r.Message
}

func NewResponse() ResponseUtil {
	return ResponseUtil{}
}
