package component

type ErrorMessage int

const (
	ErrMessageFromDB ErrorMessage = iota
	ErrMessageNotFoundFromDB
	ErrMessageInvalidPayload
	ErrMessageInternal
)

func (e ErrorMessage) String() string {
	switch e {
	case ErrMessageFromDB:
		return "error when query in db"
	case ErrMessageNotFoundFromDB:
		return "error not found any record when query in db"
	case ErrMessageInvalidPayload:
		return "invalid payload"
	case ErrMessageInternal:
		return "internal error"
	default:
		return ""
	}
}
