package helper

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type (
	Route struct {
		Name        string
		Methods     string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}
)

const (
	OPTIONS = "," + http.MethodOptions
)

func NewRouter(routes *[]Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range *routes {
		var handler http.Handler
		handler = route.HandlerFunc
		split := strings.Split(route.Methods+OPTIONS, ",")
		router.
			Methods(split...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
