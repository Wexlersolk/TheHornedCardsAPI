package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Wexler763/CoffeeApiSecond/coffee-server/db"
	"github.com/Wexler763/CoffeeApiSecond/coffee-server/router"
	"github.com/Wexler763/CoffeeApiSecond/coffee-server/services"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

func (app *Application) Serve() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading anv file")
	}

	var port = os.Getenv("PORT") //Global port variable

	fmt.Println("Api is listening on port", port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}

	return srv.ListenAndServe()

}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading anv file")
	}

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgress(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	defer dbConn.DB.Close()

	app := &Application{
		Config: cfg,
		Models: services.New(dbConn.DB),
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
