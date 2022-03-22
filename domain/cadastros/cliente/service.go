package cliente

import (
	"criptografia/config/database"
)

// Servico define a estrutura base para uso dos métodos do serviço
type Servico struct {
	repo ICliente
}

// ObterServico retorna um servico para acesso a funções de auxilio
// a lógica de negócio
func ObterServico(r ICliente) *Servico {
	return &Servico{repo: r}
}

// ObterRepo retorna um repositório para acesso à camada de dados
func ObterRepo(tx *database.DBTransacao) ICliente {
	return novoRepo(tx)
}
