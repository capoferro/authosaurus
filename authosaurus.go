package main

import (
	"log"
	"strconv"
	"net/http"

	"github.com/capoferro/authosaurus/resources"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)

func main() {
	port := 8080
	domain := "localhost"
	wsUrl := "http://" + domain + ":" + strconv.Itoa(port)
	
	wsContainer := restful.NewContainer()
	u := resources.UserResource{map[string]resources.User{}}
	u.Register(wsContainer)

	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(),
		WebServicesUrl: wsUrl,
		ApiPath:        "/api-docs",
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "./vendor/swagger-ui/dist",
	}
	swagger.RegisterSwaggerService(config, wsContainer)

	log.Printf("Listening: " + wsUrl)
	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
