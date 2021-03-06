// +build unit

package crypto

import (
	"os"
	"testing"

	"github.com/centrifuge/go-centrifuge/crypto/secp256k1"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var (
	key1      = []byte{0xbc, 0xa4, 0xe3, 0x6a, 0x57, 0x9a, 0xc1, 0xfd, 0xc1, 0xf6, 0x5b, 0xbf, 0xbc, 0x3f, 0x72, 0x16, 0xe9, 0x93, 0x74, 0x83, 0xbf, 0x9a, 0x5d, 0xd9, 0x20, 0xf4, 0x8a, 0xf7, 0xc9, 0xe3, 0x9f, 0xcd}
	key1Pub   = []byte{0x04, 0x8f, 0x06, 0x87, 0x01, 0xfa, 0xbb, 0x2f, 0xb6, 0x67, 0xa4, 0x96, 0xfc, 0xb2, 0x64, 0x41, 0x1c, 0xd8, 0x58, 0x5c, 0x3b, 0x5b, 0xd3, 0xe1, 0xab, 0xfc, 0x47, 0x64, 0x2f, 0x98, 0xd8, 0xbb, 0x8f, 0x28, 0x77, 0xbb, 0xa1, 0x1b, 0xfc, 0xa3, 0x5b, 0xf8, 0xf1, 0xf2, 0x10, 0x52, 0x39, 0x45, 0x7c, 0x24, 0x41, 0xaf, 0xb9, 0x2a, 0xf1, 0x9f, 0xe4, 0x99, 0xf4, 0xb4, 0x91, 0xcf, 0x3a, 0xb5, 0x96}
	signature = []byte{0x13, 0xa6, 0xb7, 0xff, 0x47, 0x10, 0x9f, 0x2b, 0x51, 0x3, 0x7c, 0xeb, 0x42, 0xfa, 0xc3, 0x52, 0x85, 0x2b, 0xa8, 0xea, 0xe, 0x34, 0x4f, 0xc0, 0xf7, 0x4a, 0xe9, 0x20, 0x57, 0xbd, 0x0, 0xb9, 0x7e, 0x87, 0x53, 0xa0, 0x34, 0xa8, 0xe1, 0xa, 0x80, 0xda, 0xb8, 0x82, 0x4, 0x11, 0xd3, 0x19, 0x1c, 0xbc, 0x8, 0x16, 0xd6, 0xfa, 0xae, 0x15, 0xcd, 0xb3, 0xc1, 0x97, 0x7d, 0x41, 0xd2, 0x52, 0x1}
)

func TestSign(t *testing.T) {
	sig, err := SignMessage(key1, key1Pub, CurveSecp256K1)
	assert.NoError(t, err)
	assert.NotEmpty(t, sig)
	assert.Len(t, sig, 65)
	assert.Equal(t, signature, sig)
}

func TestValidateSignature_invalid_sig(t *testing.T) {
	pubKey := key1Pub
	message := key1Pub
	signature := utils.RandomSlice(32)
	pk32 := utils.AddressTo32Bytes(common.HexToAddress(secp256k1.GetAddress(pubKey)))
	valid := VerifyMessage(pk32[:], message, signature, CurveSecp256K1)
	assert.False(t, valid, "must be false")
}

func TestSignMessageUnsupportedType(t *testing.T) {
	sig, err := SignMessage(key1, key1Pub, "rsa")
	assert.Error(t, err)
	assert.Empty(t, sig)
}

func TestSignMessageSecp256k1(t *testing.T) {

	publicKeyFile := "publicKey"
	privateKeyFile := "privateKey"
	testMsg := []byte("test")

	GenerateSigningKeyPair(publicKeyFile, privateKeyFile, CurveSecp256K1)
	privateKey, err := utils.ReadKeyFromPemFile(privateKeyFile, utils.PrivateKey)
	assert.Nil(t, err)
	publicKey, err := utils.ReadKeyFromPemFile(publicKeyFile, utils.PublicKey)
	assert.Nil(t, err)
	signature, err := SignMessage(privateKey, testMsg, CurveSecp256K1)
	assert.Nil(t, err)
	pk32 := utils.AddressTo32Bytes(common.HexToAddress(secp256k1.GetAddress(publicKey)))
	correct := VerifyMessage(pk32[:], testMsg, signature, CurveSecp256K1)

	os.Remove(publicKeyFile)
	os.Remove(privateKeyFile)

	assert.True(t, correct, "signature or verification didn't work correctly")
}

func TestSignMessageEd25519(t *testing.T) {

	publicKeyFile := "publicKey"
	privateKeyFile := "privateKey"
	testMsg := []byte("test")

	GenerateSigningKeyPair(publicKeyFile, privateKeyFile, CurveEd25519)
	privateKey, err := utils.ReadKeyFromPemFile(privateKeyFile, utils.PrivateKey)
	assert.Nil(t, err)
	publicKey, err := utils.ReadKeyFromPemFile(publicKeyFile, utils.PublicKey)
	assert.Nil(t, err)
	signature, err := SignMessage(privateKey, testMsg, CurveEd25519)
	assert.Nil(t, err)
	correct := VerifyMessage(publicKey, testMsg, signature, CurveEd25519)

	os.Remove(publicKeyFile)
	os.Remove(privateKeyFile)

	assert.True(t, correct, "signature or verification didn't work correctly")
}
