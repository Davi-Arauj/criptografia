package cliente

import (
	"criptografia/infrastructre/persistence/cadastros/cliente"
)

// ICliente define uma interface para os metodos de acesso à camada de dados
type ICliente interface {
	Adicionar(req *cliente.Cliente) error
	ConverterParaCliente(data interface{}) (*cliente.Cliente, error)
	EncryptCliente(cliente *cliente.Cliente) error
	DecryptCliente(cliente *cliente.Cliente) error
	Buscar(cliente *cliente.Cliente) error
}
