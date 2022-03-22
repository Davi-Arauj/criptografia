package cliente

import (
	"context"
	"criptografia/config/database"
	"criptografia/domain/cadastros/cliente"
	"criptografia/oops"
	"log"
)

func Adicionar(ctx context.Context, req *Req) (id *string, err error) {
	var (
		msgErrPadrao = "Erro ao cadastrar novo cliente"
	//	p utils.ParametrosRequisicao
	)
	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repoCliente := cliente.ObterRepo(tx)

	dados, err := repoCliente.ConverterParaCliente(req)
	if err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}
	dados.CreditCard = req.CreditCard
	dados.Document = req.Document
	dados.Value = req.Value
	log.Println(dados)
	if err = repoCliente.EncryptCliente(dados); err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}
	log.Println(dados)
	if err = repoCliente.Adicionar(dados); err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	id = dados.ID

	return

}
