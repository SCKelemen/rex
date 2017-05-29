package main

import (
	"fmt"
	"os"
	"bytes"
	"io"
	"strings"
	"github.com/urfave/cli"

	 
   router "github.com/SCKelemen/rex/router"
)

func main() {
	app := cli.NewApp()
  app.Name = "rex"
  r := router.NewRouter(nil)


 	app.Commands = []cli.Command{
	{
      Name:        "aes",
      Usage:       "aes",
      Subcommands: []cli.Command{
        {
          Name:  "encrypt",
          Usage: "Encrypts a files with AES",
          Action: func(c *cli.Context) error {
            k := []byte("example key 1234")
            cip := r.NewAESRouter().CTR.Encrypt([]byte("hello string"), k)
             p := r.NewAESRouter().CTR.Decrypt(cip, k)
           //plain := r.NewAESRouter().CTR.Decrypt(cipher, []byte("example key 1234"))
             fmt.Printf("cipher: %s\n", cip)
              fmt.Printf("plain: %s\n", p)

  
            return nil
          },
        },
        {
          Name:  "decrypt",
          Usage: "Decrypts",
          Action: func(c *cli.Context) error {
            fmt.Println("decrypting file: ", c.Args().First())
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



