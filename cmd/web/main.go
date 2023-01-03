package main

import (
	"fmt"
	"net/http"

	"github.com/datphat/pkg/handlers"
)

const TypePort = ":8778"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Serve run port %s", TypePort))
	_ = http.ListenAndServe(TypePort, nil)
}
