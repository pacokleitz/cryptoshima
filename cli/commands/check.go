package commands

import (
	"github.com/pacokleitz/cryptoshima/cli/subcommands/check"
	"github.com/urfave/cli/v2"
)

var CheckCmd = &cli.Command{
	Name:        "check",
	Usage:       "Check the validity of a resource",
	Subcommands: []*cli.Command{check.PadCmd},
}
