package component

type ErrorType int

const (
	ErrorEntityNotFound ErrorType = iota
	ErrorEntityExists
	ErrorWhenMapping
	ErrorDb
	ErrorInvalidPayload
	ErrorInvalidAuth
	ErrNoPermission
	ErrInternal
)

func (e ErrorType) String() string {
	switch e {
	case ErrorEntityNotFound:
		return "ERROR_ENTITY_NOT_FOUND"
	case ErrorEntityExists:
		return "ERROR_ENTITY_EXISTS"
	case ErrorWhenMapping:
		return "ERROR_WHEN_MAPPING"
	case ErrorDb:
		return "ERROR_DB"
	case ErrorInvalidPayload:
		return "ERROR_INVALID_PAYLOAD"
	case ErrorInvalidAuth:
		return "ERROR_INVALID_AUTH"
	case ErrNoPermission:
		return "ERROR_NO_PERMISSION"
	default:
		return ""
	}
}
