<div align="center">
  <h1>cryptoshima</h1>
  <img src="https://github.com/pacokleitz/cryptoshima/assets/31453761/1d3712d7-a5cd-43ca-8a4f-7b17258c899c" width="300" />
  <p>Render your seed phrase unbreakable 💣</p>
</div>

## What is this ?

Cryptoshima is a command-line tool meticulously crafted to bolster the security of your BIP-0039 recovery seed phrase. Keeping your seed phrase unencrypted in cold storage can pose significant risks. Leveraging one-time-pad encryption, Cryptoshima empowers you to render your seed phrase completely unreadable to anyone but yourself. It's highly advisable to safeguard the encrypted seed phrase on a physical medium while ensuring the one-time-pad is securely stored in multiple online storage locations.

**Use at your own risk.**

## Commands

- **generate**: Create new resources, including an encoded one-time-pad that can be used for encryption.
- **check**: Verify the validity of your resources, such as encoded one-time-pads, output the decoded pad.
- **encrypt**: Encrypt your BIP-0039 mnemonic, includes a one-time-pad option.
- **decrypt**: Decrypt an encrypted BIP-0039 mnemonic, with a one-time-pad decryption option.

For more usage instructions, run `./cryptoshima help [command]`.
