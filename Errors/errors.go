package Errors

type ErrorCode int

type CryptoError struct {
    Message string
    Code    ErrorCode
}

const (
    InsufficientIVLength ErrorCode = iota
    InsufficientKeyLength  
)

func NewCryptoError(Code ErrorCode, Message string) *CryptoError {
    return &CryptoError{Code: Code, Message: Message}
}

func (e *CryptoError) Error() string {
    return e.Message
}