package http

import (
	"encoding/json"
	"fmt"
	"jwt-auth-with-gorm/application"
	"net/http"
	"os"
	"strconv"
	"time"
)

func CreateHttpHandler() http.Handler {
	sm := http.NewServeMux()
	sm.HandleFunc("/getbook", GetBookHandler)
	return sm
}

func GetBookHandler(w http.ResponseWriter, req *http.Request) {
	param := req.FormValue("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		json.NewEncoder(w).Encode("param type must be integer")
	}

	books, err := application.Get(id)
	if err != nil {
		json.NewEncoder(w).Encode("book does not find")
	} else {
		json.NewEncoder(w).Encode(books)
	}
}

func CreateHttpServer() http.Server {
	addr := os.Getenv("ADDR")
	server := http.Server{
		Addr:              addr,
		Handler:           CreateHttpHandler(),
		MaxHeaderBytes:    1 * 1024 * 1024,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}
	fmt.Printf("http server listening on %s ", addr)
	return server
}
