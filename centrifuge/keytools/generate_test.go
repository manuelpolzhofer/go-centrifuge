package keytools

import (
	"os"
	"testing"

	"github.com/CentrifugeInc/go-centrifuge/centrifuge/keytools/io"
	"github.com/magiconair/properties/assert"
)

const (
	PrivateKeySECP256K1Len = 32
	PublicKeySECP256K1Len  = 65
	PublicKeyED25519Len    = 32
	PrivateKeyED25519Len   = 64
)

func GenerateKeyFilesForTest(t *testing.T, curve string) (publicKey, privateKey []byte) {

	publicFileName := "publicKeyFile"
	privateFileName := "privateKeyFile"

	GenerateSigningKeyPair(publicFileName, privateFileName, curve)

	_, err := os.Stat(publicFileName)

	assert.Equal(t, err != nil, false, "public key file not generated")

	_, err = os.Stat(privateFileName)

	assert.Equal(t, err != nil, false, "private key file not generated")

	publicKey, err = io.ReadKeyFromPemFile(publicFileName, PublicKey)

	if err != nil {
		log.Fatal(err)
	}

	privateKey, err = io.ReadKeyFromPemFile(privateFileName, PrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	os.Remove(publicFileName)
	os.Remove(privateFileName)

	return

}

func TestGenerateSigningKeyPairSECP256K1(t *testing.T) {

	curve := CurveSecp256K1
	publicKey, privateKey := GenerateKeyFilesForTest(t, curve)

	assert.Equal(t, len(publicKey), PublicKeySECP256K1Len, "public key length not correct")
	assert.Equal(t, len(privateKey), PrivateKeySECP256K1Len, "private key length not correct")

}

func TestGenerateSigningKeyPairED25519(t *testing.T) {

	curve := CurveEd25519
	publicKey, privateKey := GenerateKeyFilesForTest(t, curve)

	assert.Equal(t, len(publicKey), PublicKeyED25519Len, "public key length not correct")
	assert.Equal(t, len(privateKey), PrivateKeyED25519Len, "private key length not correct")
}