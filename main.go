package main

import (
	"fmt"
	"os"
	
	"github.com/urfave/cli/v2"
)

func main() {
	app := App()

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func App() *cli.App {
	app := &cli.App {
		Name: "CLI AI",
		Usage: "Bertanya ke AI melalu CLI",
		Commands: []*cli.Command{
			{
				Name: "ask",
				Action: func (c *cli.Context) error{
					fmt.Println("Hello, World,CLI AI")
					return nil
				},
			},
		},
	}
	return app
}