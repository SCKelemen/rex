package main

import (
  "os"

  "github.com/urfave/cli"
  "github.com/SCKelemen//rex/AES"
  "github.com/SCKelemen/rex/Flags/AES"
  "github.com/SCKelemen/rex/Commands/AES"
  "github.com/SCKelemen/rex/Properties"
)

func main() {
  app := cli.NewApp()
  var config Settings
  var context Properties.Context

  app.Commands = GetCommands(&context)

  app.Run(os.Args)
}


func GetCommands(c *Properties.Context) []cli.Command {
  commands := AES.Command(c)
  return commands
}