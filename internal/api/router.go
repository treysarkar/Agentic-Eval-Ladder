package api

import (
	"net/http"

	"github.com/treysarkar/Agentic-Eval-Ladder/internal/store"
)

type server struct{
	store *store.Store
}

//New router will build the whole HTTP service. Tests will depend ont his siganture
func NewRouter(s *store.Store) http.Handler{
	srv:= &server{store: s}
	mux:= http.NewServeMux()
	srv.routes(mux) //calling the method of the struct
	return mux
}

func (srv *server) routes(mux *http.ServeMux){
	// each step will register its new router here
}