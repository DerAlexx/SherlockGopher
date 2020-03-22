package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	webserver "github.com/ob-algdatii-20ss/leistungsnachweis-dievierausrufezeichen/webserver/webserver"
)

const (
	servicename string = "CrawlWebServer"
	address     string = "localhost:8081"
)

/*
getServiceName will return a the service name of this service.
*/
func getServiceName() string {
	return servicename
}

/*
getAddress will return the address of the service.
*/
func getAddress() string {
	return address
}

func main() {
	service := web.NewService(web.Name(getServiceName()))

	err := service.Init()

	if err != nil {
		log.Fatal(err)
	}

	webServerService := webserver.New()

	router := gin.Default()

	router.GET("/helloping", webServerService.Helloping)
	router.POST("search", webServerService.RecieveURL)

	router.Run(getAddress())

	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
