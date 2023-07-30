package ethereum

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

// EthereumWallet represents an Ethereum wallet containing the private key, public key, and hashed address.
type NewEthereumDisposableWallet struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Address    string `json:"address"`
}

// GenerateNewWallet generates a new Ethereum wallet and returns its details as an EthereumWallet instance.
func GenerateNewWallet() (*NewEthereumDisposableWallet, error) {
	// Generate a private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	// Convert private key to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)

	// Convert private key to a hexadecimal string
	privateKeyEncode := hexutil.Encode(privateKeyBytes)[2:]

	// Generate the public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// Convert the public key to bytes
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	// Convert the public key to a hexadecimal string
	publicKeyEncode := hexutil.Encode(publicKeyBytes[4:])

	// Calculate the hashed address using keccak256
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	hashedAddr := hexutil.Encode(hash.Sum(nil)[12:])

	// Create the EthereumWallet instance
	wallet := &NewEthereumDisposableWallet{
		PrivateKey: privateKeyEncode,
		PublicKey:  publicKeyEncode,
		Address:    hashedAddr,
	}

	return wallet, nil
}
