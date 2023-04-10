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

var blogController BlogController = InitBlogContoller()

func TestGetAllBlogs_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/blogs",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"blog\":",
	}

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)

	if assert.NoError(t, blogController.GetAll(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetBlogById_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/blogs",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"blog\":",
	}

	blog, err := database.SeedBlog()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	blogID := strconv.Itoa(int(blog.ID))

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues(blogID)

	if assert.NoError(t, blogController.GetById(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestGetBlogById_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "invalid",
		path:                   "/blogs",
		expectedStatus:         http.StatusBadRequest,
	}

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues("100")

	if assert.NoError(t, blogController.GetById(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)

	}
}

func TestCreateBlog_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/blogs",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"blog\":",
	}

	user, err := database.SeedUser()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	var blogInput models.BlogInput = models.BlogInput{
		Title: "dummy",
		Content: "dummy",
		UserID: user.ID,
	}

	jsonBody, err := json.Marshal(&blogInput)

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

	if assert.NoError(t, blogController.Create(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestCreateBlog_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "invalid",
		path:                   "/blogs",
		expectedStatus:         http.StatusBadRequest,
	}

	user, err := database.SeedUser()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	var blogInput models.BlogInput = models.BlogInput{
		Title: "dummy",
		Content: "dummy",
		UserID: user.ID,
	}

	jsonBody, err := json.Marshal(&blogInput)

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)
	// req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)

	if assert.NoError(t, blogController.Create(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestUpdateBlogById_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/blogs",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"blog\":",
	}

	blog, err := database.SeedBlog()
	
	if err != nil {
		t.Errorf("error creating blog: %v\n", err)
	}
	
	blogID := strconv.Itoa(int(blog.ID))

	blogUpdate := models.BlogInput{
		Title: "update blog",
		Content: "update blog",
	}

	jsonBody, err := json.Marshal(&blogUpdate)

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
	context.SetParamValues(blogID)

	if assert.NoError(t, blogController.Update(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))

		updatedBlog := models.Blog{}
		err := database.DB.First(&updatedBlog, blogID).Error
		if err != nil {
			t.Errorf("error retrieving blog from database: %v\n", err)
		}
		assert.Equal(t, blogUpdate.Title, updatedBlog.Title)
		assert.Equal(t, blogUpdate.Content, updatedBlog.Content)
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestUpdateBlogById_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "invalid",
		path:                   "/blogs",
		expectedStatus:         http.StatusBadRequest,
	}

	blogUpdate := models.BlogInput{
		Title: "update blog",
		Content: "update blog",
	}

	jsonBody, err := json.Marshal(&blogUpdate)

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
	context.SetParamValues("100")

	if assert.NoError(t, blogController.Update(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestDeleteBlogById_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/auth/blogs/",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	}

	blog, err := database.SeedBlog()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	blogID := strconv.Itoa(int(blog.ID))

	req := httptest.NewRequest(http.MethodDelete, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues(blogID)

	if assert.NoError(t, blogController.Delete(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestDeleteBlogById_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "invalid",
		path:                   "/auth/blogs/",
		expectedStatus:         http.StatusBadRequest,
	}

	req := httptest.NewRequest(http.MethodDelete, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues("100")

	if assert.NoError(t, blogController.Delete(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}
}







