package main 

import (
	"github.com/gin-gonic/gin"
)


func main() { 

/* Crea un router con Gin
 Esto nos permite ir agregando los distintos endpoints que tendra nuestra aplicacion
 Para ello debemos agregar, al router, distintos handlers.
*/
router := gin.Default()

// Catura la solicitud GET en la ruta /hola
router.GET("/hola", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"mensaje": "Hola",
	})
})

router.Run() // Corre el servidor en el puerto 8080
}


