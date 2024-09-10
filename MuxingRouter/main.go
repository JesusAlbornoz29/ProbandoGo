package main 

import ( 
	"github.com/gin-gonic/gin"
)

// Agrupado de endpoints
func main() {
	router := gin.Default()
	// LLamamos a GET y le pasamos una ruta,
	// en otras palabras, 
	// definimos un nuevo endpoint
	router.GET("/", Endpoint1)
	router.GET("/login", Endpoint2)
	router.GET("/login/perfil", Endpoint3)
	router.RUN(":8080")
}

// Sin agrupar
func main() {
	server := gin.Default()
	server.GET("/", bienvenida)
	/*
	Ahora agrupamos todos los endpoints
	referidos al perfil y
	dos endpoints sin agrupar
	*/
	gopher := server.Group("/perfil")
	{
		gopher.GET("/foto", fotoDePerfil)
		gopher.GET("/datos", datoPerfil)
		gopher.GET("/amigos", amigosPerfil)
	}
	server.GET("/about", HandlerAbout)
	server.Run(":8080")
}