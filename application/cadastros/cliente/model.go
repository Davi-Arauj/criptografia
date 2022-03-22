package cliente

import "time"

// Req modela uma requisição para a criação ou atualização de um cliente
type Req struct {
	Document   *string `json:"documento,omitempty" binding:"required,stringField,gte=1" minLength:"1" codinome:"documento"`
	CreditCard *string `json:"cartao,omitempty" binding:"required,stringField,gte=1" minLength:"1" codinome:"cartao"`
	Value      *string `json:"valor,omitempty" binding:"required,stringField,gte=1" minLength:"1" codinome:"valor"`
}

// Res modela uma resposta para listagem e busca de clientes
type Res struct {
	ID              *string    `json:"id,omitempty" codinome:"id"`
	DataCriacao     *time.Time `json:"data_criacao,omitempty" codinome:"data_criacao"`
	DataAtualizacao *time.Time `json:"data_atualizacao,omitempty" codinome:"data_atualizacao"`
	Document        *string    `json:"documento,omitempty" binding:"required,stringField,gte=1" minLength:"1" codinome:"documento"`
	CreditCard      *string    `json:"cartao,omitempty" binding:"required,stringField,gte=1" minLength:"1" codinome:"cartao"`
	Value           *string    `json:"valor,omitempty" binding:"required,stringField,gte=1" minLength:"1" codinome:"valor"`
}
