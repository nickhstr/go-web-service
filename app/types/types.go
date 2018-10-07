package types

// ErrorResponse provides info about the error(s) returned
type ErrorResponse struct {
	MetaData metadata `json:"metadata"`
	Errors   []err    `json:"errors"`
}

type err struct {
	Message string `json:"message"`
}

type metadata struct {
	TransactionID string `json:"transactionId"`
}

type transactionKey string

// TransactionKey provides the key used to store a request's
// transaction ID, allowing access to the ID stored in the
// request's context.
var TransactionKey = transactionKey("transaction")

// NewError returns a new ErrorResponse
func NewError(id, message string) *ErrorResponse {
	return &ErrorResponse{
		MetaData: metadata{
			TransactionID: id,
		},
		Errors: []err{
			err{message},
		},
	}
}
