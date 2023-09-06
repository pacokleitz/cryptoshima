package commands

import (
	"github.com/pacokleitz/cryptoshima/cli/subcommands/decrypt"
	"github.com/urfave/cli/v2"
)

var DecryptCmd = &cli.Command{
	Name:        "decrypt",
	Usage:       "Decrypt an encrypted BIP-0039 mnemonic",
	Subcommands: []*cli.Command{decrypt.DecryptPadCmd},
}
