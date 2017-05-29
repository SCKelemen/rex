package AES

import (
  	"crypto/aes"
  	"crypto/cipher"
  	"crypto/rand"
    "github.com/SCKelemen/rex/Errors"
    "bytes"
  	"io"
  )

// CTR is an instance of the 
// AES Counter mode
type CTR struct {

}

// Encrypt uses a []byte 'key' to encrypt []byte 'plaintext'
// using the Counter mode of AES
func (c *CTR) Encrypt(plain []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
  	if err != nil {
  		panic(err)
  	}

  	ciphertext := make([]byte, aes.BlockSize+len(plain))
  	iv := ciphertext[:aes.BlockSize]
  	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
  		panic(err)
  	}
  
  	stream := cipher.NewCTR(block, iv)
  	stream.XORKeyStream(ciphertext[aes.BlockSize:], plain)
    return ciphertext, nil

}

// Decrypt uses []byte 'key' to Decrypt
// []byte 'ciphertext' using the Counter
// mode of AES
func (c *CTR) Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    DealWith(err)
    iv := ciphertext[:aes.BlockSize]
    plaintext := make([]byte, len(ciphertext[aes.BlockSize:]))
  	stream := cipher.NewCTR(block, iv)
  	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])
    return plaintext, nil
}

// DecryptWithIV decrypts []byte 'ciphertext' with []byte 'key'
// using the Counter Mode of AES, and Initialization Vector []byte 'iv'
// Only use this method if the IV is not prepended to the ciphertext
func (c *CTR) DecryptWithIV(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    DealWith(err)
    if len(iv) < aes.BlockSize {
        var message bytes.Buffer
        message.WriteString("Invalid IV: \r\n AES requires an Initialization Vector of ")
	    message.WriteString(string(aes.BlockSize)) 
        message.WriteString("bytes, but received an IV of only ")
        message.WriteString(string(len(iv)))
        message.WriteString("bytes. ")
        return nil, Errors.NewCryptoError(Errors.InsufficientIVLength, message.String())
    }
    plaintext := make([]byte, len(ciphertext[aes.BlockSize:]))
  	stream := cipher.NewCTR(block, iv[:aes.BlockSize])
  	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])
    return plaintext, nil
}


func DealWith(e error) {
	if e != nil {
        panic(e)
    }
}



type Encryptor struct {
	w          io.Writer
  	err        error
}

func NewEncryptor(w io.Writer) *Encryptor {
  	return &Encryptor{w: w}
}

func (e *Encryptor) Encrypt(v interface{}) error {
  	if e.err != nil {
  		return e.err
  	}
	_, err := e.w.Write([]byte("dddddddddd"))
	return err
}

type Decryptor struct {
	r     	io.Reader
  	buf   	[]byte
	err		error
}

func NewDecryptor(r io.Reader) *Decryptor {
	return &Decryptor{r: r}
}
func (d *Decryptor) Decrypt(v interface{}) error {
	if d.err != nil {
		return d.err
	}

	return d.err
}






