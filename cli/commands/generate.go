package commands

import (
	"github.com/pacokleitz/cryptoshima/cli/subcommands/generate"
	"github.com/urfave/cli/v2"
)

var GenerateCmd = &cli.Command{
	Name:        "generate",
	Usage:       "Generate a new resource",
	Subcommands: []*cli.Command{generate.PadCmd},
}
