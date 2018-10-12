package hello

import (
	"fmt"
	"net/http"
)

// HelloWorld ...
func HelloWorld(w http.ResponseWriter, r *http.ResponseWriter) {
	fmt.Fprint(w, "Hello World!")
}
