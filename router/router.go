package router

import (
	"github.com/AlanEBG/school-controlAPI-go/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRouter configura todas las rutas de la API
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Grupo de rutas API
	api := router.Group("/api")
	{
		// Rutas de estudiantes
		students := api.Group("/students")
		{
			students.POST("", handlers.CreateStudent)
			students.GET("", handlers.GetAllStudents)
			students.GET("/:student_id", handlers.GetStudent)
			students.PUT("/:student_id", handlers.UpdateStudent)
			students.DELETE("/:student_id", handlers.DeleteStudent)
		}

		// Rutas de materias
		subjects := api.Group("/subjects")
		{
			subjects.POST("", handlers.CreateSubject)
			subjects.GET("/:subject_id", handlers.GetSubject)
			subjects.PUT("/:subject_id", handlers.UpdateSubject)
			subjects.DELETE("/:subject_id", handlers.DeleteSubject)
		}

		// Rutas de calificaciones
		grades := api.Group("/grades")
		{
			grades.POST("", handlers.CreateGrade)
			grades.PUT("/:grade_id", handlers.UpdateGrade)
			grades.DELETE("/:grade_id", handlers.DeleteGrade)
			grades.GET("/:grade_id/student/:student_id", handlers.GetGradeByStudentAndSubject)
			grades.GET("/student/:student_id", handlers.GetAllGradesByStudent)
		}
	}

	return router
}
