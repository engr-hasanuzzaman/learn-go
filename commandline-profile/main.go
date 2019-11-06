package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()

  app.Action = func(c *cli.Context) error {
    fmt.Println("All the arguments passed as command argument are")
    for i := 0; i < len(c.Args()); i++ {
      fmt.Printf("%d th value is %s \n", i, c.Args()[i])  
    }
    // fmt.Printf("Hello %q", c.Args().Get(0))
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}