package handlers

import (
	"net/http"
	"strconv"

	"github.com/AlanEBG/school-controlAPI-go/database"
	"github.com/AlanEBG/school-controlAPI-go/models"
	"github.com/AlanEBG/school-controlAPI-go/utils"
	"github.com/gin-gonic/gin"
)

// Crear estudiante
func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Datos invalidos: "+err.Error())
		return
	}

	// Validaciones adicionales
	if !utils.ValidateEmail(student.Email) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Formato de email invalido")
		return
	}

	if !utils.ValidateName(student.Name) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Nombre invalido")
		return
	}

	if !utils.ValidateGroup(student.Group) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Grupo invalido")
		return
	}

	// Crear estudiante en la BD
	if err := database.DB.Create(&student).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error al crear estudiante: "+err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, student)
}

// GetAllStudents obtiene todos los estudiantes
func GetAllStudents(c *gin.Context) {
	var students []models.Student

	if err := database.DB.Find(&students).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error al obtener estudiantes")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, students)
}

// GetStudent obtiene un estudiante por ID
func GetStudent(c *gin.Context) {
	id := c.Param("student_id")
	studentID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID invalido")
		return
	}

	var student models.Student
	if err := database.DB.First(&student, studentID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Estudiante no encontrado")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, student)
}

// UpdateStudent actualiza un estudiante
func UpdateStudent(c *gin.Context) {
	id := c.Param("student_id")
	studentID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID invalido")
		return
	}

	var student models.Student
	if err := database.DB.First(&student, studentID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Estudiante no encontrado")
		return
	}

	var updateData models.Student
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Datos invalidos: "+err.Error())
		return
	}

	// Validaciones
	if updateData.Email != "" && !utils.ValidateEmail(updateData.Email) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Formato de email invalido")
		return
	}

	if updateData.Name != "" && !utils.ValidateName(updateData.Name) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Nombre invalido")
		return
	}

	if updateData.Group != "" && !utils.ValidateGroup(updateData.Group) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Grupo invalido")
		return
	}

	// Actualizar campos
	if updateData.Name != "" {
		student.Name = updateData.Name
	}
	if updateData.Group != "" {
		student.Group = updateData.Group
	}
	if updateData.Email != "" {
		student.Email = updateData.Email
	}

	if err := database.DB.Save(&student).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error al actualizar estudiante")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, student)
}

// DeleteStudent elimina un estudiante
func DeleteStudent(c *gin.Context) {
	id := c.Param("student_id")
	studentID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID invalido")
		return
	}

	var student models.Student
	if err := database.DB.First(&student, studentID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Estudiante no encontrado")
		return
	}

	if err := database.DB.Delete(&student).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error al eliminar estudiante")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Estudiante eliminado correctamente"})
}
