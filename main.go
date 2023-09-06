package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/pacokleitz/cryptoshima/cli/commands"
)

func main() {
	app := &cli.App{
		Name:  "cryptoshima",
		Usage: "Make your recovery seed survive anything 💣",
		Commands: []*cli.Command{
			commands.GenerateCmd,
			commands.CheckCmd,
			commands.EncryptCmd,
			commands.DecryptCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
