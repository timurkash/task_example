package routers

import (
	"github.com/gorilla/mux"
	"github.com/timurkash/task_example/common/helper"
	"github.com/timurkash/task_example/handlers"
)

var (
	routes = []helper.Route{
		handlers.GetAddRoute(),
		handlers.GetTaskRoute(),
		handlers.GetCreateTableRoute(),
		handlers.GetIndexRoute(),
	}
)

func GetRouter() *mux.Router {
	return helper.NewRouter(&routes)
}
