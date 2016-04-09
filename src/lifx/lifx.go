package main

import (
	"github.com/codegangsta/cli"
	"lifx/client"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Name = "Lifx"
	app.Usage = "lifx [command] [...params]"
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "debug, d", Usage: "Enable debugging"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "config",
			Usage: "Gets the configuration",
			Action: func(c *cli.Context) {
				client.Config()
			},
		},
		{
			Name:  "select",
			Usage: "Select a light",
			Action: func(c *cli.Context) {
				var selector string
				if c.NArg() > 0 {
					selector = c.Args()[0]
				}
				debug := c.GlobalBool("debug")
				client.SelectLight(debug, selector)
			},
		},
		{
			Name:  "list",
			Usage: "List all bulbs",
			Action: func(c *cli.Context) {
				debug := c.GlobalBool("debug")
				client.List(debug)
			},
		},
		{
			Name:  "toggle",
			Usage: "Toggle selected/all bulbs",
			Action: func(c *cli.Context) {
				debug := c.GlobalBool("debug")
				client.Toggle(debug)
			},
		},
		{
			Name:  "on",
			Usage: "Turn light(s) on",
			Action: func(c *cli.Context) {
				debug := c.GlobalBool("debug")
				client.On(debug)
			},
		},
		{
			Name:  "off",
			Usage: "Turn light(s) off",
			Action: func(c *cli.Context) {
				debug := c.GlobalBool("debug")
				client.Off(debug)
			},
		},
		{
			Name:  "brightness",
			Usage: "Change the brightness of a light",
			Action: func(c *cli.Context) {
				var brightness string
				if c.NArg() > 0 {
					brightness = c.Args()[0]
					debug := c.GlobalBool("debug")
					client.Brightness(debug, brightness)
				} else {
					println("Enter brighness level 0.0 - 1.0")
				}
			},
		},
		{
			Name:  "color",
			Usage: "Change the color of a light",
			Action: func(c *cli.Context) {
				var color string
				if c.NArg() > 0 {
					color = c.Args()[0]
					debug := c.GlobalBool("debug")
					client.Color(debug, color)
				} else {
					println("Enter color string")
				}
			},
		},
	}
	app.Run(os.Args)
}
