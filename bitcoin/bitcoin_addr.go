package bitcoin

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"

	// "math/big"
	"crypto/rand"
	"log"

	"github.com/btcsuite/btcutil/base58"
	rp "golang.org/x/crypto/ripemd160"
)

func init() {
	log.SetPrefix("Bitcoin:")

}

const NETWORK_ID = 0x00

type BitcoinDisposableWallet struct {
	PrivateKey *ecdsa.PrivateKey `json:"priv_key"`
	PublicKey  *ecdsa.PublicKey  `json:"pub_key"`
	Address    string            `json:"bitcoin_addr"`
}

func NewBitcoinDisposableWallet() (*BitcoinDisposableWallet, error) {

	// Generate a private key
	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		// replace with error returned in response from server
		//log.Fatal("error generating new private key: %v ", err.Error())
		return nil, err
	}

	newWallet := new(BitcoinDisposableWallet)
	newWallet.PrivateKey = pk
	newWallet.PublicKey = &newWallet.PrivateKey.PublicKey

	// Generate wallet address
	firstHash := sha256.New()
	// First hashing of public key with sha256
	firstHash.Write(newWallet.PublicKey.X.Bytes())
	firstHash.Write(newWallet.PublicKey.Y.Bytes())
	firstDigest := firstHash.Sum(nil)
	// Second hashing of digest from first hashing with RIPEMD
	secondHash := rp.New()
	secondHash.Write(firstDigest)
	secondDigest := secondHash.Sum(nil)
	// Concatenating the NETWORK_ID to digest
	versionConcat := make([]byte, 21)
	versionConcat[0] = NETWORK_ID
	copy(versionConcat[1:], secondDigest[:])
	// Third hashing of values from concantenation of network id and digest from RIPEMD
	thirdHash := sha256.New()
	thirdHash.Write(versionConcat)
	thirdDigest := thirdHash.Sum(nil)
	// Fourth hashing of values from previous hash digest
	fourthHash := sha256.New()
	fourthHash.Write(thirdDigest)
	fourthDigest := fourthHash.Sum(nil)
	checkSum := fourthDigest[:4]
	checkSumConcat := make([]byte, 25)
	copy(checkSumConcat[:21], versionConcat)
	copy(checkSumConcat[21:], checkSum)
	address := base58.Encode(checkSumConcat)
	newWallet.Address = address
	return newWallet, err

}
