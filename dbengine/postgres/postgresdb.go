package postgres

import (
	"github.com/Shre1892YNWA/LibraryInventorySystem/dbengine"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgresDatabase struct {
	*sqlx.DB
}

func GetPgDatabaseEngine(dataSourceString string) (dbengine.DBEngine, error) {

	pgDB, err := sqlx.Open("postgres", dataSourceString)
	if err != nil {
		return nil, err
	}

	err = pgDB.Ping()
	if err != nil {
		return nil, err
	}

	return &postgresDatabase{
		pgDB,
	}, nil
}
