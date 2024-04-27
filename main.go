package main

import (
	"chat_app/internal/controllers"
	"chat_app/internal/repositories"
	"chat_app/internal/services"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Name     string `json:"name"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"database"`
}

func LoadConfig(filename string) Config {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
	return config
}

func main() {
	cfg := LoadConfig("config/config.json")

	connectionString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, cfg.Database.User, cfg.Database.Password)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)

	userService := services.NewUserService(userRepo)

	userController := controllers.NewUserController(userService)

	router := mux.NewRouter()

	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
