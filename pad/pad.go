package pad

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/pacokleitz/cryptoshima/random"
	"github.com/pacokleitz/cryptoshima/seed"
)

const DICTIONNARY_SIZE = 2048
const PAD_CHECKSUM_BYTES = 4
const PAD_SIZE_BYTES = 2
const WORD_BYTES = 2

func GeneratePad(nbWords uint16) ([]byte, error) {
	randomBytes := []byte{}
	for i := 0; i < int(nbWords); i++ {
		randomValue := uint16(random.GenerateRandomInt64(DICTIONNARY_SIZE))
		randomBytes = append(randomBytes, seed.Uint16ToBytes(randomValue)...)
	}
	return randomBytes, nil
}

func EncodePad(pad []byte) string {
	padBytes := seed.Uint16ToBytes(seed.MnemonicPhraseLength)
	padBytes = append(padBytes, pad...)
	checkSum := sha256.Sum256(padBytes)
	padBytes = append(padBytes, checkSum[0:PAD_CHECKSUM_BYTES]...)
	encodedPad := base64.RawURLEncoding.EncodeToString(padBytes)
	return encodedPad
}

func DecodePad(encodedPad string) ([]byte, []uint16) {
	pad, err := base64.RawURLEncoding.DecodeString(encodedPad)
	if err != nil {
		panic("error:")
	}

	padLength := len(pad)
	padChecksum := pad[padLength-PAD_CHECKSUM_BYTES:]
	computedCheckSum := sha256.Sum256(pad[:padLength-PAD_CHECKSUM_BYTES])

	if !bytes.Equal(padChecksum, computedCheckSum[:PAD_CHECKSUM_BYTES]) {
		panic("error")
	}

	phraseLength := seed.BytesToUint(pad[0:PAD_SIZE_BYTES])
	if phraseLength != (uint16(padLength)-(PAD_CHECKSUM_BYTES+PAD_SIZE_BYTES))/2 {
		panic("error")
	}

	wordList := []uint16{}
	for i := 0; i < int(phraseLength); i++ {
		wordIndex := seed.BytesToUint(pad[PAD_SIZE_BYTES+(i*WORD_BYTES):])
		wordList = append(wordList, wordIndex)
	}
	return pad[PAD_SIZE_BYTES : padLength-PAD_CHECKSUM_BYTES], wordList
}

func PrintWordList(wordList []uint16) {
	for idx, wordIndex := range wordList {
		word, err := seed.Uint16ToWord(wordIndex)
		if err != nil {
			panic("error")
		}
		fmt.Printf("[%d] - %s (%d)\n", idx, word, wordIndex)
	}
}
