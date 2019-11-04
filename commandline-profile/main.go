package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "profile"
	app.Usage = "Browse my command line profile"
	
  app.Action = func(c *cli.Context) error {
    fmt.Println("My commandline profile")
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}