package cliente

import (
	"criptografia/utils"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("", utils.AddRota("Lista clientes", "Lista clientes", listar)...)
	r.POST("", utils.AddRota("Cadastro de um novo cliente", "Cadastra um novo cliente", adicionar)...)

}

func RouterID(r *gin.RouterGroup){
	r.GET(":cliente_id",utils.AddRota("Busca um cliente","Busca um cliente",buscar)...)
}
