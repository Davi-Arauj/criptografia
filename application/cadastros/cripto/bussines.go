package cripto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

type Keypair struct {
	priv *rsa.PrivateKey
	pub  *rsa.PublicKey
}

const (
	rsaKeySize = 2048
)

var Kp Keypair
var ciphertext, signedMessage []byte
var hashed [32]byte
var cipherCard []byte
var cipherDocument []byte

func GenerateKeypair() (err error) {

	// var err error
	Kp.priv, err = rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		return err
	}

	Kp.pub = &Kp.priv.PublicKey
	return nil
}



func Sign(rng io.Reader) {
	var err error
	message := []byte("This is the plaintext to be signed")
	hashed = sha256.Sum256(message)
	signedMessage, err = rsa.SignPKCS1v15(rng, Kp.priv, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return
	}
	fmt.Printf("\nSigned Message: %x\n", signedMessage)
}

func Verify() {
	err := rsa.VerifyPKCS1v15(Kp.pub, crypto.SHA256, hashed[:], signedMessage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return
	}
	fmt.Printf("Message verified!\n")
}
