package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

var app *cli.App

func main() {

	app = cli.NewApp()
	app.Name = "Random stuffs"

	addCommands()

	app.Run(os.Args)
}

func addCommands() {

	app.Commands = []cli.Command{
		{
			Name:        "string",
			Aliases:     []string{"s"},
			Usage:       "Just generates a random str",
			Description: "",
			Action: func(c *cli.Context) {
				fmt.Println(randStr(10))
			},
		},
	}
}
