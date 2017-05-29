package Errors

type AppErrorCode int

type AppError struct {
    Message string
    ErrorCode    AppErrorCode
}

const (
    NoKeyorKeyfile AppErrorCode = iota
    NoPlaintextorPlaintextfile
    NoCiphertextorCiphertextfile
    InvalidKey
    InvalidIV
    InvalidMode
      
)

func NewAppError(Code AppErrorCode, Message string) *AppError {
    return &AppError{ErrorCode: Code, Message: Message}
}

func (e *AppError) Error() string {
    return e.Message
}

