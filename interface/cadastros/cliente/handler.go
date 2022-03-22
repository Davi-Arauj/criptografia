package cliente

import (
	"criptografia/application/cadastros/cliente"
	"criptografia/oops"
	"net/http"

	"github.com/gin-gonic/gin"
)

func adicionar(c *gin.Context) {
	var req cliente.Req

	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	id, err := cliente.Adicionar(c, &req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	c.JSON(http.StatusCreated, id)
}

func listar(c *gin.Context) {

}

func buscar(c *gin.Context) {

}
