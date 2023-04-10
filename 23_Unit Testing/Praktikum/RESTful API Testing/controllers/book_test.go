package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restful_api_testing/database"
	"restful_api_testing/models"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var bookController BookController = InitBookContoller()

func TestGetAllBooks_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/auth/books",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"book\":",
	}

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)

	if assert.NoError(t, bookController.GetAll(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetBookById_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/auth/books/",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"book\":",
	}

	book, err := database.SeedBook()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bookID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues(bookID)

	if assert.NoError(t, bookController.GetById(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestGetBookById_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "invalid",
		path:                   "/auth/books/",
		expectedStatus:         http.StatusBadRequest,
	}

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues("100")

	if assert.NoError(t, bookController.GetById(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}
}

func TestCreateBook_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/auth/books/",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"book\":",
	}

	var bookInput models.BookInput = models.BookInput{
		Title: "dummy",
		Writer: "dummy",
		Publisher: "dummy",
	}

	jsonBody, err := json.Marshal(&bookInput)

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)

	if assert.NoError(t, bookController.Create(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestCreateBook_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "invalid",
		path:                   "/auth/books/",
		expectedStatus:         http.StatusBadRequest,
	}

	var bookInput models.BookInput = models.BookInput{
		Title: "dummy",
		Writer: "dummy",
		Publisher: "dummy",
	}

	jsonBody, err := json.Marshal(&bookInput)

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)
	//req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)

	if assert.NoError(t, bookController.Create(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}
}

func TestUpdateBookById_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/auth/books",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"book\":",
	}

	book, err := database.SeedBook()
	
	if err != nil {
		t.Errorf("error creating book: %v\n", err)
	}

	bookID := strconv.Itoa(int(book.ID))

	bookUpdate := models.BookInput{
		Title:     "new title",
		Writer:    "new writer",
		Publisher: "new publisher",
	}

	jsonBody, err := json.Marshal(&bookUpdate)

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPut, testcase.path, bodyReader)
	req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues(bookID)
	
	if assert.NoError(t, bookController.Update(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))

		updatedBook := models.Book{}
		err := database.DB.First(&updatedBook, bookID).Error
		if err != nil {
			t.Errorf("error retrieving book from database: %v\n", err)
		}
		assert.Equal(t, bookUpdate.Title, updatedBook.Title)
		assert.Equal(t, bookUpdate.Writer, updatedBook.Writer)
		assert.Equal(t, bookUpdate.Publisher, updatedBook.Publisher)
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestUpdateBookById_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "invalid",
		path:                   "/auth/books",
		expectedStatus:         http.StatusBadRequest,
	}

	book, err := database.SeedBook()
	
	if err != nil {
		t.Errorf("error creating book: %v\n", err)
	}

	bookID := strconv.Itoa(int(book.ID))

	bookUpdate := models.BookInput{
		Title:     "new title",
		Writer:    "new writer",
		Publisher: "new publisher",
	}

	jsonBody, err := json.Marshal(&bookUpdate)

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPut, testcase.path, bodyReader)
	// req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues(bookID)

	if assert.NoError(t, bookController.Update(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestDeleteBookById_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/auth/books/",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	}

	book, err := database.SeedBook()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bookID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodDelete, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues(bookID)

	if assert.NoError(t, bookController.Delete(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestDeleteBookById_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "invalid",
		path:                   "/auth/books/",
		expectedStatus:         http.StatusBadRequest,
	}

	req := httptest.NewRequest(http.MethodDelete, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues("100")

	if assert.NoError(t, bookController.Delete(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}
}







