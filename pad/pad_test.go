package pad

import (
	"github.com/pacokleitz/cryptoshima/wordlists"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPadGeneration(t *testing.T) {
	const phraseLength = 24
	pad, err := GeneratePad(phraseLength)
	assert.Nil(t, err)
	encodedPad := EncodePad(pad)
	decodedPad, padWords := DecodePad(encodedPad)
	lenDecodedPad := len(decodedPad)

	for _, word := range padWords {
		assert.True(t, word < wordlists.Bip39Length)
	}

	assert.Equal(t, pad, decodedPad)
	assert.Equal(t, phraseLength, lenDecodedPad/2)
}
