package errors

var (
	ErrRegistrationIncomplete = UserlandError{
		Code:    REGISTRATION_BODY_INCOMPLETE,
		Message: REGISTRATION_BODY_INCOMPLETE_MESSAGE,
	}

	ErrRegistrationInvalid = UserlandError{
		Code:    REGISTRATION_BODY_INVALID,
		Message: REGISTRATION_BODY_INVALID_MESSAGE,
	}

	ErrRegistrationUnmatchingPassword = UserlandError{
		Code:    REGISTRATION_PASSWORD_NOT_MATCH,
		Message: REGISTRATION_GENERAL_MESSAGE,
	}

	ErrRegistrationQueryExec = UserlandError{
		Code:    REGISTRATION_UNABLE_TO_EXEC_QUERY,
		Message: REGISTRATION_GENERAL_MESSAGE,
	}

	ErrVerificationIncomplete = UserlandError{
		Code:    VERIFICATION_BODY_INCOMPLETE,
		Message: VERIFICATION_BODY_INCOMPLETE_MESSAGE,
	}

	ErrVerificationQueryExec = UserlandError{
		Code:    VERIFICATION_UNABLE_TO_EXEC_QUERY,
		Message: VERIFICATION_GENERAL_MESSAGE,
	}

	ErrLoginIncomplete = UserlandError{
		Code:    LOGIN_INCOMPLETE_CREDENTIALS,
		Message: LOGIN_INCOMPLETE_CREDENTIALS_MESSAGE,
	}

	ErrLoginUnmatch = UserlandError{
		Code:    LOGIN_PASSWORD_NOT_MATCH,
		Message: LOGIN_PASSWORD_NOT_MATCH_MESSAGE,
	}

	ErrLoginUnverified = UserlandError{
		Code:    LOGIN_ACCOUNT_UNVERIFIED,
		Message: LOGIN_ACCOUNT_UNVERIFIED_MESSAGE,
	}

	ErrLoginJWT = UserlandError{
		Code:    LOGIN_JWT_ERROR,
		Message: LOGIN_JWT_ERROR_MESSAGE,
	}

	ErrForgetPassIncomplete = UserlandError{
		Code:    FORGET_PASSWORD_INCOMPLETE_CREDENTIALS,
		Message: FORGET_PASSWORD_INCOMPLETE_CREDENTIALS_MESSAGE,
	}

	ErrForgetPassQueryExec = UserlandError{
		Code:    FORGET_PASSWORD_UNABLE_TO_EXEC_QUERY,
		Message: FORGET_PASSWORD_GENERAL_MESSAGE,
	}

	ErrResetPassInvalid = UserlandError{
		Code:    RESET_PASSWORD_BODY_INVALID,
		Message: RESET_PASSWORD_BODY_INVALID_MESSAGE,
	}

	ErrResetPassUnmatchPass = UserlandError{
		Code:    RESET_PASSWORD_PASSWORD_NOT_MATCH,
		Message: RESET_PASSWORD_PASSWORD_NOT_MATCH_MESSAGE,
	}

	ErrResetPassInvalidPass = UserlandError{
		Code:    RESET_PASSWORD_PASSWORD_INVALID,
		Message: RESET_PASSWORD_PASSWORD_INVALID_MESSAGE,
	}

	ErrResetPassQueryExec = UserlandError{
		Code:    RESET_PASSWORD_UNABLE_TO_EXEC_QUERY,
		Message: RESET_PASSWORD_GENERAL_MESSAGE,
	}

	ErrTokenNotProvided = UserlandError{
		Code:    TOKEN_NOT_PROVIDED,
		Message: TOKEN_NOT_PROVIDED_MESSAGE,
	}

	ErrTokenNotFound = UserlandError{
		Code:    TOKEN_CANNOT_BE_FOUND,
		Message: TOKEN_GENERAL_MESSAGE,
	}

	ErrTokenInvalidSignature = UserlandError{
		Code:    TOKEN_INVALID_SIGNATURE,
		Message: TOKEN_GENERAL_MESSAGE,
	}

	ErrTokenInvalidContent = UserlandError{
		Code:    TOKEN_INVALID_CONTENT,
		Message: TOKEN_GENERAL_MESSAGE,
	}

	ErrTokenExpired = UserlandError{
		Code:    TOKEN_EXPIRED,
		Message: TOKEN_GENERAL_MESSAGE,
	}

	ErrTokenUserIdDoesNotExist = UserlandError{
		Code:    TOKEN_USER_ID_DOES_NOT_EXIST,
		Message: TOKEN_GENERAL_MESSAGE,
	}
)
