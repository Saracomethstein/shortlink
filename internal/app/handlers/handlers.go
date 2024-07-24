package handlers

import (
	"fmt"
	"net/http"
)

func HandlerHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi!")
}
