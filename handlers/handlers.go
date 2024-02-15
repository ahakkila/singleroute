package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPerson(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "OK - GetPerson: "+id)
}

func GetPersonByID(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "OK - GetPersonByID: "+id)
}

func CreatePerson(c *gin.Context) {
	c.String(http.StatusOK, "OK - CreatePerson")
}

func GetAccount(c *gin.Context) {

	c.String(http.StatusOK, "OK - GetAccount")
}

func GetAccountByID(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "OK - GetAccountByID: "+id)
}

func CreateAccount(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "OK - CreateAccount: "+id)
}
