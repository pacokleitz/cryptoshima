package ciphers

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var validEncodedPads = []string{
	"ABgGeQNSBhoDeQcuBWYBxADFA4gFuwTgASQFnge7B3oBUgTPAdEDvgEdA_gDKgAiB0dkuIBo",
	"ABgEWwNjAYQBiQaGBgoBgwBCBrwEVQc2BN0AugfdAoAHfAZvB4sHygGOBX0GGgDJAF34XvXc",
	"ABgA_ACIA4MG7ASQAXUGNAcEAkUGpQDLBwUHvAFxBvAArgf2AT4EswHSAsEF2QbVBIJp5Zm0",
	"ABgAXAKhAN4BlQTzAAQAJgUeApYBoAUtBjID8gKdBMYB9QLbAI4CxgA3ACABFAWNAZA5z0c9",
	"ABgE3APOAJgD9wXBBiQAGgRwAucBXwQnBK4AWAURBFIEkwEHBAcAggGpAk8D2QLpBA6fp3_Q",
	"ABgG2Qe7AK8CIwdwAnQAuAahBbwAmgGIBhEDxQTsAckG-QMmBJIH_AIoA8UA4wOEA5JZjbfi",
	"ABgH6QZqBm0FgAWTBvgBdwMjBygGxQGxB9gAGgSPAkMCNAXaAUoBXgRpANYGgAetBsuYjy_D",
	"ABgEDwMJBhMGVQc_A48CiAaCAWIH4QJgB0MEMgGTBAsDbgFmB4kGqAW4A_4AegGpB8sapFdq",
}

func TestEncryptDecryptOTPModulo(t *testing.T) {
	cipher := NewOTPModuloCipher()

	for _, mnemonic := range validMnemonics {
		splittedMnemonic := strings.Split(mnemonic, " ")
		for _, pad := range validEncodedPads {
			encryptedMnemonic, err := cipher.Encrypt(splittedMnemonic, pad)
			assert.Nil(t, err)
			assert.NotEqual(t, splittedMnemonic, encryptedMnemonic)
			decryptedMnemonic, err := cipher.Decrypt(encryptedMnemonic, pad)
			assert.Nil(t, err)
			assert.Equal(t, splittedMnemonic, decryptedMnemonic)
		}
	}
}
