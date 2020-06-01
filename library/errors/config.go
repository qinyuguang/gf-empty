package errors

const (
	OK ResponseCode = 0

	INTERNAL_ERROR ResponseCode = 500
	UNAVAILABLE    ResponseCode = 503
)

const (
	INVALID_ARGS ResponseCode = iota + 1001
)

var messageList = map[ResponseCode]string{
	INTERNAL_ERROR: "Internal Error",
	UNAVAILABLE:    "Service Unavailable",

	INVALID_ARGS: "invalid args",
}

type ResponseCode int

func GetMessage(code ResponseCode) string {
	if msg, ok := messageList[code]; ok {
		return msg
	}

	return ""
}
