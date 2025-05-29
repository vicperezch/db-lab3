package handlers

import (
	"lab3-backend/models"
	"lab3-backend/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetEnrollmentsHandler(w http.ResponseWriter, r *http.Request) {
	var enrollments []models.EnrollmentInfo
	result := utils.DB.Find(&enrollments)

	if result.Error != nil {
		utils.RespondWithError(w, "Could not retrieve enrollments", http.StatusInternalServerError)
		log.Println(result.Error)
		return
	}

	utils.RespondWithJSON(w, enrollments)
}

func GetEnrollmentByIdHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.RespondWithError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var enrollment models.EnrollmentInfo
	utils.DB.First(&enrollment, id)

	utils.RespondWithJSON(w, enrollment)
}
