package exterror

//common error
var (
	ErrInterval          = NewError(ErrCodeInterval, "interval error")
	ErrParameterInvalid  = NewError(ErrCodeParameterInvalid, "parameter invalid error")
	ErrUnauthorized      = NewError(ErrCodeUnauthorized, "unauthorized error")
	ErrUserHasRegistered = NewError(ErrCodeUserRegistered, "user registered")
	ErrBadRole           = NewError(ErrCodeBadRole, "bad role")
	ErrUserNotExist      = NewError(ErrCodeUserNotExist, "user not exist")
	ErrPasswordWrong     = NewError(ErrCodePasswordWrong, "password is incorrect")
	ErrNodeHasRegistered = NewError(ErrCodeNodeRegistered, "node registered")
	ErrSystemHasInit     = NewError(ErrCodeSystemInit, "system had init already")
)
