package res

/*
	所有的错误码，都会放在这
*/

type ErrorCode int

const (
	SettingError ErrorCode = -1000 //系统错误
)

var (
	ErrMap = map[ErrorCode]string{
		SettingError: "系统出现问题",
	}
)
