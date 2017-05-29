package main

import (
	"fmt"
	"os"
	"bytes"
	"io"
	"strings"
	"github.com/urfave/cli"

	 "github.com/SCKelemen/rex/Errors"
   router "github.com/SCKelemen/rex/router"
)

func main() {
	app := cli.NewApp()
  app.Name = "rex"
  r := router.NewRouter(nil)
  
  var config Settings


  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name:        "mode, m",
      Value:       "ctr",
      Usage:       "Cipher Mode",
      Destination: &config.mode,
    },
    cli.StringFlag{
      Name:        "key, k",
      Value:       "",
      Usage:       "Encryption/Decryption key",
      Destination: &config.key, 
    },
    cli.StringFlag{
      Name:        "keyfile, kf",
      Value:       "",
      Usage:       "Path to KeyFile containing Encryption/Decryption key",
      Destination: &config.keyfile,
    },
    cli.StringFlag{
      Name:        "plaintext, plain, p",
      Value:       "",
      Usage:       "Plaintext to encrypt",
      Destination: &config.plaintext,
    },
    cli.StringFlag{
      Name:        "plaintextfile, plainfile, pf",
      Value:       "",
      Usage:       "Path to file containing plaintext to encrypt",
      Destination: &config.plaintextfile,
    },
    cli.StringFlag{
      Name:        "ciphertext, cipher, c",
      Value:       "",
      Usage:       "Ciphertext to decrypt",
      Destination: &config.ciphertext,
    },
    cli.StringFlag{
      Name:        "ciphertextfile, cipherfile, cf",
      Value:       "",
      Usage:       "Path to file containing ciphertext to decrypt",
      Destination: &config.mode,
    },
  }

 	app.Commands = []cli.Command{
	{
      Name:        "aes",
      Usage:       "aes",
      Subcommands: []cli.Command{
        {
          Name:  "encrypt",
          Category: "AES",
          Usage: "Encrypts a files with AES",
          Action: func(c *cli.Context) error {
            params, err := getEncryptionParameters(config)
            DealWith(err)
            var ciphertext []byte
            if params.iv != "" {
              ciphertext, err = r.NewAESRouter().CTR.EncryptWithIV([]byte(params.plain), []byte(params.key), []byte(params.iv))
              DealWith(err)
            }
            ciphertext, err = r.NewAESRouter().CTR.Encrypt([]byte(params.plain), []byte(params.key))
            DealWith(err)
            fmt.Printf("%s\n", ciphertext)
               
            return nil
          },
        },
        {
          Name:  "decrypt",
          Category: "AES",
          Usage: "Decrypts",
          Action: func(c *cli.Context) error {

            params, err := getDecryptionParameters(config)
            DealWith(err)
            var plaintext []byte
            if params.iv != "" {
              plaintext, err = r.NewAESRouter().CTR.DecryptWithIV([]byte(params.plain), []byte(params.key), []byte(params.iv))
              DealWith(err)
            }
            plaintext, err = r.NewAESRouter().CTR.Decrypt([]byte(params.plain), []byte(params.key))
            DealWith(err)
            fmt.Printf("%s\n", plaintext)

            return nil
          },
        },	 
      },
    },

  }

  app.Run(os.Args)
}


func DealWith(e error) {
	if e != nil {
        panic(e)
    }
}

type Header struct {
	path string
	name string
	extention string
	iv []byte
	hash []byte
	hashalg string
	version int
}



func EncryptToFile(filepath string) error {
	

	source, err := os.Open(filepath)
	DealWith(err)
	defer source.Close()
	
	var filename bytes.Buffer
    filename.WriteString(source.Name())
	filename.WriteString(".aes")

	destination, err := os.Create(filename.String())
	DealWith(err)
	defer destination.Close()
    
	r, w := io.Pipe()
	go func() {
  	err := NewEncryptor(w).Encrypt(&source)
  	w.Close()
	DealWith(err)
	}()

	mr := io.MultiReader(
	 	strings.NewReader("fuck \r\n"),
	 	r,
	)
	io.Copy(destination, mr)

	return nil
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



type Settings struct {
  mode            string
  key             string
  keyfile         string
  iv              string
  ivfile          string
  plaintext       string
  plaintextfile   string
  ciphertext      string
  ciphertextfile  string
}


func getEncryptionParameters(config Settings) ( EncryptionParameters, error) {
  var params EncryptionParameters

  params.key = config.keyfile
  if params.key == "" {
    params.key = config.key
  }
  if params.key == "" {
    return params, Errors.NewAppError(Errors.NoKeyorKeyfile, "No Key or Key Provided")
  }
  params.plain = config.plaintextfile
  if params.plain == "" {
    params.plain = config.plaintext
  }
  if params.plain == "" {
    return params, Errors.NewAppError(Errors.NoPlaintextorPlaintextfile, "No Plaintext or Plaintextfile Provided")
  }

  params.cipher = config.ciphertextfile
  if params.cipher == "" {
    params.cipher = "stdout"
    println("No output file provided; defaulting to stdout")
  }

  switch config.mode {
    default:
      println("Using Counter Mode")
      params.mode = "ctr"
      break;
  }
  
  params.iv = config.ivfile
  if params.iv == "" {
    params.iv = config.iv
  }

  return params, nil
}

func getDecryptionParameters(config Settings) ( EncryptionParameters, error) {
  var params EncryptionParameters

  params.key = config.keyfile
  if params.key == "" {
    params.key = config.key
  }
  if params.key == "" {
    return params, Errors.NewAppError(Errors.NoKeyorKeyfile, "No Key or Key Provided")
  }
  params.cipher = config.ciphertextfile
  if params.cipher == "" {
    params.cipher = config.ciphertext
  }
  if params.cipher == "" {
    return params, Errors.NewAppError(Errors.NoCiphertextorCiphertextfile, "No Ciphertext or Ciphertextfile Provided")
  }

  params.plain = config.plaintextfile
  if params.plain == "" {
    params.plain = "stdout"
    println("No output file provided; defaulting to stdout")
  }

  switch config.mode {
    default:
      println("Using Counter Mode")
      params.mode = "ctr"
      break;
  }
  
  params.iv = config.ivfile
  if params.iv == "" {
    params.iv = config.iv
  }

  return params, nil
}


type EncryptionParameters struct {
  mode      string
  key       string
  iv        string
  plain     string
  cipher    string
}

