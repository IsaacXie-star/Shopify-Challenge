package model

type JsonResponse struct {
	ErrCode int64       `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Page    int64       `json:"page"`
	Size    int64       `json:"size"`
	Total   int64       `json:"total"`
}

func NewResponseWithData(data interface{}) *JsonResponse {
	return &JsonResponse{Data: data}
}

func NewResponseWithMsg(data interface{}, msg string) *JsonResponse {
	return &JsonResponse{Data: data, Msg: msg}
}

func NewResponseWithErrCodeAndMsg(data interface{}, errCode int64, errMsg string) *JsonResponse {
	return &JsonResponse{Data: data, ErrCode: errCode, ErrMsg: errMsg}
}

func NewResponseWithTotal(data interface{}, total int64) *JsonResponse {
	return &JsonResponse{Data: data, Total: total}
}

func NewResponseWithPageSize(data interface{}, page int64, size int64, total int64) *JsonResponse {
	return &JsonResponse{Data: data, Page: page, Size: size, Total: total}
}
