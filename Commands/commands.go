package Commands

import (
    "github.com/urfave/cli"
)

type AES struct {
    Encryption Crypt
    Decryption Crypt
}
type Crypt struct {
    Flags []cli.Flag
    Commands []cli.Command
}

func GetAESProperties(config *Settings) AES {
   encryption := Crypt{Flags: []cli.Flag {
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
          Commands: []cli.Command{},
    }
    return AES{Encryption: encryption, Decryption: encryption}

}

type Settings struct {
    k flagValue
    iv flagValue
    p flagValue
    c flagValue
    m flagValue
}

type flagValue struct {
    value string
    filepath string
    isFile bool
}