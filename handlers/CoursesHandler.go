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

func CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Course
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	result := utils.DB.Create(&req)
	if result.Error != nil {
		utils.RespondWithError(w, "Unable to create course", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, req)
}

func GetCoursesHandler(w http.ResponseWriter, r *http.Request) {
	var courses []models.Course
	result := utils.DB.Find(&courses)

	if result.Error != nil {
		utils.RespondWithError(w, "Could not retrieve courses", http.StatusInternalServerError)
		log.Println(result.Error)
		return
	}

	utils.RespondWithJSON(w, courses)
}

func GetCourseByIdHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.RespondWithError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var course models.Course
	utils.DB.First(&course, id)

	utils.RespondWithJSON(w, course)
}

func UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.RespondWithError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req models.Course
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	var course models.Course
	result := utils.DB.First(&course, id)
	if result.Error != nil {
		utils.RespondWithError(w, "Course not found", http.StatusBadRequest)
		return
	}

	course.Description = req.Description
	course.Mode = req.Mode
	course.Name = req.Name

	result = utils.DB.Save(&course)
	if result.Error != nil {
		utils.RespondWithError(w, "Unable to save course", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, course)
}

func DeleteCourseHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.RespondWithError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result := utils.DB.Delete(&models.Course{}, id)
	if result.Error != nil {
		utils.RespondWithError(w, "Unable to delete course", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, "Course deleted successfully")
}
