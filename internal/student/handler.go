/*
Package student test contain everything for student data,
such as model, route handler, utility, etc.
*/
package student

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler route handler for student data
type Handler struct {
	DB *gorm.DB
}

// AddStudent route handler for add student data (POST /api/students)
func (h *Handler) AddStudent(ctx *gin.Context) {
	s := Student{}

	// get student data from request json
	err := ctx.BindJSON(&s)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	// set created at and updated at
	timeNow := time.Now()
	s.CreatedAt = timeNow
	s.UpdatedAt = timeNow

	// add student data to database
	result := h.DB.Create(&s)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, s)
}

// GetStudent route handler for get student data (GET /api/students/:id)
func (h *Handler) GetStudent(ctx *gin.Context) {
	s := Student{}

	// get student data
	err := h.DB.First(&s, "id = ?", ctx.Param("id")).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, s)
}

// GetStudents route handler for get students data (GET /api/students)
func (h *Handler) GetStudents(ctx *gin.Context) {
	s := []Student{}

	// get students data
	filter := make(map[string]string)
	if ctx.Query("year_of_enroll") != "" {
		filter["year_of_enroll"] = ctx.Query("year_of_enroll")
	}
	err := h.DB.Find(&s, filter).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, s)
}

// ReplaceStudent route handler for replace student data (PUT /api/students/:id)
func (h *Handler) ReplaceStudent(ctx *gin.Context) {
	sNew := Student{}

	// get new student data from request json
	err := ctx.BindJSON(&sNew)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	// get old student data
	s := Student{}
	err = h.DB.First(&s, "id = ?", ctx.Param("id")).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	// replace old student data with new student data
	s.StudentNumber = sNew.StudentNumber
	s.FullName = sNew.FullName
	s.FullAddress = sNew.FullAddress
	s.YearOfEnroll = sNew.YearOfEnroll
	s.UpdatedAt = time.Now()
	err = h.DB.Save(&s).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, s)
}

// UpdateStudent route handler for update student data (PATCH /api/students/:id)
func (h *Handler) UpdateStudent(ctx *gin.Context) {
	sNew := Student{}

	// get new student data from request json
	err := ctx.BindJSON(&sNew)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	// get old student data
	s := Student{}
	err = h.DB.First(&s, "id = ?", ctx.Param("id")).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	// update old student data with new student data
	if sNew.StudentNumber != 0 {
		s.StudentNumber = sNew.StudentNumber
	}
	if sNew.FullName != "" {
		s.FullName = sNew.FullName
	}
	if sNew.FullAddress != "" {
		s.FullAddress = sNew.FullAddress
	}
	if sNew.YearOfEnroll != 0 {
		s.YearOfEnroll = sNew.YearOfEnroll
	}
	s.UpdatedAt = time.Now()

	err = h.DB.Save(&s).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, s)
}

// DeleteStudent route handler for delete student data (DELETE /api/students/:id)
func (h *Handler) DeleteStudent(ctx *gin.Context) {
	// delete student data
	err := h.DB.Delete(&Student{}, "id = ?", ctx.Param("id")).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"message": "student deleted successfully",
	})
}
