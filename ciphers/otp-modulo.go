package ciphers

import (
	"fmt"

	"github.com/pacokleitz/cryptoshima/math"
	"github.com/pacokleitz/cryptoshima/pad"
	"github.com/pacokleitz/cryptoshima/seed"
	"github.com/pacokleitz/cryptoshima/wordlists"
)

type OTPModuloCipher struct{}

func NewOTPModuloCipher() *OTPModuloCipher {
	return &OTPModuloCipher{}
}

func (*OTPModuloCipher) Encrypt(mnemonic []string, encodedPad string) ([]string, error) {
	_, pad := pad.DecodePad(encodedPad)
	uint16Mnemonic, err := seed.MnemonicToUint16Slice(mnemonic)
	if err != nil {
		return nil, err
	}

	mnemonicLength, padLength := len(uint16Mnemonic), len(pad)

	if mnemonicLength != padLength {
		return nil, fmt.Errorf("%d words mnemonic can't be encrypted by %d words pad", mnemonicLength, padLength)
	}

	for i := 0; i < mnemonicLength; i++ {
		uint16Mnemonic[i] = math.Mod((uint16Mnemonic[i] + pad[i]), wordlists.Bip39Length)
	}

	encryptedMnemonic, err := seed.Uint16SliceToMnemonic(uint16Mnemonic)
	if err != nil {
		return nil, err
	}

	return encryptedMnemonic, nil
}

func (*OTPModuloCipher) Decrypt(encryptedMnemonic []string, encodedPad string) ([]string, error) {
	_, pad := pad.DecodePad(encodedPad)
	uint16EncryptedMnemonic, err := seed.MnemonicToUint16Slice(encryptedMnemonic)
	if err != nil {
		return nil, err
	}

	mnemonicLength, padLength := len(uint16EncryptedMnemonic), len(pad)

	if mnemonicLength != padLength {
		return nil, fmt.Errorf("%d words mnemonic can't be encrypted by %d words pad", mnemonicLength, padLength)
	}

	for i := 0; i < mnemonicLength; i++ {
		uint16EncryptedMnemonic[i] = math.Mod((uint16EncryptedMnemonic[i] - pad[i]), wordlists.Bip39Length)
	}

	mnemonic, err := seed.Uint16SliceToMnemonic(uint16EncryptedMnemonic)
	if err != nil {
		return nil, err
	}

	return mnemonic, nil
}
