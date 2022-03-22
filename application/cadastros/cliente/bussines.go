package cliente

import (
	"context"
	"criptografia/config/database"
	"criptografia/domain/cadastros/cliente"
	"criptografia/oops"
	"criptografia/utils"
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

	if err = repoCliente.EncryptCliente(dados); err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	if err = repoCliente.Adicionar(dados); err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	id = dados.ID

	return

}

//Buscar contém a lógica de negócio para buscar um cliente
func Buscar(ctx context.Context, id string) (res *Res, err error) {

	msgErrPadrao := "Erro ao buscar um cliente"

	res = new(Res)

	tx, err := database.NovaTransacao(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repoCliente := cliente.ObterRepo(tx)

	req, err := repoCliente.ConverterParaCliente(res)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}
	req.ID = &id

	if err = repoCliente.Buscar(req); err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}
	if err = repoCliente.DecryptCliente(req); err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}

	if err = utils.ConvertStruct(req, res); err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}

	return

}
