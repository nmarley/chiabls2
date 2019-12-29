package blschia_test

import (
	"bytes"
	"testing"

	bls "github.com/nmarley/chiabls2"
)

func TestExtendedPublicKey(t *testing.T) {
	xpub1 := bls.ExtendedPublicKeyFromBytes(xpubBytes)
	xpub1GotBytes := xpub1.Serialize()
	if !bytes.Equal(xpub1GotBytes, xpubBytes) {
		t.Errorf("got %v, expected %v", xpub1GotBytes, xpubBytes)
	}

	xpub3 := bls.ExtendedPublicKeyFromBytes(xpub1GotBytes)
	if !xpub1.Equal(xpub3) {
		t.Error("xpub1 should be equal to xpub3")
	}
	xpub2 := xpub1.PublicChild(1)
	if xpub1.Equal(xpub2) {
		t.Error("xpub1 should NOT be equal to xpub2")
	}

	expectedVersion := uint32(1)
	actualVersion := xpub2.GetVersion()
	if actualVersion != expectedVersion {
		t.Errorf("got version %v, expected %v", actualVersion, expectedVersion)
	}

	expectedDepth := uint8(1)
	actualDepth := xpub2.GetDepth()
	if actualDepth != expectedDepth {
		t.Errorf("got depth %v, expected %v", actualDepth, expectedDepth)
	}

	expectedParentFingerprint := uint32(0xa4700b27)
	actualParentFingerprint := xpub2.GetParentFingerprint()
	if actualParentFingerprint != expectedParentFingerprint {
		t.Errorf("got parent fingerprint %v, expected %v", actualParentFingerprint, expectedParentFingerprint)
	}

	expectedChildNumber := uint32(1)
	actualChildNumber := xpub2.GetChildNumber()
	if actualChildNumber != expectedChildNumber {
		t.Errorf("got child number %v, expected %v", actualChildNumber, expectedChildNumber)
	}

	expectedChainCode := []byte{
		0x21, 0x66, 0xf1, 0x87, 0xb5, 0x73, 0xb3, 0x54,
		0x4a, 0xc8, 0x56, 0x46, 0x14, 0x62, 0xec, 0xf9,
		0x88, 0x59, 0xff, 0x82, 0xcb, 0x68, 0x9c, 0x26,
		0x37, 0x79, 0xb8, 0x0d, 0xb8, 0x2b, 0xc8, 0x93,
	}
	actualChainCode := xpub2.GetChainCode().Serialize()
	if !bytes.Equal(expectedChainCode, actualChainCode) {
		t.Errorf("got %v, expected %v", expectedChainCode, actualChainCode)
	}

	expectedPubKey := []byte{
		0x8b, 0x07, 0x21, 0x54, 0x51, 0xa4, 0x74, 0x12,
		0x5d, 0x37, 0x5d, 0x27, 0x61, 0x1d, 0x10, 0xfb,
		0x77, 0x32, 0xc3, 0xba, 0x40, 0x9b, 0x20, 0xc0,
		0x3a, 0x5a, 0x09, 0x94, 0x77, 0xe3, 0xc4, 0x8e,
		0x8d, 0x4f, 0xae, 0xf4, 0x67, 0x74, 0x02, 0x6b,
		0xa3, 0x17, 0x6c, 0x65, 0x5e, 0x9d, 0x08, 0x6a,
	}
	actualPubKey := xpub2.GetPublicKey().Serialize()
	if !bytes.Equal(actualPubKey, expectedPubKey) {
		t.Errorf("got key bytes %v, expected %v", actualPubKey, expectedPubKey)
	}

	xpub1.Free()
	xpub2.Free()
	xpub3.Free()
}

var xpubBytes = []byte{
	0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0xd8, 0xb1, 0x25,
	0x55, 0xb4, 0xcc, 0x55, 0x78, 0x95, 0x1e, 0x4a,
	0x7c, 0x80, 0x03, 0x1e, 0x22, 0x01, 0x9c, 0xc0,
	0xdc, 0xe1, 0x68, 0xb3, 0xed, 0x88, 0x11, 0x53,
	0x11, 0xb8, 0xfe, 0xb1, 0xe3, 0x0a, 0xa5, 0x5d,
	0xb2, 0x14, 0xbc, 0x45, 0x6d, 0xe8, 0x3f, 0x84,
	0xca, 0xf1, 0x17, 0xd2, 0x5f, 0xb7, 0x6e, 0xaf,
	0xbc, 0xf2, 0x11, 0x59, 0x57, 0x1c, 0xdb, 0xc7,
	0x66, 0x27, 0xf6, 0x29, 0xb6, 0xdc, 0x93, 0x71,
	0x28, 0xc2, 0x59, 0xca, 0xe6, 0xeb, 0xaa, 0x18,
	0x0e, 0x45, 0xde, 0x95, 0x7f,
}