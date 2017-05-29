package AES

import (
    "github.com/urfave/cli"
    "github.com/SCKelemen/rex/Properties"
)

func GetEncryptionFlags(context *Properties.Context) []cli.Flag {
   return []cli.Flag {
            cli.StringFlag{
                Name:        "mode, m",
                Value:       "ctr",
                Usage:       "Cipher Mode",
                Destination: &context.M.Value,
            },
            cli.StringFlag{
                Name:        "key, k",
                Value:       "",
                Usage:       "Encryption key",
                Destination: &context.K.Value,
            },
            cli.StringFlag{
                Name:        "iv",
                Value:       "",
                Usage:       "IV",
                Destination: &context.IV.Value,
            },
            cli.StringFlag{
                Name:        "plaintext, plain, p",
                Value:       "",
                Usage:       "Plaintext",
                Destination: &context.P.Value,
            },
            cli.StringFlag{
                Name:        "keyfile, kf",
                Value:       "",
                Usage:       "Encryption key file",
                Destination: &context.K.Filepath,
            },
            cli.StringFlag{
                Name:        "ivfile, ivf",
                Value:       "",
                Usage:       "IV",
                Destination: &context.IV.Filepath,
            },
            cli.StringFlag{
                Name:        "plaintextfile, plainfile, pf",
                Value:       "",
                Usage:       "Plaintextfile",
                Destination: &context.P.Filepath,
            },
            cli.StringFlag{
                Name:        "output, out, o",
                Value:       "",
                Usage:       "Output path",
                Destination: &context.K.Filepath,
            },
          }
}

func GetDecryptionFlags(context *Properties.Context) []cli.Flag {
   return []cli.Flag {
            cli.StringFlag{
                Name:        "mode, m",
                Value:       "ctr",
                Usage:       "Cipher Mode",
                Destination: &context.M.Value,
            },
            cli.StringFlag{
                Name:        "key, k",
                Value:       "",
                Usage:       "Encryption key",
                Destination: &context.K.Value,
            },
            cli.StringFlag{
                Name:        "iv",
                Value:       "",
                Usage:       "IV",
                Destination: &context.IV.Value,
            },
            cli.StringFlag{
                Name:        "plaintext, plain, p",
                Value:       "",
                Usage:       "Plaintext",
                Destination: &context.P.Value,
            },
            cli.StringFlag{
                Name:        "keyfile, kf",
                Value:       "",
                Usage:       "Encryption key file",
                Destination: &context.K.Filepath,
            },
            cli.StringFlag{
                Name:        "ivfile, ivf",
                Value:       "",
                Usage:       "IV",
                Destination: &context.IV.Filepath,
            },
            cli.StringFlag{
                Name:        "plaintextfile, plainfile, pf",
                Value:       "",
                Usage:       "Plaintextfile",
                Destination: &context.P.Filepath,
            },
            cli.StringFlag{
                Name:        "output, out, o",
                Value:       "",
                Usage:       "Output path",
                Destination: &context.K.Filepath,
            },
          }
}