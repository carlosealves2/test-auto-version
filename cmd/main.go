package main

import (
	"fmt"
	"net/http"

	"github.com/carlosealves2/test-auto-version/configs"
	"github.com/carlosealves2/test-auto-version/internal/api"
	"github.com/carlosealves2/test-auto-version/internal/service"
)

func main() {
	fmt.Print("iniciando...")

	cfg := configs.NewConfigBuilder().WithEnv().Validate().Build()

	mux := http.NewServeMux()
	calculatorService := service.NewCalculatorService()
	apiRouter := api.NewAPIRouter(calculatorService)
	apiRouter.RegisterRoutes(mux)

	fmt.Println("Server running on port", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, mux)

}
