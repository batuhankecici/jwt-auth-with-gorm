package main

import "jwt-auth-with-gorm/transport/http"

func main() {
	server := http.CreateHttpServer()
	server.ListenAndServe()
}
