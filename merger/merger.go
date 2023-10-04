package merger

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/Salvionied/apollo"
	"github.com/Salvionied/apollo/serialization/Address"
	"github.com/Salvionied/apollo/serialization/Key"
	"github.com/Salvionied/apollo/serialization/UTxO"
	"github.com/Salvionied/apollo/txBuilding/Backend/BlockFrostChainContext"
	"github.com/joho/godotenv"
)

type Wallet struct {
	SigningKey           Key.SigningKey
	VerificationKey      Key.VerificationKey
	StakeVerificationKey Key.StakeVerificationKey
	StakeSigningKey      Key.StakeSigningKey
	Address              Address.Address
}

func NewWallet(walletName string) Wallet {
	skey, vkey, svkey, sskey, addr := load_keys(walletName)
	return Wallet{
		SigningKey:           skey,
		VerificationKey:      vkey,
		StakeVerificationKey: svkey,
		StakeSigningKey:      sskey,
		Address:              addr,
	}
}

type Merger struct {
	Wallet Wallet
	Bfc    BlockFrostChainContext.BlockFrostChainContext
}

func NewMerger(walletName string) Merger {
	wallet := NewWallet(walletName)
	godotenv.Load()
	blockfrostApiKey := os.Getenv("BLOCKFROST_API_KEY")
	backend, _ := apollo.NewBlockfrostBackend(blockfrostApiKey, apollo.MAINNET)
	return Merger{
		Wallet: wallet,
		Bfc:    backend,
	}
}

func (m Merger) Merge() {
	utxos := m.Bfc.Utxos(m.Wallet.Address)
	//split utxos in chunks of 400 utxos
	if len(utxos) < 2 {
		fmt.Println("Nothing to merge")
		return
	}
	chunks := make([][]UTxO.UTxO, 0)
	for i := 0; i < len(utxos); i += 400 {
		end := i + 400
		if end > len(utxos) {
			end = len(utxos)
		}
		chunks = append(chunks, utxos[i:end])
	}
	for _, chunk := range chunks {
		apollob := apollo.New(&m.Bfc)
		apollob = apollob.AddInput(chunk...).SetWalletFromBech32(m.Wallet.Address.String()).SetWalletAsChangeAddress()
		completed, err := apollob.Complete()
		if err != nil {
			continue
		}
		completed.SignWithSkey(m.Wallet.VerificationKey, m.Wallet.SigningKey)
		tx := completed.GetTx()
		if tx == nil {
			continue
		}
		txid, err := m.Bfc.SubmitTx(*tx)
		if err != nil {
			continue
		}
		fmt.Println("MERGED", len(chunk), "Into", hex.EncodeToString(txid.Payload))
	}

}

func (m Merger) Loop() {
	for {
		m.Merge()
		time.Sleep(60 * time.Second)

	}
}
