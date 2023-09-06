package generate

import (
	"fmt"
	"github.com/pacokleitz/cryptoshima/pad"
	"github.com/urfave/cli/v2"
	"math"
)

var DEFAULT_MNEMONIC_LENGTH uint = 24

var PadCmd = &cli.Command{
	Name:  "pad",
	Usage: "Generate a one-time-pad encoded key",
	Flags: []cli.Flag{
		&cli.UintFlag{
			Name:  "words",
			Value: DEFAULT_MNEMONIC_LENGTH,
			Usage: "Number of words in the mnemonic phrase to encrypt",
			Action: func(ctx *cli.Context, v uint) error {
				if v > math.MaxUint16 {
					return fmt.Errorf("Flag words value %v out of range[0-%d]", v, math.MaxUint16)
				}
				return nil
			},
		},
	},
	Action: handleGenerateCommand,
}

func handleGenerateCommand(cCtx *cli.Context) error {
	randomPad, err := pad.GeneratePad(uint16(cCtx.Uint("words")))
	if err != nil {
		return err
	}

	encodedPad := pad.EncodePad(randomPad)

	_, wordList := pad.DecodePad(encodedPad)
	pad.PrintWordList(wordList)
	fmt.Printf("%s\n", encodedPad)
	return nil
}
