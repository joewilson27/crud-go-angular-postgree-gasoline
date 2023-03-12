package main

import (
	"fmt"
	"go-postgres-crud/router"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Server dijalankan pada port 8080...")

	log.Fatal(http.ListenAndServe(":8080", &CORSRouterDecorator{r}))
}

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Accept-Language,"+
				" X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
