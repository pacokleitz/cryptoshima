package commands

import (
	"github.com/pacokleitz/cryptoshima/cli/subcommands/encrypt"
	"github.com/urfave/cli/v2"
)

var EncryptCmd = &cli.Command{
	Name:        "encrypt",
	Usage:       "Encrypt a BIP-0039 mnemonic",
	Subcommands: []*cli.Command{encrypt.EncryptPadCmd},
}
