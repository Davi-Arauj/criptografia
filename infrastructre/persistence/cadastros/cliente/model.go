package cliente

import "time"

//Cliente estrutura para definiçãe de modelo de cliente para uso na camada de dados
type Cliente struct {
	ID              *string    `sql:"id" codinome:"id"`
	DataCriacao     *time.Time `sql:"data_criacao::TIMESTAMPTZ" codinome:"data_criacao"`
	DataAtualizacao *time.Time `sql:"data_atualizacao::TIMESTAMPTZ" codinome:"data_atualizacao"`
	Document        *string    `sql:"userdocument" codinome:"userdocument"`
	CreditCard      *string    `sql:"creditcard" codinome:"creditcard"`
	Value           *string    `sql:"value" codinome:"value"`
}

