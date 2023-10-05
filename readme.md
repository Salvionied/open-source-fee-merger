# Fee Merger for cardano

## Before Running
set Environment variables as:

MNEMONIC="your Seed Phrase"

FREQUENCY=every how many seconds it needs to run

BLOCKFROST_API_KEY=your mainnet blockfrost api key


## how to run:

`go run main.go`

This will create a folder named keys containing a wallet folder
within wallet folder which contains 6 files.

an address.txt containing the bech32 encoded address.
a mnemonic file containing the mnemonic to activate this wallet on the preferred nami or eternl wallet
and 4 key files containing stakes public and private, and payment public and private


## DockerHub

[https://hub.docker.com/repository/docker/zhaata/cardano-fee-merger/general](https://hub.docker.com/r/zhaata/cardano-fee-merger)https://hub.docker.com/r/zhaata/cardano-fee-merger
