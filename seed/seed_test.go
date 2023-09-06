package seed

import (
	"testing"

	"github.com/pacokleitz/cryptoshima/math"
	"github.com/pacokleitz/cryptoshima/wordlists"
	"github.com/stretchr/testify/assert"
)

type validWord struct {
	word  string
	value uint16
}

var validMnemonics = [][]validWord{
	{{"fault", 671}, {"network", 1190}, {"forum", 733}, {"supreme", 1743}, {"brisk", 226}, {"alcohol", 48}, {"science", 1543}, {"rose", 1504}, {"artwork", 105}, {"foam", 720}, {"question", 1404}, {"dove", 526}, {"pluck", 1333}, {"maximum", 1099}, {"baby", 136}, {"blood", 191}, {"asset", 109}, {"burden", 244}, {"south", 1665}, {"replace", 1462}, {"shrimp", 1593}, {"decline", 455}, {"police", 1341}, {"coach", 357}},
}

func TestValidWordToUint(t *testing.T) {
	for _, mnemonic := range validMnemonics {
		for _, testWord := range mnemonic {
			value, err := WordToUInt16(testWord.word)
			assert.Nil(t, err)
			assert.Equal(t, value, testWord.value)
		}
	}
}

func TestValidUintToWord(t *testing.T) {
	for _, mnemonic := range validMnemonics {
		for _, testWord := range mnemonic {
			word, err := Uint16ToWord(testWord.value)
			assert.Nil(t, err)
			assert.Equal(t, word, testWord.word)
		}
	}
}

func TestMnemonicToMnemonic(t *testing.T) {
	for _, validMnemonic := range validMnemonics {
		mnemonicBytes, err := MnemonicToBytes(validWordsToWords(validMnemonic))
		assert.Nil(t, err)

		extractedMnemonic, err := BytesToMnemonic(mnemonicBytes)
		assert.Nil(t, err)

		for i := 0; i < len(extractedMnemonic); i++ {
			assert.Equal(t, extractedMnemonic[i], validMnemonic[i].word)
		}
	}
}

func TestMnemonicToMnemonicWithModuloPad(t *testing.T) {
	for _, validMnemonic := range validMnemonics {
		mnemonicValues, err := MnemonicToUint16Slice(validWordsToWords(validMnemonic))
		assert.Nil(t, err)

		for i := 0; i < len(mnemonicValues); i++ {
			mnemonicValues[i] = math.Mod((mnemonicValues[i] + 1984), wordlists.Bip39Length)
		}

		mnemonicBytes, err := Uint16SliceToBytes(mnemonicValues)
		assert.Nil(t, err)

		extractedMnemonic, err := BytesToMnemonic(mnemonicBytes)
		assert.Nil(t, err)

		for i := 0; i < len(extractedMnemonic); i++ {
			targetPaddedWord, err := Uint16ToWord(math.Mod((validMnemonic[i].value + 1984), uint16(wordlists.Bip39Length)))
			assert.Nil(t, err)
			assert.Equal(t, extractedMnemonic[i], targetPaddedWord)
		}
	}
}

func validWordsToWords(validWords []validWord) []string {
	words := []string{}
	for _, validWord := range validWords {
		words = append(words, validWord.word)
	}
	return words
}
