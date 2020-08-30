package v1

const (
	// Port on which the service will listen
	Port = "9000"

	// MesssageInvalidUserID is the message response when the user-id is invalid in payload
	MesssageInvalidUserID = "Invalid User-ID"

	// MesssageCannotProcessRequest is the message response when the server is not able to process the request
	MesssageCannotProcessRequest = "Cannot Process Request"
)

// HTTPResponse defines response struct for api's to client
type HTTPResponse struct {
	Data  interface{}
	Error error
}

// UserInfo holds the information corresponding to the user-id
type UserInfo struct {
	Message int `json:"message"`
}

// ErrorMessage defines
type ErrorMessage struct {
	Message string `json:"message"`
}
