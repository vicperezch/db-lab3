package handlers

import (
	"lab3-backend/models"
	"lab3-backend/utils"
	"log"
	"net/http"
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
