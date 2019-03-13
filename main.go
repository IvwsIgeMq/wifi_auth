package main

import (
	"flag"
	"github.com/Nrehearsal/wifi_auth/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	port := flag.String("port", "8081", "http server port")
	sslOn := flag.String("ssl-on", "no", "is auth server ssl on")
	sslCert := flag.String("ssl-cert", "ssl.crt", "ssl certificate file")
	sslKey := flag.String("ssl-key", "ssl.key", "private key file for ssl certificate")

	flag.Parse()

	router := gin.Default()
	router.StaticFS("/static", http.Dir("static"))
	router.LoadHTMLGlob("static/*.html")

	router.GET("/ping", handler.Ping)

	router.GET("/login", handler.Login)

	router.POST("/logincheck", handler.LoginCheck)

	router.GET("/portal", handler.Portal)

	router.GET("/auth", handler.Auth)

	router.GET("/msg", handler.Msg)

	if *sslOn == "yes" {
		router.RunTLS(":"+*port, *sslCert, *sslKey)
	} else {
		router.Run(":" + *port)
	}
	return
}