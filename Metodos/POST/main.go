package main 

import "github.com/gin-gonic/gin"

type Login struct {
	Email string `json:"email" blinding:"required"`
	Password string `json:"password" blinding:"required"`

}


func main () {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var login Login
		// Una peticion va acompa;ada de 3 componentes 
		// 1. URL  -> http://localhost:8080/login
		// 2. Headers  -> Content-Type: application/json
		// 3. Body  -> JSON
		c.BindJSON(&login) // BindJSON es un metodo que se encarga de leer el body de la peticion y asignarlo a la variable login

		c.JSON(200, gin.H{
			"status": login.Email})
	})

	r.Run(":8080")
}