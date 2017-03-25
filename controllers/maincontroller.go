package controllers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/alexkavon/loaf/models"
	"gitlab.com/alexkavon/loaf/utils"
)

type MainCtrl struct {
	*Controller
}

func Feed(w http.ResponseWriter, r *http.Request) {
	channel := models.Channel{Name: "Bob"}
	channel.Name = "John"

	utils.RenderTemplate(w, "index.tmpl", channel)
}

func Discover(w http.ResponseWriter, r *http.Request) {
	channel := models.Channel{Name: "Discover Channel"}

	json.NewEncoder(w).Encode(channel)
}
