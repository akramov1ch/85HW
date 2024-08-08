package handlers

import (
	"fmt"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Rate Limited API!")
}
