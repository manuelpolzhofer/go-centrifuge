package documents

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInt256(t *testing.T) {
	tests := []struct {
		name  string
		val   string
		isErr bool
	}{
		{
			"zero",
			"0",
			false,
		},
		{
			"int256 max",
			"57896044618658097711785492504343953926634992332820282019728792003956564819967",
			false,
		},
		{
			"int256 min",
			"-57896044618658097711785492504343953926634992332820282019728792003956564819967",
			false,
		},
		{
			"int256 positive intermediate",
			"323213",
			false,
		},
		{
			"int256 negative intermediate",
			"-10",
			false,
		},
		{
			"allow bit length 256, but fails because max",
			"115792089237316195423570985008687907853269984665640564039457584007913129639935",
			true,
		},
		{
			"not allow bit length more than 256 = 257",
			"115792089237316195423570985008687907853269984665640564039457584007913129639936",
			true,
		},
		{
			"more than allowed int256 (+1)",
			"57896044618658097711785492504343953926634992332820282019728792003956564819968",
			true,
		},
		{
			"less than allowed int256 (-1)",
			"-57896044618658097711785492504343953926634992332820282019728792003956564819968",
			true,
		},
		{
			"negative decimals aren't accepted",
			"-57896044618658097711785492.8990",
			true,
		},
		{
			"positive decimals aren't accepted",
			"4461865809771.89905453435435",
			true,
		},
		{
			"reject arbitrary strings",
			"ewrkwebj1232312.323",
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			i, err := NewInt256(test.val)
			if test.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, i)
				exp := new(big.Int)
				exp.SetString(test.val, 10)
				assert.True(t, i.v.Cmp(exp) == 0)
			}
		})
	}
}

func TestInt256_Bytes(t *testing.T) {
	tests := []struct {
		name string
		val  string
	}{
		{
			"zero",
			"0",
		},
		{
			"int256 max",
			"57896044618658097711785492504343953926634992332820282019728792003956564819967",
		},
		{
			"int256 min",
			"-57896044618658097711785492504343953926634992332820282019728792003956564819967",
		},
		{
			"positive intermediate",
			"5789604461865809771178549250434",
		},
		{
			"negative intermediate",
			"-5789604461865809771178549250434",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expectBig := intForString(test.val)
			intVal, err := NewInt256(test.val)
			assert.NoError(t, err)
			b := intVal.Bytes()
			intValGot, err := FromBytes(b[:])
			assert.NoError(t, err)
			assert.True(t, expectBig.Cmp(&intValGot.v) == 0)
			assert.True(t, intValGot.Equals(intVal))
		})
	}
}

func TestFromBytes(t *testing.T) {
	tests := []struct {
		name  string
		val   []byte
		isErr bool
	}{
		{
			"zero",
			intBytesForString(t, "0"),
			false,
		},
		{
			"int256 max",
			intBytesForString(t, "57896044618658097711785492504343953926634992332820282019728792003956564819967"),
			false,
		},
		{
			"int256 min",
			intBytesForString(t, "-57896044618658097711785492504343953926634992332820282019728792003956564819967"),
			false,
		},
		{
			"reject arbitrary byte length",
			[]byte{1, 1},
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			i, err := FromBytes(test.val)
			if test.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, i)
			}
		})
	}
}

func intBytesForString(t *testing.T, s string) []byte {
	n, err := NewInt256(s)
	assert.NoError(t, err)
	b := n.Bytes()
	return b[:]
}

func intForString(s string) *big.Int {
	n := new(big.Int)
	n.SetString(s, 10)
	return n
}
