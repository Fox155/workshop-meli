package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"workshop-meli/models"
	"workshop-meli/services"
	"workshop-meli/tools"

	"github.com/gin-gonic/gin"
)

//Recibe un contexto de gin
func AltaUsuario(c *gin.Context) {
	usuario := models.User{}

	if err := c.BindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if !usuario.IsValid() {
		c.JSON(http.StatusBadRequest, "Invalid purchase params")
		return
	}
	newUsuario, err := services.AltaUsuario(usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newUsuario)
}

func MostrarUsuarios(c *gin.Context) {
	c.JSON(http.StatusOK, services.MostrarTodosUsuarios())
}

func MostrarUsuario(c *gin.Context) {
	id := c.Param("id")
	numid, errorr := strconv.Atoi(id)

	if isValid := tools.ValidateString(id); !isValid && errorr != nil {
		c.JSON(http.StatusBadRequest, "Parametros Invalidos")
		return
	}

	if usuario, err := services.MostrarUsuarioPorID(numid); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	} else {
		c.JSON(http.StatusOK, usuario)
		return
	}
}

func ActualizarUsuario(c *gin.Context) {
	id := c.Param("id")
	usuario := models.User{}
	numid, errorr := strconv.Atoi(id)

	if err := c.BindJSON(&usuario); err != nil && errorr != nil {
		c.JSON(http.StatusBadRequest, "Tipo o formato de datos invalidos")
		return
	}
	nuevoUsuario, err := services.ActualizarUsuario(numid, usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, nuevoUsuario)
}

func EliminarUsuario(c *gin.Context) {
	id := c.Param("id")
	numid, errorr := strconv.Atoi(id)

	if isValid := tools.ValidateString(id); !isValid && errorr != nil {
		c.JSON(http.StatusBadRequest, "Parametros Invalidos")
		return
	}
	c.JSON(http.StatusAccepted, fmt.Sprint(services.EliminarUsuario(numid)))
	return
}
