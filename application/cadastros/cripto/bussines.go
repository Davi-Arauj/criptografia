package cripto

import (
	"criptografia/application/cadastros/cliente"
	"criptografia/oops"
	"criptografia/utils"
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

func Encrypt(cliente cliente.Req, rng io.Reader) (res cliente.Req, err error) {

	msgErrPadrao := "Erro ao encriptar os dados"
	label := []byte("davimoreiraaraujo")

	cipherCard = []byte(*cliente.CreditCard)
	cipherCard, err = rsa.EncryptOAEP(sha256.New(), rng, Kp.pub, cipherCard, label)
	res.CreditCard = utils.PonteiroString(string(cipherCard))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption:%s\n", err)
		return res, oops.Wrap(err, msgErrPadrao)
	}

	cipherDocument = []byte(*cliente.Document)
	cipherDocument, err = rsa.EncryptOAEP(sha256.New(), rng, Kp.pub, cipherDocument, label)
	res.Document = utils.PonteiroString(string(cipherDocument))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption:%s\n", err)
		return res, oops.Wrap(err, msgErrPadrao)

	} // Como a criptografia é uma função aleatória, o texto cifrado será diferente a cada vez.

	fmt.Printf("CipherCard:-->%x\n", cipherCard)
	fmt.Printf("\nCipherDocument:-->%x\n", cipherDocument)

	return res, nil

}

func Decrypt(rng io.Reader) {
	label := []byte("davimoreiraaraujo")
	plainCard, err := rsa.DecryptOAEP(sha256.New(), rng, Kp.priv, cipherCard, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return
	}
	fmt.Printf("\nPlaincard: %s\n", string(plainCard))

	plainDocument, err := rsa.DecryptOAEP(sha256.New(), rng, Kp.priv, cipherDocument, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return
	}
	fmt.Printf("\nPlainDocument: %s\n", string(plainDocument))
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
