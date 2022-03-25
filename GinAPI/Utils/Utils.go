package Utils

const(
	SUCCESS_CODE = 200
	ERROR_CODE = 300
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}


func ResponseSuccess(message string, data interface{}) Response{
	return Response{
		Code :    SUCCESS_CODE,
		Message : message,
		Data  :   data,
	}
}


func ResponseError(message string, data interface{}) Response {
	return Response{
		Code :    ERROR_CODE,
		Message : message,
		Data  :   data,
	}

}

