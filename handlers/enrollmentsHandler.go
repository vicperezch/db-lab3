package handlers

import (
	"encoding/json"
	"lab3-backend/models"
	"lab3-backend/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func CreateEnrollmentHandler(w http.ResponseWriter, r *http.Request) {
	var req models.EnrollmentDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	enrollment := models.Enrollment{
		StudentID:      req.StudentID,
		CourseID:       req.CourseID,
		EnrollmentDate: req.EnrollmentDate,
	}

	result := utils.DB.Create(&enrollment)
	if result.Error != nil {
		utils.RespondWithError(w, "Unable to create enrollment", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, enrollment)
}

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

func UpdateEnrollmentHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.RespondWithError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req models.EnrollmentDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	var enrollment models.Enrollment
	result := utils.DB.First(&enrollment, id)
	if result.Error != nil {
		utils.RespondWithError(w, "Enrollment not found", http.StatusBadRequest)
		return
	}

	enrollment.StudentID = req.StudentID
	enrollment.CourseID = req.CourseID
	enrollment.EnrollmentDate = req.EnrollmentDate

	result = utils.DB.Save(&enrollment)
	if result.Error != nil {
		utils.RespondWithError(w, "Unable to save enrollment", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, enrollment)
}

func DeleteEnrollmentHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.RespondWithError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result := utils.DB.Delete(&models.Enrollment{}, id)
	if result.Error != nil {
		utils.RespondWithError(w, "Unable to delete enrollment", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, "Enrollment deleted successfully")
}
