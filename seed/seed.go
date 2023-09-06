package seed

import (
	"encoding/binary"
	"fmt"
	"github.com/pacokleitz/cryptoshima/wordlists"
)

const MnemonicPhraseLength uint16 = 24
const Uint16ByteSize = 2

func Uint16ToBytes(value uint16) []byte {
	bs := make([]byte, Uint16ByteSize)
	binary.BigEndian.PutUint16(bs, value)
	return bs
}

func BytesToUint(bytes []byte) uint16 {
	return binary.BigEndian.Uint16(bytes)
}

// Return the index/value of a BIP-0039 word (TODO: optimize by using a word to value hashmap)
func WordToUInt16(word string) (uint16, error) {
	for index, w := range wordlists.Bip39Words {
		if w == word {
			return uint16(index), nil
		}
	}
	return 0, fmt.Errorf("%s is not a valid BIP-0039 word", word)
}

// Return the BIP-0039 word at a specific index
func Uint16ToWord(index uint16) (string, error) {
	if index >= wordlists.Bip39Length {
		return "", fmt.Errorf("%d is not a valid BIP-0039 index number", index)
	}
	return wordlists.Bip39Words[index], nil
}

func MnemonicToBytes(mnemonic []string) ([]byte, error) {
	values := []uint16{}
	for _, word := range mnemonic {
		value, err := WordToUInt16(word)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}

	bytes := []byte{}
	for _, value := range values {
		bytes = append(bytes, Uint16ToBytes(value)...)
	}
	return bytes, nil
}

func MnemonicToUint16Slice(mnemonic []string) ([]uint16, error) {
	values := []uint16{}
	for _, word := range mnemonic {
		value, err := WordToUInt16(word)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

func Uint16SliceToBytes(values []uint16) ([]byte, error) {
	bytes := []byte{}
	for _, value := range values {
		bytes = append(bytes, Uint16ToBytes(value)...)
	}
	return bytes, nil
}

func Uint16SliceToMnemonic(values []uint16) ([]string, error) {
	words := []string{}
	for _, value := range values {
		word, err := Uint16ToWord(value)
		if err != nil {
			return nil, err
		}
		words = append(words, word)
	}
	return words, nil
}

func BytesToMnemonic(bytes []byte) ([]string, error) {
	bytesLength := len(bytes)
	if bytesLength != int(MnemonicPhraseLength)*2 {
		return nil, fmt.Errorf("%d bytes array can't be converted to %d words mnemonic", bytesLength, MnemonicPhraseLength)
	}

	valuesFromBytes := []uint16{}
	for i := 0; i < len(bytes); i += 2 {
		valueFromBytes := BytesToUint(bytes[i:])
		valuesFromBytes = append(valuesFromBytes, valueFromBytes)
	}

	wordsFromValues := []string{}
	for _, value := range valuesFromBytes {
		word, err := Uint16ToWord(value)
		if err != nil {
			return nil, err
		}
		wordsFromValues = append(wordsFromValues, word)
	}
	return wordsFromValues, nil
}
