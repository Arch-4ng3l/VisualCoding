package www

import (
	"net/http"

	"github.com/Arch-4ng3l/VisualCoding/types"
	"github.com/gorilla/mux"
)

type Data struct {
	Content []types.Node
}

func AddFrontend(router *mux.Router) {

	fs := http.FileServer(http.Dir("./www/svelte/public"))
	router.PathPrefix("/").Handler(fs)

}
