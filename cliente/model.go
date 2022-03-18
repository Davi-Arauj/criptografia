package cliente

type Cliente struct {
	ID         string `json:"id"`
	Document   string `json:"documento"`
	CreditCard string `json:"cartao"`
	Value      string `json:"valor"`
}
