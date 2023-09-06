package check

import (
	"github.com/urfave/cli/v2"
	"github.com/pacokleitz/cryptoshima/pad"
)

var PadCmd = &cli.Command{
	Name:  "pad",
	Usage: "Check the BIP-0039 words and validity of a one-time-pad encoded key",
	Action: handleCheckPadCommand,
}

func handleCheckPadCommand(cCtx *cli.Context) error {
	encodedPad := cCtx.Args().Get(0)
	_, wordList := pad.DecodePad(encodedPad)
	pad.PrintWordList(wordList)
	return nil
}
