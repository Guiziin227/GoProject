package router

import (
	"api/src/router/rotas"
	"github.com/gorilla/mux"
)

// Gerar retorna um router com rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
