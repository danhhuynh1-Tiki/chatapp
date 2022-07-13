package responses

type Response struct {
	Status        int                    `json:"status"`         // error code, such as 200, 401, 500
	Message       string                 `json:"message"`        // quickly announce to know success for fail
	DetailMessage string                 `json:"attach_message"` // show details about the message, such as database fail, bad data
	Data          map[string]interface{} `json:"data"`  // the responses data to client
}
