package postgre

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgre struct {
	// We can add multiple database connection by creating new struct here
	// and add the connection string on initialization of postgre repository package
	newsDatabase *PgDB
}

func New(
	newsDatabaseConnectionString string,
) (*Postgre, error) {
	var err error

	newsDatabase, err := InitPostgreDB(newsDatabaseConnectionString, "")
	if err != nil {
		return nil, fmt.Errorf("[Postgre][Init] Failed init user database, trace %v", err)
	}

	return &Postgre{
		newsDatabase: &newsDatabase,
	}, nil
}

type PgDB struct {
	Master *sqlx.DB

	// Slave added on case if we need to add slave to its database due to performance issue
	Slave *sqlx.DB
}

func InitPostgreDB(
	masterDriver string,
	slaveDriver string,
) (db PgDB, err error) {
	db.Master, err = sqlx.Connect(DB_DRIVER_NAME_POSTGRE, masterDriver)
	if err != nil {
		return db, err
	}

	if slaveDriver != "" {
		db.Slave, err = sqlx.Connect(DB_DRIVER_NAME_POSTGRE, slaveDriver)
		if err != nil {
			return db, err
		}
	}

	return db, err
}
