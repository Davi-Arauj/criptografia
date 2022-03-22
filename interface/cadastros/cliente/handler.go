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

func buscar(c *gin.Context) {
	id := c.Param("cliente_id")

	res, err := cliente.Buscar(c, id)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(200, res)
}
