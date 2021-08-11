package helper

var (
	// MsgUsersNotFound email / pass not found in database
	MsgUsersNotFound = "Email is invalid. Please check your email is valid before continue"

	// MsgInvalidCredentials email / pass not valid
	MsgInvalidCredentials = "Email / Password is invalid. Please try again"

	// MsgSignatureTokenFailed email / pass not valid
	MsgSignatureTokenFailed = "signature of token is invalid. Please try again"

	// MsgInvalidPermission permission is invalid
	MsgInvalidPermission = "You doesn't have permission to take this action"

	// MsgFailedCreateData create is failed
	MsgFailedCreateData = "Create data failed, please try again or contact administrator."

	// MsgFailedUpdateData update is failed
	MsgFailedUpdateData = "Updating data failed, please try again or contact administrator."

	// MsgDataNotFound data not found
	MsgDataNotFound = "Data not found."

	// MsgDataNotFound data invalid
	MsgInvalidData = "Invalid data"

	// MsgRequired required fill
	MsgRequired = "Some fill required"

	// MsgErrorHashPassword Hash password error
	MsgErrorHashPassword = "Error hashing password"

	// MsgNoAuth No auth provided
	MsgNoAuth = "No Authorization header provided"

	// Msg Invalid token
	MsgInvalidToken = "Token invalid"
)
