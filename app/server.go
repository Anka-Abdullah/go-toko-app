package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Anka-Abdullah/Go-toko-1/database/seeders"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("wellcome to" + appConfig.AppName)

	// server.InitializeDB(dbConfig)
	server.InitializeRoutes()
	// seeders.DBSeed(server.DB)
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) InitializeDB(dbConfig DBConfig) {
	host := dbConfig.DBHost
	dbuser := dbConfig.DBUser
	password := dbConfig.DBPassword
	dbname := dbConfig.DBName
	port := dbConfig.DBPort

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, dbuser, password, dbname, port)

	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed on connecting to DB")
	}

}

func (server *Server) dbMigrate() {
	for _, model := range RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully")
}

func (server *Server) initCommands(config AppConfig, dbConfig DBConfig) {
	server.InitializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}
	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Run() {
	server := Server{}
	appConfig := AppConfig{}
	dbConfig := DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppEnv = os.Getenv("APP_ENV")
	appConfig.AppPort = os.Getenv("APP_PORT")

	dbConfig.DBHost = os.Getenv("DB_HOST")
	dbConfig.DBUser = os.Getenv("DB_USER")
	dbConfig.DBPassword = os.Getenv("DB_PASSWORD")
	dbConfig.DBName = os.Getenv("DB_NAME")
	dbConfig.DBPort = os.Getenv("DB_PORT")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.initCommands(appConfig, dbConfig)
	} else {

		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}

}
