package cliente

import (
	"criptografia/config/database"
	"criptografia/infrastructre/persistence/cadastros/cliente"
	"criptografia/infrastructre/persistence/cadastros/cliente/postgres"
	"criptografia/oops"
	"criptografia/utils"
)

type repositorio struct {
	pg *postgres.PGCliente
}

func novoRepo(db *database.DBTransacao) *repositorio {
	return &repositorio{
		pg: &postgres.PGCliente{DB: db},
	}
}

func (r *repositorio) Adicionar(req *cliente.Cliente) error {
	return r.pg.Adicionar(req)
}

func (r *repositorio) EncryptCliente(req *cliente.Cliente) error {
	return r.pg.EncryptCliente(req)
}

func (r *repositorio) ConverterParaCliente(dados interface{}) (*cliente.Cliente, error) {
	res := &cliente.Cliente{}
	if err := utils.ConvertStruct(dados, res); err != nil {
		return res, oops.Err(err)
	}
	return res, nil
}
