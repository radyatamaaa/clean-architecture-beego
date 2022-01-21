package models

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrUnAuthorize = errors.New("Unauthorize")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist or duplicate")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Bad Request")

	ErrPublicKey = errors.New("invalid Public Key")

	ErrInvalidLatitude = errors.New("invalid Latitude")

	ErrInvalidLongitude = errors.New("invalid Longitude")

	ErrInvalidDataType = errors.New("invalid data type")

	ErrIsRequired = errors.New("is required")

	ErrInvalidValue = errors.New("invalid value")

	ErrForbidden = errors.New("you don't have permission to access this resource")
)

//authValidation
var (
	ErrInvalidToken              = errors.New("Invalid authorization token")
	ErrInvalidTokenType          = errors.New("Authorization token type does not match")
	ErrNotMatchTokenCredentials  = errors.New("Authorization token credentials do not match")
	ErrInvalidTokenCredentials   = errors.New("Invalid authorization token credentials")
	ErrInvalidTokenExpired       = errors.New("Authorization token has expired")
	ErrUsername                  = errors.New("Please Check again your Username")
	ErrPassword                  = errors.New("Please Check again your Password")
	LoginFailedMessage           = errors.New("Kombinasi personal number dan password Anda tidak tepat. Mohon periksa kembali. Demi alasan keamanan, akun Anda akan terkunci sementara jika 3x gagal percobaan login.")
	LoginFailedMessageLockedUser = errors.New("Demi alasan keamanan ,Akun Anda terkunci sementara selama 15 menit karena telah gagal 3x percobaan silahkan coba lagi beberapa saat.")

	PersonalNumberNotFound = errors.New("Personal number Anda tidak terdaftar di BRIBrain Webview karena untuk saat ini BRIBrain hanya dapat dilihat oleh personal number tertentu saja.")
)

var (
	GeneralSuccess = "Success"

	ErrGeneralMessage      = errors.New("something wrong")
	ErrMessageCaptcha      = errors.New("Captcha yang Anda masukkan salah. Silakan coba kembali.")
	ErrRekomendasiNotFound = errors.New("Rekomendasi Tidak Di temukan")
)
