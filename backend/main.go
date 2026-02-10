package main

import (
	"LearningCampusControlContinu/config"
	"fmt"
	"net/http"
)

func main() {
	config.ConnectDB()
	fmt.Println("Serveur démarré sur hhtp://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
