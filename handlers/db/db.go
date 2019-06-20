package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/timurkash/task_example/common/config"
	"github.com/timurkash/task_example/common/dbclient"
	"github.com/timurkash/task_example/models"
)

var (
	DBHOST          = config.GetEnv("DBHOST", "postgres")
	DBPORT          = config.GetEnv("DBPORT", "5432")
	DBLOGINPASSWORD = config.GetEnv("DBLOGINPASSWORD", "postgres:super123")
	DBNAME          = config.GetEnv("DBNAME", "postgres?sslmode=disable")
	DBSCHEME        = config.GetEnv("DBSCHEME", "task_example")
	db              *sql.DB
)

func GetTask(guid string) (*models.TaskModelSQL, error) {
	taskModelSQL := models.TaskModelSQL{}
	db, err := dbclient.Open(db, DBHOST, DBPORT, DBLOGINPASSWORD, DBNAME)
	if err != nil {
		return nil, err
	}
	err = db.QueryRow("SELECT status,uDate FROM "+DBSCHEME+".tasks where guid = $1", guid).Scan(&taskModelSQL.Status, &taskModelSQL.Timestamp)
	if err != nil {
		return nil, err
	}
	return &taskModelSQL, nil
}

func AddTask(guid string) error {
	db, err := dbclient.Open(db, DBHOST, DBPORT, DBLOGINPASSWORD, DBNAME)
	if err != nil {
		return err
	}
	status := "created"
	_, err = db.Exec("INSERT INTO "+DBSCHEME+".tasks (guid, status) values ($1, $2)", guid, status)
	if err != nil {
		return err
	}
	return nil
}

func IsExists(guid string) (bool, error) {
	var (
		count int
		err   error
	)
	db, err = dbclient.Open(db, DBHOST, DBPORT, DBLOGINPASSWORD, DBNAME)
	if err != nil {
		return false, err
	}
	err = db.QueryRow("SELECT COUNT(*) FROM "+DBSCHEME+".tasks where guid = $1", guid).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

func DeleteTask(guid string) error {
	db, err := dbclient.Open(db, DBHOST, DBPORT, DBLOGINPASSWORD, DBNAME)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE from "+DBSCHEME+".tasks where guid = $1", guid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateStatus(guid string, status string) error {
	db, err := dbclient.Open(db, DBHOST, DBPORT, DBLOGINPASSWORD, DBNAME)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE  "+DBSCHEME+".tasks set status = $2, uDate = now() where guid = $1", guid, status)
	if err != nil {
		return err
	}
	return nil
}

func CreateTable() error {
	db, err := dbclient.Open(db, DBHOST, DBPORT, DBLOGINPASSWORD, DBNAME)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE SCHEMA ` + DBSCHEME)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE ` + DBSCHEME + `.tasks (
	guid uuid NOT NULL,
	status varchar(10) NOT NULL,
	cdate timestamp NOT NULL DEFAULT now(),
	udate timestamp NOT NULL DEFAULT now(),
	CONSTRAINT tasks_pk PRIMARY KEY (guid)
);`)
	if err != nil {
		return err
	}
	return nil
}
