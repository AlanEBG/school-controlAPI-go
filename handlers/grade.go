package handlers

import (
	"net/http"
	"strconv"

	"github.com/AlanEBG/school-controlAPI-go/database"
	"github.com/AlanEBG/school-controlAPI-go/models"
	"github.com/AlanEBG/school-controlAPI-go/utils"
	"github.com/gin-gonic/gin"
)

// CreateGrade crea una nueva calificación
func CreateGrade(c *gin.Context) {
	var grade models.Grade

	if err := c.ShouldBindJSON(&grade); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Datos invalidos: "+err.Error())
		return
	}

	// Validar calificación
	if !utils.ValidateGrade(grade.Grade) {
		utils.ErrorResponse(c, http.StatusBadRequest, "La calificacion debe estar entre 0 y 100")
		return
	}

	// Validar que el estudiante exista
	var student models.Student
	if err := database.DB.First(&student, grade.StudentID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Estudiante no encontrado")
		return
	}

	// Validar que la materia exista
	var subject models.Subject
	if err := database.DB.First(&subject, grade.SubjectID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Materia no encontrada")
		return
	}

	// Crear calificación en la BD
	if err := database.DB.Create(&grade).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error al crear calificacion: "+err.Error())
		return
	}

	// Cargar relaciones para la respuesta
	database.DB.Preload("Student").Preload("Subject").First(&grade, grade.GradeID)

	utils.SuccessResponse(c, http.StatusCreated, grade)
}

// UpdateGrade actualiza una calificación
func UpdateGrade(c *gin.Context) {
	id := c.Param("grade_id")
	gradeID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID invalido")
		return
	}

	var grade models.Grade
	if err := database.DB.First(&grade, gradeID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Calificacion no encontrada")
		return
	}

	var updateData models.Grade
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Datos invalidos: "+err.Error())
		return
	}

	// Validar calificación si se proporciona
	if updateData.Grade != 0 {
		if !utils.ValidateGrade(updateData.Grade) {
			utils.ErrorResponse(c, http.StatusBadRequest, "La calificacion debe estar entre 0 y 100")
			return
		}
		grade.Grade = updateData.Grade
	}

	// Validar estudiante si se proporciona
	if updateData.StudentID != 0 {
		var student models.Student
		if err := database.DB.First(&student, updateData.StudentID).Error; err != nil {
			utils.ErrorResponse(c, http.StatusNotFound, "Estudiante no encontrado")
			return
		}
		grade.StudentID = updateData.StudentID
	}

	// Validar materia si se proporciona
	if updateData.SubjectID != 0 {
		var subject models.Subject
		if err := database.DB.First(&subject, updateData.SubjectID).Error; err != nil {
			utils.ErrorResponse(c, http.StatusNotFound, "Materia no encontrada")
			return
		}
		grade.SubjectID = updateData.SubjectID
	}

	if err := database.DB.Save(&grade).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error al actualizar calificacion")
		return
	}

	// Cargar relaciones para la respuesta
	database.DB.Preload("Student").Preload("Subject").First(&grade, grade.GradeID)

	utils.SuccessResponse(c, http.StatusOK, grade)
}

// DeleteGrade elimina una calificación
func DeleteGrade(c *gin.Context) {
	id := c.Param("grade_id")
	gradeID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID invalido")
		return
	}

	var grade models.Grade
	if err := database.DB.First(&grade, gradeID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Calificacion no encontrada")
		return
	}

	if err := database.DB.Delete(&grade).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error al eliminar calificacion")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Calificacion eliminada correctamente"})
}

// GetGradeByStudentAndSubject obtiene la calificación de un estudiante en una materia
func GetGradeByStudentAndSubject(c *gin.Context) {
	gradeID := c.Param("grade_id")
	studentID := c.Param("student_id")

	gID, err1 := strconv.Atoi(gradeID)
	sID, err2 := strconv.Atoi(studentID)

	if err1 != nil || err2 != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "IDs invalidos")
		return
	}

	var grade models.Grade
	if err := database.DB.Preload("Student").Preload("Subject").
		Where("grade_id = ? AND student_id = ?", gID, sID).
		First(&grade).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Calificacion no encontrada")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, grade)
}

// GetAllGradesByStudent obtiene todas las calificaciones de un estudiante
func GetAllGradesByStudent(c *gin.Context) {
	studentID := c.Param("student_id")
	sID, err := strconv.Atoi(studentID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID invalido")
		return
	}

	// Verificar que el estudiante exista
	var student models.Student
	if err := database.DB.First(&student, sID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Estudiante no encontrado")
		return
	}

	var grades []models.Grade
	if err := database.DB.Preload("Subject").
		Where("student_id = ?", sID).
		Find(&grades).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error al obtener calificaciones")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, grades)
}
