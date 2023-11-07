package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Wexler763/TheHornedCardsAPI/coffee-server/helpers"
	"github.com/Wexler763/TheHornedCardsAPI/coffee-server/services"
	"github.com/go-chi/chi/v5"
)

var card services.Card
var group services.Group

// GET/coffees
func GetAllCards(w http.ResponseWriter, r *http.Request) {
	all, err := card.GetAllCards()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"cards": all})
}

// GET//coffees/coffee/{id}
func GetCardById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	card, err := card.GetCardById(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, card)
}

// POST/coffees/coffee
func CreateCard(w http.ResponseWriter, r *http.Request) {
	var cardData services.Card
	err := json.NewDecoder(r.Body).Decode(&cardData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	cardCreated, err := card.CreateCard(cardData)
	// CHECK
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, cardCreated)
}

// PUT/coffees/coffee/{id}
func UpdateCard(w http.ResponseWriter, r *http.Request) {
	var cardData services.Card
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&cardData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cardUpdated, err := card.UpdateCard(id, cardData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, cardUpdated)
}

// DELETE/coffees/coffee/{id}
func DeleteCard(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := card.DeleteCard(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "successfull deletion"})
}

func DeleteAllCards(w http.ResponseWriter, r *http.Request) {
	err := card.DeleteAllCards()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Envelope{"error": "Failed to delete all cards"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "All cards deleted"})
}

// --------------------------------------- GROUPS ---------------------------------------------------

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	var groupData services.Group
	err := json.NewDecoder(r.Body).Decode(&groupData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	groupCreated, err := group.CreateGroup(groupData)
	// CHECK
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, groupCreated)
}

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	all, err := group.GetAllGroups()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"groups": all})
}

func DeleteGroupById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := group.DeleteGroupById(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "successfull deletion"})
}

func DeleteGroupByName(w http.ResponseWriter, r *http.Request) {
	group_name := chi.URLParam(r, "group_name")
	err := group.DeleteGroupByName(group_name)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "successfull deletion"})
}

func DeleteAllGroups(w http.ResponseWriter, r *http.Request) {
	err := group.DeleteAllGroups()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.Envelope{"error": "Failed to delete all groups"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "All cards groups"})
}
