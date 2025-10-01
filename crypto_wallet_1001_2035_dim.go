// 代码生成时间: 2025-10-01 20:35:43
package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "log"
)

// CryptoWallet represents a cryptocurrency wallet
type CryptoWallet struct {
    privateKey *ecdsa.PrivateKey
}

// NewCryptoWallet generates a new cryptographic wallet
func NewCryptoWallet() (*CryptoWallet, error) {
    privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        return nil, fmt.Errorf("failed to generate private key: %w", err)
    }
    return &CryptoWallet{privateKey: privateKey}, nil
}

// GetPublicKey returns the public key of the wallet
func (w *CryptoWallet) GetPublicKey() (string, error) {
    publicKeyBytes := elliptic.Marshal(elliptic.P256(), w.privateKey.X, w.privateKey.Y)
    publicKeyHex := hex.EncodeToString(publicKeyBytes)
    return publicKeyHex, nil
}

// GetPrivateKey returns the private key of the wallet
func (w *CryptoWallet) GetPrivateKey() (string, error) {
    privateKeyBytes := w.privateKey.D.Bytes()
    privateKeyHex := hex.EncodeToString(privateKeyBytes)
    return privateKeyHex, nil
}

func main() {
    // Create a new cryptocurrency wallet
    wallet, err := NewCryptoWallet()
    if err != nil {
        log.Fatalf("failed to create wallet: %s", err)
    }
    fmt.Println("Wallet created successfully")

    // Get and print the public key
    publicKey, err := wallet.GetPublicKey()
    if err != nil {
        log.Fatalf("failed to get public key: %s", err)
    }
    fmt.Printf("Public Key: %s
", publicKey)

    // Get and print the private key
    privateKey, err := wallet.GetPrivateKey()
    if err != nil {
        log.Fatalf("failed to get private key: %s", err)
    }
    fmt.Printf("Private Key: %s
", privateKey)
}
