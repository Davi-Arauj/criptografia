package main

import (
	"criptografia/cliente"
	"criptografia/db"
	"crypto/rand"
)

var Davi = cliente.Cliente{}

func main() {
	db.Connection()
	defer db.DB.Close()
	Davi.ID = "1"
	Davi.Document = "documento"
	Davi.CreditCard = "cartão de debito"

	rng := rand.Reader
	cliente.GenerateKeypair()
	cliente.Encrypt(Davi, rng)
	cliente.Decrypt(rng)
	cliente.Sign(rng)
	cliente.Verify()
}
