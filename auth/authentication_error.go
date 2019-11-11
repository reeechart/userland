package auth

const (
	REQUEST_BODY_UNDECODABLE               = 100
	REGISTRATION_BODY_INCOMPLETE           = 101
	REGISTRATION_PASSWORD_NOT_MATCH        = 102
	REGISTRATION_UNABLE_TO_EXEC_QUERY      = 103
	VERIFICATION_BODY_INCOMPLETE           = 104
	VERIFICATION_UNABLE_TO_EXEC_QUERY      = 105
	LOGIN_INCOMPLETE_CREDENTIALS           = 106
	LOGIN_PASSWORD_DOES_NOT_MATCH          = 107
	LOGIN_JWT_ERROR                        = 108
	FORGET_PASSWORD_INCOMPLETE_CREDENTIALS = 109
	FORGET_PASSWORD_UNABLE_TO_EXEC_QUERY   = 110
	RESET_PASSWORD_PASSWORD_NOT_MATCH      = 111
	RESET_PASSWORD_UNABLE_TO_EXEC_QUERY    = 112
	TOKEN_NOT_PROVIDED                     = 113
	TOKEN_CANNOT_BE_FOUND                  = 114
	TOKEN_INVALID_SIGNATURE                = 115
	TOKEN_INVALID_CONTENT                  = 116
	TOKEN_EXPIRED                          = 117
	TOKEN_EMAIL_DOES_NOT_EXIST             = 118
)
