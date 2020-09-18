package constants

const (
	TokenSecretKey        = "blueBird_sEcrET_key$"
	TokenExpiredInMinutes = 8 * 60 * 60
)

const (
	ERR_CODE_00     = "00"
	ERR_CODE_00_MSG = "Success"

	ERR_CODE_03     = "03"
	ERR_CODE_03_MSG = "Error, unmarshall body Request"

	ERR_CODE_54     = "54"
	ERR_CODE_54_MSG = "Invalid Authorization"

	ERR_CODE_55     = "55"
	ERR_CODE_55_MSG = "Token expired"

	ERR_CODE_50     = "50"
	ERR_CODE_50_MSG = "Invalid username / password"

	ERR_CODE_51     = "51"
	ERR_CODE_51_MSG = "Error connection to database"

	ERR_CODE_53     = "53"
	ERR_CODE_53_MSG = "Failed generate token"
)