# Crypto Keys

Repo to take a BIP39 seed phrase and a BIP44 path and show the associated address and private key. This is useful so that you can generate the private keys used by Ledger hardware wallet, which is useful for back-up purposes.

## Set-up

Install the golang tool chain.

## How to run
On a Mac, in a command window:

```
export WORDS="health turn prison primary wave rule broken target alarm clinic border crunch era correct retreat mixture salmon evidence only cycle ask comfort science mixed"
go run src/cryptokeys/crypto.go "m/44'/60'/0'/0/0"
```

## Security
Don't disclose your private keys or seed phrase to anyone. Back up either off-line or to secure storage such as 1Passpord.


