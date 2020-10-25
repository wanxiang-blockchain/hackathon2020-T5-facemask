package context

type Response struct {
	Success bool        `json:"success"`
	Msg     interface{} `json:"msg"`
}

func NewResponse(success bool, msg interface{}) *Response {
	return &Response{Success: success, Msg: msg}
}

func NewOkResponse(msg interface{}) *Response {
	return NewResponse(true, msg)
}

func NewFailedResponse(i interface{}) *Response {
	return NewResponse(false, i)
}
