package main

import (
  "os"

  "github.com/urfave/cli"
  "github.com/SCKelemen/rex/Flags/AES"
  "github.com/SCKelemen/rex/Commands/AES"
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
          Flags: Flags{}.GetEncryptionFlagsForAES(&config),
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


type Flags struct {

}



