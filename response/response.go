package response

const (
	DATA_KEY        = "_data_"
	SUCCESS_CODE    = 200
	SUCCESS_MESSAGE = "Success"
)

// R API统一返回结构体
type R struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Error(code int, message string) *R {
	return &R{
		Code:    code,
		Message: message,
	}
}

func Success(data any) *R {
	return &R{
		Code:    SUCCESS_CODE,
		Message: SUCCESS_MESSAGE,
		Data:    data,
	}
}
