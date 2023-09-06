package decrypt

import (
	"fmt"
	"strings"

	"github.com/pacokleitz/cryptoshima/ciphers"
	"github.com/urfave/cli/v2"
)

var DecryptPadCmd = &cli.Command{
	Name:   "pad",
	Usage:  "Decrypt an encrypted BIP-0039 mnemonic using a one-time-pad",
	Action: handleDecryptPadCommand,
}

func handleDecryptPadCommand(cCtx *cli.Context) error {
	encryptedMnemonic := strings.Split(cCtx.Args().Get(0), " ")
	encodedPad := cCtx.Args().Get(1)

	cipher := ciphers.NewOTPModuloCipher()
	mnemonic, err := cipher.Decrypt(encryptedMnemonic, encodedPad)
	if err != nil {
		return err
	}

	for idx, word := range mnemonic {
		fmt.Printf("[%d] - %s\n", idx, word)
	}
	fmt.Printf("%s\n", strings.Join(mnemonic, " "))
	return nil
}
