package main

import (
	"log"
	"net/http"
	"strings"

	"categories/database"
	"categories/handlers"
	"categories/repositories"
	"categories/services"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	// Load .env into environment variables
	_ = godotenv.Load()

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// log.Println("Config:", config)

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc("/api/products", productHandler.HandleProducts)
	http.HandleFunc("/api/products/", productHandler.HandleProductByID)

	log.Println("Server running on port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
