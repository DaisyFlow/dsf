package main

import (
	"fmt"
	"os"

	"github.com/daisyflow/dsf/cmd/dsf/app"
)

func main() {
	command := app.NewDsfCommand()
	if err := command.Run(os.Args); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
