package main

import (
  "os"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  var config Settings

  app.Commands = []cli.Command{
    {
      Name:     "aes",
      Subcommands: []cli.Command{
        {
          Name:  "encrypt",
          Usage: "Encrypts a file with AES",
          Category: "AES",
          Flags: []cli.Flag {
            cli.StringFlag{
                Name:        "mode, m",
                Value:       "ctr",
                Usage:       "Cipher Mode",
                Destination: &config.m.value,
            },
            cli.StringFlag{
                Name:        "key, k",
                Value:       "",
                Usage:       "Encryption key",
                Destination: &config.k.value,
            },
            cli.StringFlag{
                Name:        "iv",
                Value:       "",
                Usage:       "IV",
                Destination: &config.iv.value,
            },
            cli.StringFlag{
                Name:        "plaintext, plain, p",
                Value:       "",
                Usage:       "Plaintext",
                Destination: &config.p.value,
            },
            cli.StringFlag{
                Name:        "keyfile, kf",
                Value:       "",
                Usage:       "Encryption key file",
                Destination: &config.k.filepath,
            },
            cli.StringFlag{
                Name:        "ivfile, ivf",
                Value:       "",
                Usage:       "IV",
                Destination: &config.iv.filepath,
            },
            cli.StringFlag{
                Name:        "plaintextfile, plainfile, pf",
                Value:       "",
                Usage:       "Plaintextfile",
                Destination: &config.p.filepath,
            },
            cli.StringFlag{
                Name:        "output, out, o",
                Value:       "",
                Usage:       "Output path",
                Destination: &config.k.filepath,
            },
          },
          Action: func(c *cli.Context) error {
           println("encrypting...")
           println(config.k.value)
            return nil
          },
        },
        {
          Name:  "decrypt",
          Usage: "Decrypts a file with AES",
          Category: "AES",
          Action: func(c *cli.Context) error {
            println("decrypting...")
            return nil
          },
        },	 
      },
    },
  }

  app.Run(os.Args)
}

type flagValue struct {
    value string
    filepath string
    isFile bool
}

type Settings struct {
    k flagValue
    iv flagValue
    p flagValue
    c flagValue
    m flagValue
}


