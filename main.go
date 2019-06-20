package main

import (
	"github.com/timurkash/task_example/consts"
	"github.com/timurkash/task_example/routers"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started on :" + consts.TASKS_PORT)
	log.Fatal(http.ListenAndServe(":"+consts.TASKS_PORT, routers.GetRouter()))
}
