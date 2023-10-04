package merger

import (
	"fmt"

	"log"
	"os"

	"github.com/Salvionied/apollo/serialization/Address"
	"github.com/Salvionied/apollo/serialization/HDWallet"
	"github.com/Salvionied/apollo/serialization/Key"
)

func CreateWallet(folder_name string) (Key.SigningKey, Key.VerificationKey, Key.StakeVerificationKey, Key.StakeSigningKey, Address.Address) {
	fmt.Println("Directories do not exist, Generating new ones")
	if _, err := os.Stat("keys" + "/" + folder_name); !os.IsNotExist(err) {
		log.Fatal("Wallet already exists")
	}
	os.Mkdir("keys"+"/"+folder_name, 0755)

	mnemonic, _ := HDWallet.GenerateMnemonic()
	os.WriteFile("keys"+"/"+folder_name+"/"+"mnemonic", []byte(mnemonic), 0644)
	hdWall, _ := HDWallet.NewHDWalletFromMnemonic(mnemonic, "")
	paymentPath := "m/1852'/1815'/0'/0/0"
	stakingPath := "m/1852'/1815'/0'/2/0"
	paymentKeyPath, _ := hdWall.DerivePath(paymentPath)
	verificationKey_bytes := paymentKeyPath.XPrivKey.PublicKey()
	signingKey_bytes := paymentKeyPath.XPrivKey.Bytes()
	stakingKeyPath, _ := hdWall.DerivePath(stakingPath)
	stakeVerificationKey_bytes := stakingKeyPath.XPrivKey.PublicKey()
	stakeSigningKey_bytes := stakingKeyPath.XPrivKey.Bytes()
	signingKey := Key.SigningKey{signingKey_bytes}
	verificationKey := Key.VerificationKey{verificationKey_bytes}
	stakeSigningKey := Key.StakeSigningKey{stakeSigningKey_bytes}
	stakeVerificationKey := Key.StakeVerificationKey{stakeVerificationKey_bytes}
	stakeVerKey := Key.VerificationKey{Payload: stakeVerificationKey_bytes}
	skh, _ := stakeVerKey.Hash()
	vkh, _ := verificationKey.Hash()

	os.WriteFile("keys"+"/"+folder_name+"/"+"public.key", verificationKey.Payload, 0644)
	os.WriteFile("keys"+"/"+folder_name+"/"+"private.key", signingKey.Payload, 0644)
	os.WriteFile("keys"+"/"+folder_name+"/"+"stake.public.key", stakeSigningKey.Payload, 0644)
	os.WriteFile("keys"+"/"+folder_name+"/"+"stake.private.key", stakeVerificationKey.Payload, 0644)
	addr := Address.Address{StakingPart: skh[:], PaymentPart: vkh[:], Network: 1, AddressType: Address.KEY_KEY, HeaderByte: 0b00000001, Hrp: "addr"}
	os.WriteFile("keys"+"/"+folder_name+"/"+"address.txt", []byte(addr.String()), 0644)
	return signingKey, verificationKey, stakeVerificationKey, stakeSigningKey, addr

}

func LoadWallet(folder_name string) (Key.SigningKey, Key.VerificationKey, Key.StakeVerificationKey, Key.StakeSigningKey, Address.Address) {
	signingKey_bytes, err := os.ReadFile("keys" + "/" + folder_name + "/" + "private.key")
	if err != nil {
		log.Fatal(err)
	}
	verificationKey_bytes, err := os.ReadFile("keys" + "/" + folder_name + "/" + "public.key")
	if err != nil {
		log.Fatal(err)
	}
	stakeSigningKey_bytes, err := os.ReadFile("keys" + "/" + folder_name + "/" + "stake.private.key")
	if err != nil {
		log.Fatal(err)
	}
	stakeVerificationKey_bytes, err := os.ReadFile("keys" + "/" + folder_name + "/" + "stake.public.key")
	if err != nil {
		log.Fatal(err)
	}
	seed_bytes, err := os.ReadFile("keys" + "/" + folder_name + "/" + "mnemonic")
	if err != nil {
		signingKey := Key.SigningKey{Payload: signingKey_bytes}
		verificationKey := Key.VerificationKey{Payload: verificationKey_bytes}
		stakeSigningKey := Key.StakeSigningKey{Payload: stakeSigningKey_bytes}
		stakeVerificationKey := Key.StakeVerificationKey{Payload: stakeVerificationKey_bytes}
		stakeVerKey := Key.VerificationKey{Payload: stakeVerificationKey_bytes}
		skh, _ := stakeVerKey.Hash()
		vkh, _ := verificationKey.Hash()

		addr := Address.Address{StakingPart: skh[:], PaymentPart: vkh[:], Network: 1, AddressType: Address.KEY_KEY, HeaderByte: 0b00000001, Hrp: "addr"}
		return signingKey, verificationKey, stakeVerificationKey, stakeSigningKey, addr

	}
	paymentPath := "m/1852'/1815'/0'/0/0"
	stakingPath := "m/1852'/1815'/0'/2/0"
	hdWall, _ := HDWallet.NewHDWalletFromMnemonic(string(seed_bytes), "")
	paymentKeyPath, _ := hdWall.DerivePath(paymentPath)
	verificationKey_bytes = paymentKeyPath.XPrivKey.PublicKey()
	signingKey_bytes = paymentKeyPath.XPrivKey.Bytes()
	stakingKeyPath, _ := hdWall.DerivePath(stakingPath)
	stakeVerificationKey_bytes = stakingKeyPath.XPrivKey.PublicKey()
	stakeSigningKey_bytes = stakingKeyPath.XPrivKey.Bytes()
	//stake := stakingKeyPath.RootXprivKey.Bytes()
	signingKey := Key.SigningKey{Payload: signingKey_bytes}
	verificationKey := Key.VerificationKey{Payload: verificationKey_bytes}
	stakeSigningKey := Key.StakeSigningKey{Payload: stakeSigningKey_bytes}
	stakeVerificationKey := Key.StakeVerificationKey{Payload: stakeVerificationKey_bytes}
	stakeVerKey := Key.VerificationKey{Payload: stakeVerificationKey_bytes}
	skh, _ := stakeVerKey.Hash()
	vkh, _ := verificationKey.Hash()

	addr := Address.Address{StakingPart: skh[:], PaymentPart: vkh[:], Network: 1, AddressType: Address.KEY_KEY, HeaderByte: 0b00000001, Hrp: "addr"}

	return signingKey, verificationKey, stakeVerificationKey, stakeSigningKey, addr

}

func load_wallet_from_mnemonic(mnemonic string) (Key.SigningKey, Key.VerificationKey, Key.StakeVerificationKey, Key.StakeSigningKey, Address.Address) {
	hdWall, _ := HDWallet.NewHDWalletFromMnemonic(mnemonic, "")
	paymentPath := "m/1852'/1815'/0'/0/0"
	stakingPath := "m/1852'/1815'/0'/2/0"
	paymentKeyPath, _ := hdWall.DerivePath(paymentPath)
	verificationKey_bytes := paymentKeyPath.XPrivKey.PublicKey()
	signingKey_bytes := paymentKeyPath.XPrivKey.Bytes()
	stakingKeyPath, _ := hdWall.DerivePath(stakingPath)
	stakeVerificationKey_bytes := stakingKeyPath.XPrivKey.PublicKey()
	stakeSigningKey_bytes := stakingKeyPath.XPrivKey.Bytes()
	signingKey := Key.SigningKey{signingKey_bytes}
	verificationKey := Key.VerificationKey{verificationKey_bytes}
	stakeSigningKey := Key.StakeSigningKey{stakeSigningKey_bytes}
	stakeVerificationKey := Key.StakeVerificationKey{stakeVerificationKey_bytes}
	stakeVerKey := Key.VerificationKey{Payload: stakeVerificationKey_bytes}
	skh, _ := stakeVerKey.Hash()
	vkh, _ := verificationKey.Hash()
	addr := Address.Address{StakingPart: skh[:], PaymentPart: vkh[:], Network: 1, AddressType: Address.KEY_KEY, HeaderByte: 0b00000001, Hrp: "addr"}
	return signingKey, verificationKey, stakeVerificationKey, stakeSigningKey, addr
}

// func ImportWallet(name string) (Key.SigningKey, Key.VerificationKey, Key.StakeVerificationKey, Key.StakeSigningKey, Address.Address) {
// 	return signingKey, verificationKey, stakeVerificationKey, stakeSigningKey, addr

// }

func load_keys(folder_name string) (Key.SigningKey, Key.VerificationKey, Key.StakeVerificationKey, Key.StakeSigningKey, Address.Address) {
	var signingKey Key.SigningKey
	var verificationKey Key.VerificationKey
	var stakeSigningKey Key.StakeSigningKey
	var stakeVerificationKey Key.StakeVerificationKey
	var addr Address.Address
	_, err := os.Stat("keys" + "/")
	if os.IsNotExist(err) {
		fmt.Println("Directories do not exist, Generating new ones")
		err := os.Mkdir("keys"+"/", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat("keys" + "/" + folder_name); os.IsNotExist(err) {
		// fmt.Println("Error loading keys, Generating new ones")
		// fmt.Println("Use Python Script")

		fmt.Println("Directories do not exist, Generating new ones")
		os.Mkdir("keys"+"/"+folder_name, 0755)

		mnemonic, _ := HDWallet.GenerateMnemonic()
		os.WriteFile("keys"+"/"+folder_name+"/"+"mnemonic", []byte(mnemonic), 0644)
		hdWall, _ := HDWallet.NewHDWalletFromMnemonic(mnemonic, "")
		paymentPath := "m/1852'/1815'/0'/0/0"
		stakingPath := "m/1852'/1815'/0'/2/0"
		paymentKeyPath, _ := hdWall.DerivePath(paymentPath)
		verificationKey_bytes := paymentKeyPath.XPrivKey.PublicKey()
		signingKey_bytes := paymentKeyPath.XPrivKey.Bytes()
		stakingKeyPath, _ := hdWall.DerivePath(stakingPath)
		stakeVerificationKey_bytes := stakingKeyPath.XPrivKey.PublicKey()
		stakeSigningKey_bytes := stakingKeyPath.XPrivKey.Bytes()
		signingKey = Key.SigningKey{signingKey_bytes}
		verificationKey = Key.VerificationKey{verificationKey_bytes}
		stakeSigningKey = Key.StakeSigningKey{stakeSigningKey_bytes}
		stakeVerificationKey = Key.StakeVerificationKey{stakeVerificationKey_bytes}
		stakeVerKey := Key.VerificationKey{Payload: stakeVerificationKey_bytes}
		skh, _ := stakeVerKey.Hash()
		vkh, _ := verificationKey.Hash()

		os.WriteFile("keys"+"/"+folder_name+"/"+"public.key", verificationKey.Payload, 0644)
		os.WriteFile("keys"+"/"+folder_name+"/"+"private.key", signingKey.Payload, 0644)
		os.WriteFile("keys"+"/"+folder_name+"/"+"stake.public.key", stakeSigningKey.Payload, 0644)
		os.WriteFile("keys"+"/"+folder_name+"/"+"stake.private.key", stakeVerificationKey.Payload, 0644)
		addr = Address.Address{StakingPart: skh[:], PaymentPart: vkh[:], Network: 1, AddressType: Address.KEY_KEY, HeaderByte: 0b00000001, Hrp: "addr"}
		os.WriteFile("keys"+"/"+folder_name+"/"+"address.txt", []byte(addr.String()), 0644)
	} else {
		signingKey_bytes, err := os.ReadFile("keys" + "/" + folder_name + "/" + "private.key")
		if err != nil {
			log.Fatal(err)
		}
		verificationKey_bytes, err := os.ReadFile("keys" + "/" + folder_name + "/" + "public.key")
		if err != nil {
			log.Fatal(err)
		}
		stakeSigningKey_bytes, err := os.ReadFile("keys" + "/" + folder_name + "/" + "stake.private.key")
		if err != nil {
			log.Fatal(err)
		}
		stakeVerificationKey_bytes, err := os.ReadFile("keys" + "/" + folder_name + "/" + "stake.public.key")
		if err != nil {
			log.Fatal(err)
		}
		seed_bytes, err := os.ReadFile("keys" + "/" + folder_name + "/" + "mnemonic")
		if err != nil {
			signingKey = Key.SigningKey{Payload: signingKey_bytes}
			verificationKey = Key.VerificationKey{Payload: verificationKey_bytes}
			stakeSigningKey = Key.StakeSigningKey{Payload: stakeSigningKey_bytes}
			stakeVerificationKey = Key.StakeVerificationKey{Payload: stakeVerificationKey_bytes}
			stakeVerKey := Key.VerificationKey{Payload: stakeVerificationKey_bytes}
			skh, _ := stakeVerKey.Hash()
			vkh, _ := verificationKey.Hash()

			addr = Address.Address{StakingPart: skh[:], PaymentPart: vkh[:], Network: 1, AddressType: Address.KEY_KEY, HeaderByte: 0b00000001, Hrp: "addr"}
			return signingKey, verificationKey, stakeVerificationKey, stakeSigningKey, addr

		}
		paymentPath := "m/1852'/1815'/0'/0/0"
		stakingPath := "m/1852'/1815'/0'/2/0"
		hdWall, _ := HDWallet.NewHDWalletFromMnemonic(string(seed_bytes), "")
		paymentKeyPath, _ := hdWall.DerivePath(paymentPath)
		verificationKey_bytes = paymentKeyPath.XPrivKey.PublicKey()
		signingKey_bytes = paymentKeyPath.XPrivKey.Bytes()
		stakingKeyPath, _ := hdWall.DerivePath(stakingPath)
		stakeVerificationKey_bytes = stakingKeyPath.XPrivKey.PublicKey()
		stakeSigningKey_bytes = stakingKeyPath.XPrivKey.Bytes()
		//stake := stakingKeyPath.RootXprivKey.Bytes()
		signingKey = Key.SigningKey{Payload: signingKey_bytes}
		verificationKey = Key.VerificationKey{Payload: verificationKey_bytes}
		stakeSigningKey = Key.StakeSigningKey{Payload: stakeSigningKey_bytes}
		stakeVerificationKey = Key.StakeVerificationKey{Payload: stakeVerificationKey_bytes}
		stakeVerKey := Key.VerificationKey{Payload: stakeVerificationKey_bytes}
		skh, _ := stakeVerKey.Hash()
		vkh, _ := verificationKey.Hash()

		addr = Address.Address{StakingPart: skh[:], PaymentPart: vkh[:], Network: 1, AddressType: Address.KEY_KEY, HeaderByte: 0b00000001, Hrp: "addr"}

	}

	return signingKey, verificationKey, stakeVerificationKey, stakeSigningKey, addr
}
