package AES

import (
    "github.com/urfave/cli"
    "github.com/SCKelemen/rex/Properties"
)

func Command(c *Properties.Context) cli.Command {
    return cli.Command{
        Name:     "aes",
      Subcommands: []cli.Command{
        {
          Name:  "encrypt",
          Usage: "Encrypts a file with AES",
          Category: "AES",
          Flags: GetEncryptionFlags(c),
          Action: func(c *cli.Context) error {
           println("encrypting...")

            return nil
          },
        },
        {
          Name:  "decrypt",
          Usage: "Decrypts a file with AES",
          Category: "AES",
          Flags: GetDecryptionFlags(c),
          Action: func(c *cli.Context) error {
            println("decrypting...")
            return nil
          },
        },	 
      },
    }


}



