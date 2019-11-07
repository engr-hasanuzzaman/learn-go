package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli"
)

func main() {
  var language string

  app := cli.NewApp()

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name:        "lang",
      Value:       "english",
      Usage:       "language for the greeting",
      // Destination: &language,
    },
    {
      Name:        "lang",
      Value:       "bn",
      Usage:       "language for the greeting",
      // Destination: &language,
    },
  }

  app.Action = func(c *cli.Context) error {
    name := "someone"
    language = c.String("lang")
    if c.NArg() > 0 {
      name = c.Args()[0]
    }

    if language == "bd" {
      fmt.Println("Apnake sagotom ", name)
    }else if language == "spanish" {
      fmt.Println("Hola", name)
    } else {
      fmt.Println("Hello", name)
    }
    
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}