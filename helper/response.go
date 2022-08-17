package helper

type res struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

func ResSuccess(message string, data interface{}) res {
	return res{
		Status:  true,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
}

func ResFailed(message string, err interface{}) res {
	return res{
		Status:  false,
		Message: message,
		Data:    nil,
		Errors:  err,
	}
}
