package AES

import (
  	"crypto/aes"
  	"crypto/cipher"
  	"crypto/rand"

  	"fmt"
  	"io"
  )

// CTR is an instance of the 
// AES Counter mode
type CTR struct {

}

// Encrypt uses a []byte 'key' to encrypt []byte 'plaintext'
// using the Counter mode of AES
func (c *CTR) Encrypt(plaintext []byte, key []byte) []byte {

  	block, err := aes.NewCipher(key)
  	if err != nil {
  		panic(err)
  	}
  
  	// The IV needs to be unique, but not secure. Therefore it's common to
  	// include it at the beginning of the ciphertext.
  	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
  	iv := ciphertext[:aes.BlockSize]
  	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
  		panic(err)
  	}
  
  	stream := cipher.NewCTR(block, iv)
  	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
  
  
  
  	fmt.Printf("%s\n", plaintext)
    return plaintext
  
}

// Decrypt uses []byte 'key' to Decrypt
// []byte 'ciphertext' using the Counter
// mode of AES
func (c *CTR) Decrypt(ciphertext []byte, key []byte) []byte {
    return c.Encrypt(ciphertext, key)
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