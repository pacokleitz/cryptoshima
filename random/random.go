package random

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

func init() {
	assertAvailablePRNG()
}

func assertAvailablePRNG() {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Sprintf("crypto/rand is unavailable: Read() failed with %#v", err))
	}
}

func GenerateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomInt64(upperBound int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(upperBound))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}
