package routes

import (
	"Groupietracker/controlers"
	"net/http"
)

func SetupConnectionRoutes() {
	http.HandleFunc("/connection", controlers.ConnectionControler)
}
