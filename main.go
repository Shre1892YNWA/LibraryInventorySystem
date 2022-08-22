package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/Shre1892YNWA/LibraryInventorySystem/dbengine"
	"github.com/Shre1892YNWA/LibraryInventorySystem/dbengine/inmem"
	"github.com/Shre1892YNWA/LibraryInventorySystem/dbengine/mongo"
	"github.com/Shre1892YNWA/LibraryInventorySystem/dbengine/postgres"
	"github.com/Shre1892YNWA/LibraryInventorySystem/literals"
	"github.com/Shre1892YNWA/LibraryInventorySystem/router"
	"github.com/joho/godotenv"
)

var (
	dbName = flag.String("db", literals.DBNameInmem, "name of database")
)

func main() {
	flag.Parse()

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbEngine, err := getDatabaseByName(*dbName)
	if err != nil {
		panic(err)
	}

	appRouter, err := router.InitializeRoutes(dbEngine)
	if err != nil {
		panic(err)
	}

	fmt.Println("serving at port 8080")

	http.ListenAndServe(":8080", appRouter)
}

func getDatabaseByName(dbName string) (dbengine.DBEngine, error) {
	switch dbName {
	case literals.DBNameMongo:
		return mongo.GetMongodbEngine("mongodb://localhost:2717")
	case literals.DBNamePG:
		return postgres.GetPgDatabaseEngine(getPostgresDataString())
	default:
		return inmem.GetInMemEngine(), nil
	}
}

func getPostgresDataString() string {
	Host := os.Getenv("POSTGRES_HOST")
	Port := os.Getenv("POSTGRES_PORT")
	User := os.Getenv("POSTGRES_USER")
	Password := os.Getenv("POSTGRES_PASSWORD")
	DBname := os.Getenv("POSTGRES_DB")

	dataSourceString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DBname)

	return dataSourceString

}
