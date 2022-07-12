package responses

import (
	"net/http"
)

var ErrorInvalidData = func(err error) Response {
	return Response{
		Status: http.StatusBadRequest,
		Message: "failed",
		DetailMessage: "invalid data",
		Data: map[string]interface{}{
			"error": err.Error(),
		},
	}
}

var ErrorInternalServer = func(err error) Response {
	return Response{
		Status: http.StatusBadRequest,
		Message: "failed",
		DetailMessage: "internal server error",
		Data: map[string]interface{}{
			"error": err.Error(),
		},
	}
}

var ErrorNotFound = func(err error) Response {
	return Response{
		Status: http.StatusNotFound,
		Message: "failed",
		DetailMessage: "not found",
		Data: map[string]interface{}{
			"error": err.Error(),
		},
	}
}

var ErrorUnauthorized = func() Response {
	return Response{
		Status: http.StatusUnauthorized,
		Message: "failed",
		DetailMessage: "unauthorized",
		Data: map[string]interface{}{
			"error": "unauthorized",
		},
	}
}

var Create = func(id string) Response {
	return Response{
		Status: http.StatusOK,
		Message: "success",
		DetailMessage: "created",
		Data: map[string]interface{}{
			"id": id, 
		},
	}
}