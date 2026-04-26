package response

type ApiResponse struct {
	Status 	bool 	`json:"status"`
	Message string	`json:"message"`
	Data 	any 	`json:"data"`
}

func Success(message string, data any) ApiResponse {
	return ApiResponse{
		Status: true,
		Message: message,
		Data: data,
	}
}

func Error(messsage string) ApiResponse {
	return ApiResponse{
		Status: false,
		Message: messsage,
		Data: nil,
	}
}