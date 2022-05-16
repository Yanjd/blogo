package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	ErrUsernameUsed  = 1001
	ErrPasswordWrong = 1002
	ErrUserNotExist  = 1003
	ErrTokenExist    = 1004
	ErrTokenRuntime  = 1005
	ErrTokenWrong    = 1006
	ErrTokenFmtWrong = 1007
	ErrUserFmtWrong  = 1008
	// 2xxx - article

	// 3xxx - category

	ErrCateNameUsed = 1001
	ErrCateFmtWrong = 1002
)

var codeMsg = map[int]string{
	SUCCESS:          "OK",
	ERROR:            "Fail",
	ErrUsernameUsed:  "user name existed",
	ErrPasswordWrong: "password error",
	ErrUserNotExist:  "user not exist",
	ErrTokenExist:    "token not exist",
	ErrTokenRuntime:  "token expire",
	ErrTokenWrong:    "token error",
	ErrTokenFmtWrong: "token format error",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
