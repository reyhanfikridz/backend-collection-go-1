/*
Package student test
contain testing for package student
*/
package student_test

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/reyhanfikridz/backend-collection-go-1/internal/student"
	"github.com/reyhanfikridz/backend-collection-go-1/internal/test_utils"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
)

// TestMain
func TestMain(m *testing.M) {
	// get test app
	app, err := test_utils.GetTestApp()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// delete student data before all running all test
	err = app.DB.Delete(student.Student{},
		"student_number IN ?", []int{
			20182304981, 20182304982, 20182304983, 20182304984, 20182304985,
			20182304986, 20182304987, 20182304988, 20182304989, 201823049810,
		}).Error
	if err != nil {
		log.Fatalf(err.Error())
	}

	// run all test in the package
	m.Run()

	// delete student data before all running all test
	err = app.DB.Delete(student.Student{},
		"student_number IN ?", []int{
			20182304981, 20182304982, 20182304983, 20182304984, 20182304985,
			20182304986, 20182304987, 20182304988, 20182304989, 201823049810,
		}).Error
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// TestAddStudent test API with handler AddStudent
func TestAddStudent(t *testing.T) {
	// get test app
	app, err := test_utils.GetTestApp()
	if err != nil {
		t.Fatalf(err.Error())
	}

	// test API add student
	apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Post("/api/students").
		JSON(student.Student{
			StudentNumber: 20182304981,
			FullName:      "Tom Bradney",
			FullAddress:   "Tom Street",
			YearOfEnroll:  2018,
		}).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

// TestGetStudent test API with handler GetStudent
func TestGetStudent(t *testing.T) {
	// get test app
	app, err := test_utils.GetTestApp()
	if err != nil {
		t.Fatalf(err.Error())
	}

	// add student first
	//// call API add student
	result := apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Post("/api/students").
		JSON(student.Student{
			StudentNumber: 20182304982,
			FullName:      "Tom Bradney 2",
			FullAddress:   "Tom Street 2",
			YearOfEnroll:  2018,
		}).
		Expect(t).
		Status(http.StatusCreated).
		End()

	//// get added student data
	var addedStudent student.Student
	err = test_utils.ParseDataFromApitestResult(result, &addedStudent)
	if err != nil {
		t.Fatalf(err.Error())
	}

	addedStudent.CreatedAt = addedStudent.CreatedAt.Round(time.Millisecond)
	addedStudent.UpdatedAt = addedStudent.UpdatedAt.Round(time.Millisecond)

	// test API get student
	result = apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Get(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		Expect(t).
		Status(http.StatusOK).
		End()

	var gettedStudent student.Student
	err = test_utils.ParseDataFromApitestResult(result, &gettedStudent)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, addedStudent, gettedStudent)
}

// TestGetStudents test API with handler GetStudents
func TestGetStudents(t *testing.T) {
	// get test app
	app, err := test_utils.GetTestApp()
	if err != nil {
		t.Fatalf(err.Error())
	}

	// add student first
	apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Post("/api/students").
		JSON(student.Student{
			StudentNumber: 20182304983,
			FullName:      "Tom Bradney 3",
			FullAddress:   "Tom Street 3",
			YearOfEnroll:  2018,
		}).
		Expect(t).
		Status(http.StatusCreated).
		End()

	apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Post("/api/students").
		JSON(student.Student{
			StudentNumber: 20182304984,
			FullName:      "Tom Bradney 4",
			FullAddress:   "Tom Street 4",
			YearOfEnroll:  2018,
		}).
		Expect(t).
		Status(http.StatusCreated).
		End()

	apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Post("/api/students").
		JSON(student.Student{
			StudentNumber: 20182304985,
			FullName:      "Tom Bradney 5",
			FullAddress:   "Tom Street 5",
			YearOfEnroll:  2019,
		}).
		Expect(t).
		Status(http.StatusCreated).
		End()

	// test API get students
	result := apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Get("/api/students").
		QueryParams(map[string]string{
			"year_of_enroll": "2018",
		}).
		Expect(t).
		Status(http.StatusOK).
		End()

	var gettedStudents []student.Student
	err = test_utils.ParseDataFromApitestResult(result, &gettedStudents)
	if err != nil {
		t.Fatalf(err.Error())
	}

	for _, student := range gettedStudents {
		assert.Equal(t, student.YearOfEnroll, 2018)
	}
}

// TestReplaceStudent test API with handler ReplaceStudent
func TestReplaceStudent(t *testing.T) {
	// get test app
	app, err := test_utils.GetTestApp()
	if err != nil {
		t.Fatalf(err.Error())
	}

	// add student first
	//// call API add student
	result := apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Post("/api/students").
		JSON(student.Student{
			StudentNumber: 20182304986,
			FullName:      "Tom Bradney 6",
			FullAddress:   "Tom Street 6",
			YearOfEnroll:  2018,
		}).
		Expect(t).
		Status(http.StatusCreated).
		End()

	//// get added student data
	var addedStudent student.Student
	err = test_utils.ParseDataFromApitestResult(result, &addedStudent)
	if err != nil {
		t.Fatalf(err.Error())
	}

	addedStudent.CreatedAt = addedStudent.CreatedAt.Round(time.Millisecond)
	addedStudent.UpdatedAt = addedStudent.UpdatedAt.Round(time.Millisecond)

	// test API replace student
	//// call API replace student
	result = apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Put(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		JSON(student.Student{
			StudentNumber: 20182304987,
			FullName:      "Tom Bradney 7",
			FullAddress:   "Tom Street 7",
			YearOfEnroll:  2022,
		}).
		Expect(t).
		Status(http.StatusOK).
		End()

	var studentAfterReplaced student.Student
	err = test_utils.ParseDataFromApitestResult(result, &studentAfterReplaced)
	if err != nil {
		t.Fatalf(err.Error())
	}

	studentAfterReplaced.CreatedAt = studentAfterReplaced.CreatedAt.
		Round(time.Millisecond)
	studentAfterReplaced.UpdatedAt = studentAfterReplaced.UpdatedAt.
		Round(time.Millisecond)

	//// get student for comparison
	result = apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Get(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		Expect(t).
		Status(http.StatusOK).
		End()

	var gettedStudent student.Student
	err = test_utils.ParseDataFromApitestResult(result, &gettedStudent)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, gettedStudent, studentAfterReplaced)
}

// TestUpdateStudent test API with handler UpdateStudent
func TestUpdateStudent(t *testing.T) {
	// get test app
	app, err := test_utils.GetTestApp()
	if err != nil {
		t.Fatalf(err.Error())
	}

	// add student first
	//// call API add student
	result := apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Post("/api/students").
		JSON(student.Student{
			StudentNumber: 20182304988,
			FullName:      "Tom Bradney 8",
			FullAddress:   "Tom Street 8",
			YearOfEnroll:  2018,
		}).
		Expect(t).
		Status(http.StatusCreated).
		End()

	//// get added student data
	var addedStudent student.Student
	err = test_utils.ParseDataFromApitestResult(result, &addedStudent)
	if err != nil {
		t.Fatalf(err.Error())
	}

	addedStudent.CreatedAt = addedStudent.CreatedAt.Round(time.Millisecond)
	addedStudent.UpdatedAt = addedStudent.UpdatedAt.Round(time.Millisecond)

	// test API update student
	//// test update student_number only
	result = apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Patch(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		JSON(map[string]any{"student_number": 20182304989}).
		Expect(t).
		Status(http.StatusOK).
		End()

	var studentAfterUpdated student.Student
	err = test_utils.ParseDataFromApitestResult(result, &studentAfterUpdated)
	if err != nil {
		t.Fatalf(err.Error())
	}

	studentAfterUpdated.CreatedAt = studentAfterUpdated.CreatedAt.
		Round(time.Millisecond)
	studentAfterUpdated.UpdatedAt = studentAfterUpdated.UpdatedAt.
		Round(time.Millisecond)

	result = apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Get(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		Expect(t).
		Status(http.StatusOK).
		End()

	var gettedStudent student.Student
	err = test_utils.ParseDataFromApitestResult(result, &gettedStudent)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, gettedStudent, studentAfterUpdated)
	assert.Equal(t, gettedStudent.ID, addedStudent.ID)
	assert.NotEqual(t, gettedStudent.StudentNumber, addedStudent.StudentNumber)
	assert.Equal(t, gettedStudent.FullName, addedStudent.FullName)
	assert.Equal(t, gettedStudent.FullAddress, addedStudent.FullAddress)
	assert.Equal(t, gettedStudent.YearOfEnroll, addedStudent.YearOfEnroll)

	//// test update full_name only
	result = apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Patch(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		JSON(map[string]any{"full_name": "Tom Bradney 2"}).
		Expect(t).
		Status(http.StatusOK).
		End()

	err = test_utils.ParseDataFromApitestResult(result, &studentAfterUpdated)
	if err != nil {
		t.Fatalf(err.Error())
	}

	studentAfterUpdated.CreatedAt = studentAfterUpdated.CreatedAt.
		Round(time.Millisecond)
	studentAfterUpdated.UpdatedAt = studentAfterUpdated.UpdatedAt.
		Round(time.Millisecond)

	result = apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Get(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		Expect(t).
		Status(http.StatusOK).
		End()

	err = test_utils.ParseDataFromApitestResult(result, &gettedStudent)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, gettedStudent, studentAfterUpdated)
	assert.Equal(t, gettedStudent.ID, addedStudent.ID)
	assert.NotEqual(t, gettedStudent.StudentNumber, addedStudent.StudentNumber)
	assert.NotEqual(t, gettedStudent.FullName, addedStudent.FullName)
	assert.Equal(t, gettedStudent.FullAddress, addedStudent.FullAddress)
	assert.Equal(t, gettedStudent.YearOfEnroll, addedStudent.YearOfEnroll)

	//// test update full_address only
	result = apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Patch(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		JSON(map[string]any{"full_address": "Tom Street 2"}).
		Expect(t).
		Status(http.StatusOK).
		End()

	err = test_utils.ParseDataFromApitestResult(result, &studentAfterUpdated)
	if err != nil {
		t.Fatalf(err.Error())
	}

	studentAfterUpdated.CreatedAt = studentAfterUpdated.CreatedAt.
		Round(time.Millisecond)
	studentAfterUpdated.UpdatedAt = studentAfterUpdated.UpdatedAt.
		Round(time.Millisecond)

	result = apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Get(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		Expect(t).
		Status(http.StatusOK).
		End()

	err = test_utils.ParseDataFromApitestResult(result, &gettedStudent)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, gettedStudent, studentAfterUpdated)
	assert.Equal(t, gettedStudent.ID, addedStudent.ID)
	assert.NotEqual(t, gettedStudent.StudentNumber, addedStudent.StudentNumber)
	assert.NotEqual(t, gettedStudent.FullName, addedStudent.FullName)
	assert.NotEqual(t, gettedStudent.FullAddress, addedStudent.FullAddress)
	assert.Equal(t, gettedStudent.YearOfEnroll, addedStudent.YearOfEnroll)

	//// test update year_of_enroll only
	result = apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Patch(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		JSON(map[string]any{"year_of_enroll": 2022}).
		Expect(t).
		Status(http.StatusOK).
		End()

	err = test_utils.ParseDataFromApitestResult(result, &studentAfterUpdated)
	if err != nil {
		t.Fatalf(err.Error())
	}

	studentAfterUpdated.CreatedAt = studentAfterUpdated.CreatedAt.
		Round(time.Millisecond)
	studentAfterUpdated.UpdatedAt = studentAfterUpdated.UpdatedAt.
		Round(time.Millisecond)

	result = apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Get(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		Expect(t).
		Status(http.StatusOK).
		End()

	err = test_utils.ParseDataFromApitestResult(result, &gettedStudent)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, gettedStudent, studentAfterUpdated)
	assert.Equal(t, gettedStudent.ID, addedStudent.ID)
	assert.NotEqual(t, gettedStudent.StudentNumber, addedStudent.StudentNumber)
	assert.NotEqual(t, gettedStudent.FullName, addedStudent.FullName)
	assert.NotEqual(t, gettedStudent.FullAddress, addedStudent.FullAddress)
	assert.NotEqual(t, gettedStudent.YearOfEnroll, addedStudent.YearOfEnroll)
}

// TestDeleteHandler test API with handler DeleteHandler
func TestDeleteHandler(t *testing.T) {
	// get test app
	app, err := test_utils.GetTestApp()
	if err != nil {
		t.Fatalf(err.Error())
	}

	// add student first
	//// call API add student
	result := apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Post("/api/students").
		JSON(student.Student{
			StudentNumber: 201823049810,
			FullName:      "Tom Bradney 10",
			FullAddress:   "Tom Street 10",
			YearOfEnroll:  2018,
		}).
		Expect(t).
		Status(http.StatusCreated).
		End()

	//// get added student data
	var addedStudent student.Student
	err = test_utils.ParseDataFromApitestResult(result, &addedStudent)
	if err != nil {
		t.Fatalf(err.Error())
	}

	addedStudent.CreatedAt = addedStudent.CreatedAt.Round(time.Millisecond)
	addedStudent.UpdatedAt = addedStudent.UpdatedAt.Round(time.Millisecond)

	// test API delete student
	//// call API delete student
	apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Delete(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		Expect(t).
		Status(http.StatusOK).
		End()

	//// call API get student
	apitest.New().
		HandlerFunc(app.Router.Handler().ServeHTTP).
		Get(fmt.Sprintf("/api/students/%d", addedStudent.ID)).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}
