package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/timurkash/task_example/common/helper"
	"github.com/timurkash/task_example/handlers/db"
	"github.com/timurkash/task_example/models"
	"github.com/timurkash/task_example/uuid"
	"net/http"
	"time"
)

func GetIndexRoute() helper.Route {
	return helper.Route{
		"Index",
		"GET",
		"/",
		Index,
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	url := "https://github.com/timurkash/task_example"
	bytes := []byte("<a target='_blank' href='" + url + "'>" + url + "</a>")
	(w).Write(bytes)
}

func GetAddRoute() helper.Route {
	return helper.Route{
		"Add",
		"POST",
		"/task",
		Add,
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	guid := uuid.Generate()
	err := db.AddTask(guid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	addModel := models.AddModel{Guid: guid}
	bytes, err := json.Marshal(addModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(bytes)
	go func() {
		db.UpdateStatus(guid, "running")
		time.AfterFunc(2*time.Minute, func() {
			db.UpdateStatus(guid, "finished")
		})
	}()
}

func GetTaskRoute() helper.Route {
	return helper.Route{
		"Get",
		"GET",
		"/task/{guid}",
		Get,
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	if !uuid.IsValidUUID(guid) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	isExists, err := db.IsExists(guid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isExists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	taskModelSQL, err := db.GetTask(guid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	taskModel := models.TaskModel{Status: taskModelSQL.Status, Timestamp: taskModelSQL.Timestamp.Format(time.RFC3339)}
	bytes, err := json.Marshal(taskModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	//w.WriteHeader(http.StatusOK)
}

func GetCreateTableRoute() helper.Route {
	return helper.Route{
		"CreateTable",
		"GET",
		"/createTable",
		CreateTable,
	}
}

func CreateTable(w http.ResponseWriter, r *http.Request) {
	err := db.CreateTable()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
