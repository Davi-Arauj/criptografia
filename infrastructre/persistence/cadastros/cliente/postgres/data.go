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
	"os"

	"github.com/Masterminds/squirrel"
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

//Buscar busca um cliente no banco de dados do postgres
func (pg *PGCliente) Buscar(req *cliente.Cliente) (err error) {
	if err = pg.DB.Builder.
		Select(`id,data_criacao::TIMESTAMPTZ,
				   data_atualizacao::TIMESTAMPTZ,
				   userdocument,
				   creditcard,value`).
		From(`t_cliente`).
		Where(squirrel.Eq{
			"id": req.ID,
		}).
		Scan(&req.ID, &req.DataCriacao, &req.DataAtualizacao, &req.Document, &req.CreditCard, &req.Value); err != nil {
		return oops.Err(err)
	}
	return
}

func GenerateKeypair() (err error) {

	Kp.priv, err = rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		return
	}

	Kp.pub = &Kp.priv.PublicKey
	return
}

func (pg *PGCliente) EncryptCliente(cliente *cliente.Cliente) error {
	GenerateKeypair()

	var (
		err            error
		msgErrPadrao   = "Erro ao encriptar os dados"
		rng            = rand.Reader
		label          = []byte("davimoreiraaraujo")
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
		rng          = rand.Reader
		label        = []byte("davimoreiraaraujo")
	)

	cipherCard, err := hex.DecodeString(*cliente.CreditCard)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	cipherDocument, err := hex.DecodeString(*cliente.Document)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
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
