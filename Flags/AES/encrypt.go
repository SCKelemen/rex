package Flags

import (
    "os"
    "github.com/urfave/cli"
    "github.com/SCKelemen/rex/cmd"
)



func (aes *AES) GetEncryptionFlags(config *Settings) []cli.Flag {
    return []cli.Flag {
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
          }
}