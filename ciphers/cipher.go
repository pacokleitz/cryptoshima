package ciphers

type Cipher interface {
	Encrypt(mnemonic []string, secret string) ([]string, error)
	Decrypt(encryptedMnemonic []string, secret string) ([]string, error)
}
