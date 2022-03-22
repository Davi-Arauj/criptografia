package cliente

import "time"

//Cliente estrutura para definiçãe de modelo de cliente para uso na camada de dados
type Cliente struct {
	ID              *string    `sql:"id" codinome:"id"`
	DataCriacao     *time.Time `sql:"data_criacao::TIMESTAMPTZ" codinome:"data_criacao"`
	DataAtualizacao *time.Time `sql:"data_atualizacao::TIMESTAMPTZ" codinome:"data_atualizacao"`
	Document        *string    `sql:"userdocument" codinome:"documento"`
	CreditCard      *string    `sql:"creditcard" codinome:"cartao"`
	Value           *string    `sql:"value" codinome:"valor"`
}
