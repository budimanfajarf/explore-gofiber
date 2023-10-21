package constant

const (
	ErrInvalidCredentials = "INVALID_CREDENTIALS"
	ErrInvalidPayload     = "INVALID_PAYLOAD"
)

var DefaultErrMessage = map[string]string{
	"INVALID_CREDENTIALS": "invalid email or password",
	"INVALID_PAYLOAD":     "invalid payload",
}
