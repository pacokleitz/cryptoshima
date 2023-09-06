package encrypt

import (
	"fmt"
	"strings"

	"github.com/pacokleitz/cryptoshima/ciphers"
	"github.com/urfave/cli/v2"
)

var EncryptPadCmd = &cli.Command{
	Name:   "pad",
	Usage:  "Encrypt BIP-0039 mnemonic using a one-time-pad",
	Action: handleEncryptPadCommand,
}

func handleEncryptPadCommand(cCtx *cli.Context) error {
	mnemonic := strings.Split(cCtx.Args().Get(0), " ")
	encodedPad := cCtx.Args().Get(1)

	cipher := ciphers.NewOTPModuloCipher()
	encryptedMnemonic, err := cipher.Encrypt(mnemonic, encodedPad)
	if err != nil {
		return err
	}

	for idx, word := range encryptedMnemonic {
		fmt.Printf("[%d] - %s\n", idx, word)
	}

	fmt.Printf("%s\n", strings.Join(encryptedMnemonic, " "))
	return nil
}
