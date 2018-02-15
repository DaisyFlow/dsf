package app

import "github.com/urfave/cli"

// NewDsfCommand returns the cli App used to run the application.
func NewDsfCommand() *cli.App {
	app := cli.NewApp()
	app.Name = "dsf"
	app.Usage = "Manage code"

	return app
}
