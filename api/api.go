package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Arch-4ng3l/VisualCoding/www"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listeningAddr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr,
	}
}

func (s *APIServer) Init() {
	router := mux.NewRouter()

	router.HandleFunc("/api/node", nil)

	www.AddFrontend(router)
	fmt.Println("Listening on Port " + s.listeningAddr)
	http.ListenAndServe(s.listeningAddr, router)
}

func createFunc(f func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusForbidden, err)
		}
	}

}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)

}
