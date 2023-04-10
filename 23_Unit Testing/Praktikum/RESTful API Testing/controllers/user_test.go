package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restful_api_testing/database"
	"restful_api_testing/middlewares"
	"restful_api_testing/models"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var jwtConfig middlewares.JWTConfig = middlewares.JWTConfig{
	SecretKey:       "sssstt_rahasia",
	ExpiresDuration: 1,
}

var userController UserController = InitUserContoller(&jwtConfig)

func TestGetAllUsers_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/auth/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	}

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)

	if assert.NoError(t, userController.GetAll(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetUserById_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/auth/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	}

	user, err := database.SeedUser()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	userID := strconv.Itoa(int(user.ID))

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues(userID)

	if assert.NoError(t, userController.GetById(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestGetUserById_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/auth/users",
		expectedStatus:         http.StatusBadRequest,
	}

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues("100")

	if assert.NoError(t, userController.GetById(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}
}

func TestCreateUser_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/skipAuth/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte("dummy"), bcrypt.DefaultCost) 
	
	if err != nil {
		t.Errorf("error create password: %v\n", err)
	}

	var userInput models.UserInput = models.UserInput{
		Name: "dummy",
		Email: "dummy",
		Password: string(hashedPass),
	}

	jsonBody, err := json.Marshal(&userInput)

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

	if assert.NoError(t, userController.Create(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestCreateUser_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "invalid",
		path:                   "/skipAuth/users",
		expectedStatus:         http.StatusBadRequest,
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte("dummy"), bcrypt.DefaultCost) 
	
	if err != nil {
		t.Errorf("error create password: %v\n", err)
	}

	var userInput models.UserInput = models.UserInput{
		Name: "dummy",
		Email: "dummy",
		Password: string(hashedPass),
	}

	jsonBody, err := json.Marshal(&userInput)

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

	if assert.NoError(t, userController.Create(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}
}

func TestUpdateUserById_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/auth/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	}

	user, err := database.SeedUser()
	
	if err != nil {
		t.Errorf("error creating user: %v\n", err)
	}

	userID := strconv.Itoa(int(user.ID))

	var userUpdate models.UserInput = models.UserInput{
		Name: "updated user",
		Email: "updated user",
		Password: "updated user",
	}

	jsonBody, err := json.Marshal(&userUpdate)

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
	context.SetParamValues(userID)
	

	if assert.NoError(t, userController.Update(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))

		updatedUser := models.User{}
		err := database.DB.First(&updatedUser, userID).Error
		if err != nil {
			t.Errorf("error retrieving user from database: %v\n", err)
		}
		assert.Equal(t, userUpdate.Name, updatedUser.Name)
		assert.Equal(t, userUpdate.Email, updatedUser.Email)
		err = bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte("updated user"))
		assert.NoError(t, err)
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestUpdateUserById_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "invalid",
		path:                   "/auth/users",
		expectedStatus:         http.StatusBadRequest,
	}

	var userUpdate models.UserInput = models.UserInput{
		Name: "updated user",
		Email: "updated user",
		Password: "updated user",
	}

	jsonBody, err := json.Marshal(&userUpdate)

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
	context.SetParamValues("100")
	

	if assert.NoError(t, userController.Update(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}
}

func TestDeleteUserById_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/auth/users/",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	}

	user, err := database.SeedUser()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	userID := strconv.Itoa(int(user.ID))

	req := httptest.NewRequest(http.MethodDelete, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues(userID)

	if assert.NoError(t, userController.Delete(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestDeleteUserById_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "invalid",
		path:                   "/auth/users/",
		expectedStatus:         http.StatusBadRequest,
	}

	req := httptest.NewRequest(http.MethodDelete, testcase.path, nil)

	recorder := httptest.NewRecorder()

	e := InitEcho()

	context := e.NewContext(req, recorder)
	context.SetPath(testcase.path)
	context.SetParamNames("id")
	context.SetParamValues("100")

	if assert.NoError(t, userController.Delete(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}
}

func TestLogin_Valid(t *testing.T) {
	testcase := testCase{
		name:                   "valid",
		path:                   "/skipAuth/users/login",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	}

	user, err := database.SeedUser()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	var userInput models.UserInput = models.UserInput{
		Name: user.Name,
		Email: user.Email,
		Password: "dummy",
	}

	jsonBody, err := json.Marshal(&userInput)

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

	if assert.NoError(t, userController.Login(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}

func TestLogin_Invalid(t *testing.T) {
	testcase := testCase{
		name:                   "invalid",
		path:                   "/skipAuth/users/login",
		expectedStatus:         http.StatusBadRequest,
	}

	user, err := database.SeedUser()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	var userInput models.UserInput = models.UserInput{
		Name: user.Name,
		Email: user.Email,
		Password: "dummy",
	}

	jsonBody, err := json.Marshal(&userInput)

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

	if assert.NoError(t, userController.Login(context)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	}

	t.Cleanup(func() {
		database.CleanSeeders()
	})
}








