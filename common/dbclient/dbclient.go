package dbclient

import (
	"database/sql"
	"fmt"
)

func Open(db *sql.DB, host, port, loginpassword, name string) (*sql.DB, error) {
	if db == nil {
		connectString := fmt.Sprintf("postgres://%s@%s:%s/%s", loginpassword, host, port, name)
		db_, err := sql.Open("postgres", connectString)
		if err != nil {
			return nil, err
		}
		//if db_ == nil {
		//	return nil, errors.New("db is nil")
		//}
		//if err := db.Ping(); err != nil {
		//	return nil, err
		//}
		return db_, nil
	}
	return db, nil
}
