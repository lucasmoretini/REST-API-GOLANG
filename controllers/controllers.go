package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api-go/database"
	"rest-api-go/models"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)

	json.NewEncoder(w).Encode(personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade

	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
}
