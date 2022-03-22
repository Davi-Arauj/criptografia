package cadastros

import (
	"criptografia/interface/cadastros/cliente"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	//Grupamento das rotas de clientes em cadastros
	cliente.Router(r.Group("clientes"))
	cliente.RouterID(r.Group("cliente"))

}
