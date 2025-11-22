package handlers

import (
	"net/http"
	"strconv"

	"github.com/AlanEBG/school-controlAPI-go/database"
	"github.com/AlanEBG/school-controlAPI-go/models"
	"github.com/AlanEBG/school-controlAPI-go/utils"
	"github.com/gin-gonic/gin"
)

// CreateSubject crea una nueva materia
func CreateSubject(c *gin.Context) {
	var subject models.Subject

	if err := c.ShouldBindJSON(&subject); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Datos invalidos: "+err.Error())
		return
	}

	// Validación del nombre
	if !utils.ValidateName(subject.Name) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Nombre de materia invalido")
		return
	}

	// Crear materia en la BD
	if err := database.DB.Create(&subject).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error al crear materia: "+err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, subject)
}

// GetSubject obtiene una materia por ID
func GetSubject(c *gin.Context) {
	id := c.Param("subject_id")
	subjectID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID invalido")
		return
	}

	var subject models.Subject
	if err := database.DB.First(&subject, subjectID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Materia no encontrada")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, subject)
}

// UpdateSubject actualiza una materia
func UpdateSubject(c *gin.Context) {
	id := c.Param("subject_id")
	subjectID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID invalido")
		return
	}

	var subject models.Subject
	if err := database.DB.First(&subject, subjectID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Materia no encontrada")
		return
	}

	var updateData models.Subject
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Datos invalidos: "+err.Error())
		return
	}

	// Validación
	if updateData.Name != "" && !utils.ValidateName(updateData.Name) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Nombre de materia invalido")
		return
	}

	// Actualizar campo
	if updateData.Name != "" {
		subject.Name = updateData.Name
	}

	if err := database.DB.Save(&subject).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error al actualizar materia")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, subject)
}

// DeleteSubject elimina una materia
func DeleteSubject(c *gin.Context) {
	id := c.Param("subject_id")
	subjectID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID invalido")
		return
	}

	var subject models.Subject
	if err := database.DB.First(&subject, subjectID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Materia no encontrada")
		return
	}

	if err := database.DB.Delete(&subject).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error al eliminar materia")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Materia eliminada correctamente"})
}
