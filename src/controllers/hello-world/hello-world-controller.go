package helloworld

import (
	"encoding/json"
	"net/http"
)

func HelloWord() {
	http.HandleFunc("/hello-world", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)

		json.NewEncoder(rw).Encode("{message:'Hello World'}")
	})
}
