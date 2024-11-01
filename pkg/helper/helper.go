package helper

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

func StringToECDSA(hexKey string) (*ecdsa.PrivateKey, error) {
	// Decode the hex string to bytes
	keyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex string: %v", err)
	}

	// Convert the byte slice to an ECDSA private key
	privateKey, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
	}

	return privateKey, nil
}
