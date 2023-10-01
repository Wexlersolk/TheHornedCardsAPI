package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Wexler763/CoffeeApiSecond/coffee-server/helpers"
	"github.com/Wexler763/CoffeeApiSecond/coffee-server/services"
)

var coffee services.Coffee

func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	all, err := coffee.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
}

func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	coffeeCreated, err := coffee.CreateCoffee(coffeeData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Panicln(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, coffeeCreated)
}
