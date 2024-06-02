package main

import (
	"log"
	"net/http"
	"zota-project/config"
	web "zota-project/web"
	"zota-project/services"
)

func main() {
	cfg := config.LoadConfig()

	depositService := services.NewDepositService(cfg.SecretKey)
	statusService := services.NewStatusService(cfg.SecretKey)

	handler := web.NewHandler(depositService, statusService)

	http.HandleFunc("/deposit", handler.DepositHandler)
	http.HandleFunc("/status", handler.StatusHandler)
	http.HandleFunc("/poll-status", handler.PollStatusHandler)

	http.Handle("/", web.LoggingMiddleware(http.DefaultServeMux))

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}