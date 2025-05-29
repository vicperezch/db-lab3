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

func CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Student
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	result := utils.DB.Create(&req)
	if result.Error != nil {
		utils.RespondWithError(w, "Unable to create student", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, req)
}

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	var students []models.Student
	result := utils.DB.Find(&students)

	if result.Error != nil {
		utils.RespondWithError(w, "Could not retrieve students", http.StatusInternalServerError)
		log.Println(result.Error)
		return
	}

	utils.RespondWithJSON(w, students)
}

func GetStudentByIdHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.RespondWithError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var student models.Student
	utils.DB.First(&student, id)

	utils.RespondWithJSON(w, student)
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.RespondWithError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req models.Student
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	var student models.Student
	result := utils.DB.First(&student, id)
	if result.Error != nil {
		utils.RespondWithError(w, "Student not found", http.StatusBadRequest)
		return
	}

	student.Email = req.Email
	student.FirstName = req.FirstName
	student.LastName = req.LastName
	student.Gender = req.Gender
	student.Level = req.Level

	result = utils.DB.Save(&student)
	if result.Error != nil {
		utils.RespondWithError(w, "Unable to save student", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, student)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.RespondWithError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result := utils.DB.Delete(&models.Student{}, id)
	if result.Error != nil {
		utils.RespondWithError(w, "Unable to delete student", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, "Student deleted successfully")
}
