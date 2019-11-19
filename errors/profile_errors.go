package errors

var (
	ErrUserInfoInvalid = UserlandError{
		Code:    UPDATE_PROFILE_USER_INFO_INVALID,
		Message: UPDATE_PROFILE_USER_INFO_INVALID_MESSAGE,
	}

	ErrUpdateProfileQueryExec = UserlandError{
		Code:    UPDATE_PROFILE_UNABLE_TO_EXEC_QUERY,
		Message: UPDATE_PROFILE_UNABLE_TO_EXEC_QUERY_MESSAGE,
	}

	ErrChangeEmailQueryExec = UserlandError{
		Code:    CHANGE_EMAIL_UNABLE_TO_EXEC_QUERY,
		Message: CHANGE_EMAIL_UNABLE_TO_EXEC_QUERY_MESSAGE,
	}

	ErrChangeEmailInvalidEmail = UserlandError{
		Code:    CHANGE_EMAIL_EMAIL_INVALID,
		Message: CHANGE_EMAIL_EMAIL_INVALID_MESSAGE,
	}

	ErrChangePasswordPasswordUnmatch = UserlandError{
		Code:    CHANGE_PASSWORD_PASSWORD_NOT_MATCH,
		Message: CHANGE_PASSWORD_PASSWORD_NOT_MATCH_MESSAGE,
	}

	ErrChangePasswordInvalidPassword = UserlandError{
		Code:    CHANGE_PASSWORD_PASSWORD_INVALID,
		Message: CHANGE_PASSWORD_PASSWORD_INVALID_MESSAGE,
	}

	ErrChangePasswordIncorrectCurrentPass = UserlandError{
		Code:    CHANGE_PASSWORD_INCORRECT_CURRENT_PASSWORD,
		Message: INCORRECT_PASSWORD_GENERAL_MESSAGE,
	}

	ErrDeleteAccountIncorrectPass = UserlandError{
		Code:    DELETE_ACCOUNT_INCORRECT_PASSWORD,
		Message: INCORRECT_PASSWORD_GENERAL_MESSAGE,
	}

	ErrUpdatePictureQueryExec = UserlandError{
		Code:    PICTURE_UNABLE_TO_EXEC_QUERY,
		Message: PICTURE_GENERAL_MESSAGE,
	}

	ErrUpdatePicturePicCantBeFetched = UserlandError{
		Code:    PICTURE_CANNOT_BE_FETCHED_FROM_FORM,
		Message: PICTURE_FORMAT_GENERAL_MESSAGE,
	}

	ErrUpdatePictureCantBeRead = UserlandError{
		Code:    PICTURE_CANNOT_BE_READ,
		Message: PICTURE_FORMAT_GENERAL_MESSAGE,
	}
)
