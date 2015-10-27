package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "random"
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

	app.Run(os.Args)
}
