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

//Adicionar adiciona um registro no banco
func (r *repositorio) Adicionar(req *cliente.Cliente) error {
	return r.pg.Adicionar(req)
}

//Buscar busca um registro no banco
func (r *repositorio) Buscar(req *cliente.Cliente) error{
	return r.pg.Buscar(req)
}

//EncryptCliente realiza a criptografia dos dados necessários
func (r *repositorio) EncryptCliente(req *cliente.Cliente) error {
	return r.pg.EncryptCliente(req)
}

//DecryptCliente realiza a criptografia dos dados necessários
func (r *repositorio) DecryptCliente(req *cliente.Cliente) error {
	return r.pg.DecryptCliente(req)
}

//Converte uma struct comun em uma struct de um Cliente
func (r *repositorio) ConverterParaCliente(dados interface{}) (*cliente.Cliente, error) {
	res := &cliente.Cliente{}
	if err := utils.ConvertStruct(dados, res); err != nil {
		return res, oops.Err(err)
	}
	return res, nil
}
