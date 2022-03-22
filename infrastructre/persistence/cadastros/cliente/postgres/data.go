package postgres

import (
	"criptografia/config/database"
	"criptografia/infrastructre/persistence/cadastros/cliente"
	"criptografia/oops"
	"criptografia/utils"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

type PGCliente struct {
	DB *database.DBTransacao
}
type Keypair struct {
	priv *rsa.PrivateKey
	pub  *rsa.PublicKey
}

var Kp Keypair

const (
	rsaKeySize = 700
)

//Adicionar adiciona um novo cliente ao banco de dados do postgres
func (pg *PGCliente) Adicionar(req *cliente.Cliente) (err error) {

	cols, vals, err := utils.FormatarInsertUpdate(req)
	if err != nil {
		return oops.Err(err)
	}

	if err = pg.DB.Builder.
		Insert("t_cliente").
		Columns(cols...).
		Values(vals...).
		Suffix(`RETURNING "id"`).
		Scan(&req.ID); err != nil {
		return oops.Err(err)
	}
	return
}

func GenerateKeypair() (err error) {

	Kp.priv, err = rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		return err
	}

	Kp.pub = &Kp.priv.PublicKey
	return nil
}

var rng = rand.Reader
var label = []byte("davimoreiraaraujo")

func (pg *PGCliente) EncryptCliente(cliente *cliente.Cliente) error {
	GenerateKeypair()
	var (
		err          error
		msgErrPadrao = "Erro ao encriptar os dados"

		cipherCard     = []byte(*cliente.CreditCard)
		cipherDocument = []byte(*cliente.Document)
	)

	cipherCard, err = rsa.EncryptOAEP(sha256.New(), rng, Kp.pub, cipherCard, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption:%s\n", err)
		return oops.Wrap(err, msgErrPadrao)
	}
	cliente.CreditCard = utils.PonteiroString(hex.EncodeToString(cipherCard))

	cipherDocument, err = rsa.EncryptOAEP(sha256.New(), rng, Kp.pub, cipherDocument, label)
	cliente.Document = utils.PonteiroString(hex.EncodeToString(cipherDocument))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption:%s\n", err)
		return oops.Wrap(err, msgErrPadrao)

	} // Como a criptografia é uma função aleatória, o texto cifrado será diferente a cada vez.

	return nil

}

func (pg *PGCliente) DecryptCliente(cliente *cliente.Cliente) error {

	var (
		msgErrPadrao = "Erro ao Desencriptar os dados"
	)

	cipherCard, err := (hex.DecodeString(*cliente.CreditCard))
	if err != nil {
		log.Println(err)
	}
	cipherDocument, err := hex.DecodeString(*cliente.Document)
	if err != nil {
		log.Println(err)
	}

	plainCard, err := rsa.DecryptOAEP(sha256.New(), rng, Kp.priv, cipherCard, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Card Error from decryption: %s\n", err)
		return oops.Wrap(err, msgErrPadrao)
	}
	cliente.CreditCard = utils.PonteiroString(string(plainCard))

	plainDocument, err := rsa.DecryptOAEP(sha256.New(), rng, Kp.priv, cipherDocument, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Document Error from decryption: %s\n", err)
		return oops.Wrap(err, msgErrPadrao)
	}

	cliente.Document = utils.PonteiroString(string(plainDocument))

	return nil
}
